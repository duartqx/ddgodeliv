package vehicle

type VehicleModel struct {
	Id           int    `db:"id" json:"id"`
	Name         string `db:"name" json:"name" validate:"required,min=3"`
	Manufacturer string `db:"manufacturer" json:"manufacturer" validate:"required,min=3"`
	Year         int    `db:"year" json:"year" validate:"gt=1900,lt=9999"`
	Transmission string `db:"transmission" json:"transmission" validate:"required"`
	Type         string `db:"type" json:"type" validate:"required"`
}

func GetNewVehicleModel() *VehicleModel {
	return &VehicleModel{}
}

func (m VehicleModel) GetId() int {
	return m.Id

}

func (m *VehicleModel) SetId(id int) IVehicleModel {
	m.Id = id
	return m
}

func (m VehicleModel) GetName() string {
	return m.Name

}

func (m *VehicleModel) SetName(name string) IVehicleModel {
	m.Name = name
	return m
}

func (m VehicleModel) GetManufacturer() string {
	return m.Manufacturer
}

func (m *VehicleModel) SetManufacturer(manufacturer string) IVehicleModel {
	m.Manufacturer = manufacturer
	return m
}

func (m VehicleModel) GetYear() int {
	return m.Year
}

func (m *VehicleModel) SetYear(year int) IVehicleModel {
	m.Year = year
	return m
}

func (m VehicleModel) GetTransmission() string {
	return m.Transmission
}

func (m *VehicleModel) SetTransmission(transmission string) IVehicleModel {
	m.Transmission = transmission
	return m
}

func (m VehicleModel) GetType() string {
	return m.Type
}

func (m *VehicleModel) SetType(modelType string) IVehicleModel {
	m.Type = modelType
	return m
}
