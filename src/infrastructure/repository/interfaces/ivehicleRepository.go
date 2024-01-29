package repository

import (
	m "ddgodeliv/domains/models"
)

type IVehicleModelRepository interface {
	FindById(id int) (m.IVehicleModel, error)

	Create(model m.IVehicleModel) error
	Update(model m.IVehicleModel) error
	Delete(model m.IVehicleModel) error
}

type IVehicleRepository interface {
	FindById(id int) (m.IVehicle, error)

	FindByCompanyId(id int) (*[]m.IVehicle, error)
	ExistsByCompanyId(id int) *bool

	Create(vehicle m.IVehicle) error
	Update(vehicle m.IVehicle) error
	Delete(vehicle m.IVehicle) error
}
