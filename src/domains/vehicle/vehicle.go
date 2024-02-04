package vehicle

import c "ddgodeliv/domains/company"

type Vehicle struct {
	Id        int    `db:"id" json:"id"`
	ModelId   int    `db:"model_id" json:"model_id" validate:"required,gte=1"`
	CompanyId int    `db:"company_id" json:"company_id" validate:"required,gte=1"`
	LicenseId string `db:"license_id" json:"license_id" validate:"required"`

	Model   VehicleModel `db:"model" json:"model" validate:"-"`
	Company c.Company    `db:"company" json:"company" validate:"-"`
}

func GetNewVehicle() *Vehicle {
	return &Vehicle{
		Model:   *GetNewVehicleModel(),
		Company: *c.GetNewCompany(),
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
	return &v.Model
}

func (v *Vehicle) SetModel(model IVehicleModel) IVehicle {
	v.Model.
		SetId(model.GetId()).
		SetYear(model.GetYear()).
		SetMaxLoad(model.GetMaxLoad()).
		SetManufacturer(model.GetManufacturer())
	return v
}

func (v Vehicle) GetCompany() c.ICompany {
	return &v.Company
}

func (v *Vehicle) SetCompany(company c.ICompany) IVehicle {
	v.Company.
		SetId(company.GetId()).
		SetName(company.GetName()).
		SetOwnerId(company.GetOwnerId())
	return v
}
