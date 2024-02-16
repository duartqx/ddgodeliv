package postgresql

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/jmoiron/sqlx"

	e "ddgodeliv/common/errors"
	v "ddgodeliv/domains/vehicle"
)

type VehicleRepository struct {
	db *sqlx.DB
}

var vehicleRepository *VehicleRepository

func GetVehicleRepository(db *sqlx.DB) *VehicleRepository {
	if vehicleRepository == nil {
		vehicleRepository = &VehicleRepository{db: db}
	}
	return vehicleRepository
}

func (vr VehicleRepository) FindById(vehicle v.IVehicle) error {
	if err := vr.db.Get(
		vehicle,
		"SELECT * FROM vehicles WHERE id = $1 AND company_id = $2",
		vehicle.GetId(), vehicle.GetCompanyId(),
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return e.NotFoundError
		}
		return err
	}
	return nil
}

func (vr VehicleRepository) FindByCompanyId(id int) (*[]v.IVehicle, error) {
	vehicles := []v.IVehicle{}

	rows, err := vr.db.Queryx(
		`SELECT
			v.id AS id,
			v.model_id AS model_id,
			v.company_id AS company_id,
			v.license_id AS license_id,

			m.id AS "model.id",
			m.name AS "model.name",
			m.manufacturer AS "model.manufacturer",
			m.year AS "model.year",
			m.transmission AS "model.transmission",
			m.type AS "model.type",

			c.id AS "company.id",
			c.owner_id AS "company.owner_id",
			c.name AS "company.name"
		FROM vehicles v
		INNER JOIN vehiclemodels m ON v.model_id = m.id
		INNER JOIN companies c ON v.company_id = c.id
		WHERE company_id = $1`,
		id,
	)
	if err != nil {
		return nil, err
	}

	for rows.Next() {

		vehicle := v.GetNewVehicle()

		if err := rows.StructScan(vehicle); err != nil {
			return nil, err
		}

		var castedVehicle v.IVehicle = vehicle

		vehicles = append(vehicles, castedVehicle)
	}

	return &vehicles, nil

}

func (vr VehicleRepository) ExistsByCompanyId(id int) (exists bool) {
	vr.db.QueryRow(
		"SELECT EXISTS (SELECT 1 FROM vehicles WHERE company_id = $1)",
		id,
	).Scan(&exists)

	return exists
}

func (vr VehicleRepository) ModelExists(id int) (exists bool) {
	vr.db.QueryRow(
		"SELECT EXISTS (SELECT 1 FROM vehiclemodels WHERE id = $1)",
		id,
	).Scan(&exists)

	return exists
}

func (vr VehicleRepository) Create(vehicle v.IVehicle) error {
	if err := vr.db.Get(
		vehicle,
		`
			WITH new_vehicle AS (
				INSERT INTO vehicles (model_id, company_id, license_id)
				VALUES ($1, $2, $3)
				RETURNING id, model_id
			)
			SELECT 
				v.id AS "id",

				m.id AS "model.id",
				m.name AS "model.name",
				m.manufacturer AS "model.manufacturer",
				m.year AS "model.year",
				m.transmission AS "model.transmission",
				m.type AS "model.type"
			FROM new_vehicle v
			INNER JOIN vehiclemodels m ON m.id = v.model_id
		`,
		vehicle.GetModelId(),
		vehicle.GetCompanyId(),
		strings.ToLower(vehicle.GetLicenseId()),
	); err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (vr VehicleRepository) Update(vehicle v.IVehicle) error {
	_, err := vr.db.Exec(
		`
			UPDATE vehicles
			SET model_id = $1, company_id = $2, license_id = $3
			WHERE id = $2 AND company_id = $2
		`,
		vehicle.GetModelId(),
		vehicle.GetCompanyId(),
		strings.ToLower(vehicle.GetLicenseId()),
	)
	if err != nil {
		return fmt.Errorf("Error trying to update vehicle: %v", err.Error())
	}

	return err
}

func (vr VehicleRepository) Delete(vehicle v.IVehicle) error {
	if vehicle.GetId() == 0 {
		return fmt.Errorf("Invalid Vehicle Id")
	}

	_, err := vr.db.Exec(
		"DELETE FROM vehicles WHERE id = $1 AND company_id = $2",
		vehicle.GetId(),
		vehicle.GetCompanyId(),
	)
	if err != nil {
		return fmt.Errorf("Error trying to exec Delete query: %v", err.Error())
	}

	return nil
}
