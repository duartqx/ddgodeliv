package models

type IVehicleModel interface {
	GetId() int
	SetId(id int) IVehicleModel

	GetManufacturer() string // Normalize
	SetManufacturer(manufacturer string) IVehicleModel

	GetYear() int
	SetYear(year int) IVehicleModel

	GetMaxLoad() int
	SetMaxLoad(maxLoad int) IVehicleModel
}
