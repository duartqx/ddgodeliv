package vehicle

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

type IVehicle interface {
	GetId() int
	SetId(id int) IVehicle

	GetModelId() int
	SetModelId(modelId int) IVehicle

	GetLicenseId() string
	SetLicenseId(licenseId string) IVehicle

	GetModel() IVehicleModel
}
