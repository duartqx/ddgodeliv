package vehicle

import c "ddgodeliv/domains/entities/company"
import m "ddgodeliv/domains/models"

type Vehicle struct {
	Id        int    `db:"id" json:"id"`
	ModelId   int    `db:"model_id" json:"model_id"`
	CompanyId int    `db:"company_id" json:"company_id"`
	LicenseId string `db:"license_id" json:"license_id"`

	Model   m.IVehicleModel `json:"model"`
	Company m.ICompany      `json:"company"`
}

func GetNewVehicle() *Vehicle {
	return &Vehicle{
		Model:   &VehicleModel{},
		Company: c.GetNewCompany(),
	}
}

func (v Vehicle) GetId() int {
	return v.Id
}

func (v *Vehicle) SetId(id int) m.IVehicle {
	v.Id = id
	return v
}

func (v Vehicle) GetModelId() int {
	return v.ModelId
}

func (v *Vehicle) SetModelId(modelId int) m.IVehicle {
	v.ModelId = modelId
	return v
}

func (v Vehicle) GetCompanyId() int {
	return v.CompanyId
}

func (v *Vehicle) SetCompanyId(companyId int) m.IVehicle {
	v.CompanyId = companyId
	return v
}

func (v Vehicle) GetLicenseId() string {
	return v.LicenseId
}

func (v *Vehicle) SetLicenseId(licenseId string) m.IVehicle {
	v.LicenseId = licenseId
	return v
}

func (v Vehicle) GetModel() m.IVehicleModel {
	return v.Model
}

func (v *Vehicle) SetModel(model m.IVehicleModel) m.IVehicle {
	v.Model = model
	return v
}

func (v Vehicle) GetCompany() m.ICompany {
	return v.Company
}

func (v *Vehicle) SetCompany(company m.ICompany) m.IVehicle {
	v.Company = company
	return v
}
