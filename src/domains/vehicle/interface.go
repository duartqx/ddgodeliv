package vehicle

import c "ddgodeliv/domains/company"

type IVehicle interface {
	GetId() int
	SetId(id int) IVehicle

	GetModelId() int
	SetModelId(modelId int) IVehicle

	GetCompanyId() int
	SetCompanyId(companyId int) IVehicle

	GetLicenseId() string
	SetLicenseId(licenseId string) IVehicle

	GetModel() IVehicleModel
	SetModel(model IVehicleModel) IVehicle

	GetCompany() c.ICompany
	SetCompany(company c.ICompany) IVehicle
}

type IVehicleModel interface {
	GetId() int
	SetId(id int) IVehicleModel

	GetName() string
	SetName(name string) IVehicleModel

	GetManufacturer() string // Normalize
	SetManufacturer(manufacturer string) IVehicleModel

	GetYear() int
	SetYear(year int) IVehicleModel

	GetTransmission() string
	SetTransmission(transmission string) IVehicleModel

	GetType() string
	SetType(modelType string) IVehicleModel
}
