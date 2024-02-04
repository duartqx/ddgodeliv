package vehicle

type VehicleModel struct {
	Id           int    `db:"id" json:"id"`
	Name         string `db:"name" json:"name" validate:"required,min=3"`
	Manufacturer string `db:"manufacturer" json:"manufacturer" validate:"required,min=3"`
	Year         int    `db:"year" json:"year" validate:"gt=1900,lt=9999"`
	MaxLoad      int    `db:"max_load" json:"max_load" validate:"required"`
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

func (m VehicleModel) GetMaxLoad() int {
	return m.MaxLoad
}

func (m *VehicleModel) SetMaxLoad(maxLoad int) IVehicleModel {
	m.MaxLoad = maxLoad
	return m
}
