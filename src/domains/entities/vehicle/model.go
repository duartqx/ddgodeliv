package vehicle

import m "ddgodeliv/domains/models"

type VehicleModel struct {
	Id           int    `db:"id" json:"id"`
	Manufacturer string `db:"manufacturer" json:"manufacturer"`
	Year         int    `db:"year" json:"year"`
	MaxLoad      int    `db:"max_load" json:"max_load"`
}

func GetNewVehicleModel() *VehicleModel {
	return &VehicleModel{}
}

func (m VehicleModel) GetId() int {
	return m.Id

}

func (m *VehicleModel) SetId(id int) m.IVehicleModel {
	m.Id = id
	return m
}

func (m VehicleModel) GetManufacturer() string {
	return m.Manufacturer
}

func (m *VehicleModel) SetManufacturer(manufacturer string) m.IVehicleModel {
	m.Manufacturer = manufacturer
	return m
}

func (m VehicleModel) GetYear() int {
	return m.Year
}

func (m *VehicleModel) SetYear(year int) m.IVehicleModel {
	m.Year = year
	return m
}

func (m VehicleModel) GetMaxLoad() int {
	return m.MaxLoad
}

func (m *VehicleModel) SetMaxLoad(maxLoad int) m.IVehicleModel {
	m.MaxLoad = maxLoad
	return m
}
