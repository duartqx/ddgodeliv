package postgresql

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	v "ddgodeliv/domains/vehicle"
)

type VehicleModelRepository struct {
	db *sqlx.DB
}

func GetNewVehicleModelRepository(db *sqlx.DB) *VehicleModelRepository {
	return &VehicleModelRepository{db: db}
}

func (vmr VehicleModelRepository) FindById(id int) (v.IVehicleModel, error) {

	vehicleModel := v.GetNewVehicleModel()

	if err := vmr.db.Get(
		vehicleModel,
		"SELECT * FROM users WHERE id = $1",
		id,
	); err != nil {
		return nil, err
	}

	return vehicleModel, nil
}

func (vmr VehicleModelRepository) All() (*[]v.IVehicleModel, error) {
	models := []v.IVehicleModel{}

	rows, err := vmr.db.Queryx("SELECT * FROM vehiclemodels")
	if err != nil {
		return nil, err
	}

	for rows.Next() {

		driver := v.GetNewVehicleModel()

		if err := rows.StructScan(driver); err != nil {
			return nil, err
		}

		var castedDriver v.IVehicleModel = driver

		models = append(models, castedDriver)
	}

	return &models, nil
}

func (vmr VehicleModelRepository) Create(model v.IVehicleModel) error {
	var id int

	if err := vmr.db.QueryRow(
		`
			INSERT INTO vehiclemodels (name, manufacturer, year, transmission, type)
			VALUES ($1, $2, $3, $4, $5)
			RETURNING id
		`,
		model.GetName(),
		model.GetManufacturer(),
		model.GetYear(),
		model.GetTransmission(),
		model.GetType(),
	).Scan(&id); err != nil {
		return err
	}

	model.SetId(id)

	return nil
}

func (vmr VehicleModelRepository) Update(model v.IVehicleModel) error {
	_, err := vmr.db.Exec(
		`
			UPDATE vehiclemodels
			SET name = $1, manufacturer = $2, year = $3, transmission = $4, type = $5
			WHERE id = $2
		`,
		model.GetName(),
		model.GetManufacturer(),
		model.GetYear(),
		model.GetTransmission(),
		model.GetType(),
	)

	return err
}

func (vmr VehicleModelRepository) Delete(model v.IVehicleModel) error {
	if model.GetId() == 0 {
		return fmt.Errorf("Invalid Vehicle Model Id")
	}

	_, err := vmr.db.Exec("DELETE FROM vehiclemodels WHERE id = $1", model.GetId())

	return err

}
