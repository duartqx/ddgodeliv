package postgresql

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"

	v "ddgodeliv/domains/vehicle"
)

type VehicleRepository struct {
	db *sqlx.DB
}

func GetNewVehicleRepository(db *sqlx.DB) *VehicleModelRepository {
	return &VehicleModelRepository{db: db}
}

func (vr VehicleRepository) simpleValidate(vehicle v.IVehicle) error {
	if vehicle.GetModelId() == 0 || vehicle.GetCompanyId() == 0 || vehicle.GetLicenseId() == "" {
		return fmt.Errorf("Invalid vehicle: Missing Model, Company or License")
	}
	return nil
}

func (vr VehicleRepository) FindById(vehicle v.IVehicle) error {
	if err := vr.db.Get(
		vehicle,
		"SELECT * FROM vehicles WHERE id = $1 AND company_id = $2",
		vehicle.GetId(), vehicle.GetCompanyId(),
	); err != nil {
		return err
	}
	return nil
}

func (vr VehicleRepository) FindByCompanyId(id int) (*[]v.IVehicle, error) {
	vehicles := []v.IVehicle{}

	rows, err := vr.db.Query("SELECT * FROM vehicles WHERE company_id = $1", id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {

		vehicle := v.GetNewVehicle()

		if err := rows.Scan(vehicle); err != nil {
			return nil, err
		}

		var castedVehicle v.IVehicle = vehicle

		vehicles = append(vehicles, castedVehicle)
	}

	return &vehicles, nil

}

func (vr VehicleRepository) ExistsByCompanyId(id int) (exists *bool) {
	vr.db.QueryRow(
		"SELECT EXISTS (SELECT 1 FROM vehicles WHERE company_id = $1)",
		id,
	).Scan(&exists)

	return exists
}

func (vr VehicleRepository) Create(vehicle v.IVehicle) error {
	if err := vr.simpleValidate(vehicle); err != nil {
		return err
	}

	var id int

	if err := vr.db.QueryRow(
		`
			INSERT INTO vehicles (model_id, company_id, license_id)
			VALUES ($1, $2, $3)
			RETURNING id
		`,
		vehicle.GetModelId(),
		vehicle.GetCompanyId(),
		strings.ToLower(vehicle.GetLicenseId()),
	).Scan(&id); err != nil {
		return err
	}

	vehicle.SetId(id)

	return nil
}

func (vr VehicleRepository) Update(vehicle v.IVehicle) error {
	if err := vr.simpleValidate(vehicle); err != nil {
		return err
	}
	if vehicle.GetId() == 0 {
		return fmt.Errorf("Invalid Vehicle Id")
	}

	res, err := vr.db.Exec(
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

	if count, err := res.RowsAffected(); err != nil {
		return fmt.Errorf("Error trying to count affected rows: %v", err.Error())
	} else if count < 1 {
		return fmt.Errorf("No rows were affected!")
	}
	return err
}

func (vr VehicleRepository) Delete(vehicle v.IVehicle) error {
	if vehicle.GetId() == 0 {
		return fmt.Errorf("Invalid Vehicle Id")
	}

	res, err := vr.db.Exec(
		"DELETE FROM vehicles WHERE id = $1 AND company_id = $2",
		vehicle.GetId(),
		vehicle.GetCompanyId(),
	)
	if err != nil {
		return fmt.Errorf("Error trying to exec Delete query: %v", err.Error())
	}

	if count, err := res.RowsAffected(); err != nil {
		return fmt.Errorf("Error trying to count affected rows: %v", err.Error())
	} else if count < 1 {
		return fmt.Errorf("No rows were affected!")
	}

	return nil
}
