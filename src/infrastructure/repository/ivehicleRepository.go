package repository

import (
	v "ddgodeliv/src/domains/vehicle"
)

type IVehicleModelRepository interface {
	FindById(id int) (v.IVehicleModel, error)

	Create(model v.IVehicleModel) error
	Update(model v.IVehicleModel) error
	Delete(model v.IVehicleModel) error
}

type IVehicleRepository interface {
	FindById(id int) (v.IVehicle, error)

	FindByCompanyId(id int) (*[]v.IVehicle, error)
	ExistsByCompanyId(id int) *bool

	Create(vehicle v.IVehicle) error
	Update(vehicle v.IVehicle) error
	Delete(vehicle v.IVehicle) error
}
