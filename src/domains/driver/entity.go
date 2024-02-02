package driver

import (
	c "ddgodeliv/domains/company"
	u "ddgodeliv/domains/user"
)

type Driver struct {
	Id        int    `db:"id" json:"id"`
	UserId    int    `db:"user_id" json:"user_id" validate:"required,gt=0"`
	CompanyId int    `db:"company_id" json:"company_id" validate:"required,gt=0"`
	LicenseId string `db:"license_id" json:"license_id" validate:"required,min=3,max=250"`

	User    u.IUser    `json:"user"`
	Company c.ICompany `json:"company"`
}

func GetNewDriver() *Driver {
	return &Driver{
		User:    u.GetNewUser(),
		Company: c.GetNewCompany(),
	}
}

func (d Driver) GetId() int {
	return d.Id
}

func (d *Driver) SetId(id int) IDriver {
	d.Id = id
	return d
}

func (d Driver) GetUserId() int {
	return d.UserId
}

func (d *Driver) SetUserId(userId int) IDriver {
	d.UserId = userId
	return d
}

func (d Driver) GetLicenseId() string {
	return d.LicenseId
}

func (d *Driver) SetLicenseId(licenseId string) IDriver {
	d.LicenseId = licenseId
	return d
}

func (d Driver) GetCompanyId() int {
	return d.CompanyId
}

func (d *Driver) SetCompanyId(companyId int) IDriver {
	d.CompanyId = companyId
	return d
}

func (d Driver) GetUser() u.IUser {
	return d.User
}

func (d Driver) GetCompany() c.ICompany {
	return d.Company
}

func (d Driver) HasInvalidId() bool {
	return d.GetId() == 0
}
