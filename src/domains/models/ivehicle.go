package models

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

	GetCompany() ICompany
	SetCompany(company ICompany) IVehicle
}
