package vehicle

import "ddgodeliv/src/domains/company"

type Vehicle struct {
	Id        int    `db:"id" json:"id"`
	ModelId   int    `db:"model_id" json:"model_id"`
	CompanyId int    `db:"company_id" json:"company_id"`
	LicenseId string `db:"license_id" json:"license_id"`

	Model   IVehicleModel    `json:"model"`
	Company company.ICompany `json:"company"`
}

func GetNewVehicle() *Vehicle {
	return &Vehicle{
		Model:   &VehicleModel{},
		Company: &company.Company{},
	}
}

func (v Vehicle) GetId() int {
	return v.Id
}

func (v *Vehicle) SetId(id int) IVehicle {
	v.Id = id
	return v
}

func (v Vehicle) GetModelId() int {
	return v.ModelId
}

func (v *Vehicle) SetModelId(modelId int) IVehicle {
	v.ModelId = modelId
	return v
}

func (v Vehicle) GetCompanyId() int {
	return v.CompanyId
}

func (v *Vehicle) SetCompanyId(companyId int) IVehicle {
	v.CompanyId = companyId
	return v
}

func (v Vehicle) GetLicenseId() string {
	return v.LicenseId
}

func (v *Vehicle) SetLicenseId(licenseId string) IVehicle {
	v.LicenseId = licenseId
	return v
}

func (v Vehicle) GetModel() IVehicleModel {
	return v.Model
}

func (v *Vehicle) SetModel(model IVehicleModel) IVehicle {
	v.Model = model
	return v
}

func (v Vehicle) GetCompany() company.ICompany {
	return v.Company
}

func (v *Vehicle) SetCompany(company company.ICompany) IVehicle {
	v.Company = company
	return v
}
