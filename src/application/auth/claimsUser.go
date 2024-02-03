package auth

import d "ddgodeliv/domains/driver"

type ClaimsUser struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`

	Driver struct {
		Id        int `json:"driver_id"`
		CompanyId int `json:"company_id"`
	} `json:"driver"`
}

func (u ClaimsUser) GetId() int {
	return u.Id
}

func (u ClaimsUser) GetEmail() string {
	return u.Email
}

func (u ClaimsUser) GetName() string {
	return u.Name
}

func (u *ClaimsUser) SetDriver(driver d.IDriver) *ClaimsUser {
	u.Driver.Id = driver.GetId()
	u.Driver.CompanyId = driver.GetCompanyId()
	return u
}

func (u ClaimsUser) HasInvalidCompany() bool {
	return u.Driver.CompanyId == 0
}

func (u ClaimsUser) GetCompanyId() int {
	return u.Driver.CompanyId
}

func (u ClaimsUser) GetDriverId() int {
	return u.Driver.Id
}

func (u *ClaimsUser) SetFromAnother(user *ClaimsUser) {
	u.Id = user.Id
	u.Email = user.Email
	u.Name = user.Name
	u.Driver.Id = user.Driver.Id
	u.Driver.CompanyId = user.Driver.CompanyId
}
