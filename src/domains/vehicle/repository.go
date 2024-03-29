package vehicle

type IVehicleModelRepository interface {
	FindById(id int) (IVehicleModel, error)
	All() (*[]IVehicleModel, error)

	Create(model IVehicleModel) error
	Update(model IVehicleModel) error
	Delete(model IVehicleModel) error
}

type IVehicleRepository interface {
	FindById(vehicle IVehicle) error

	FindByCompanyId(id int) (*[]IVehicle, error)
	ExistsByCompanyId(id int) bool

	ModelExists(id int) bool

	Create(vehicle IVehicle) error
	Update(vehicle IVehicle) error
	Delete(vehicle IVehicle) error
}
