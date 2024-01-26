package vehicle

type Vehicle struct {
	Id        int    `db:"id" json:"id"`
	ModelId   int    `db:"model_id" json:"model_id"`
	LicenseId string `db:"license_id" json:"license_id"`

	Model IVehicleModel `json:"model"`
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
