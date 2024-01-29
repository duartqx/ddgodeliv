package driver

import (
	c "ddgodeliv/domains/entities/company"
	u "ddgodeliv/domains/entities/user"
	m "ddgodeliv/domains/models"
)

type Driver struct {
	Id        int    `db:"id" json:"id"`
	UserId    int    `db:"user_id" json:"user_id" validate:"required,gt=0"`
	CompanyId int    `db:"company_id" json:"company_id" validate:"required,gt=0"`
	LicenseId string `db:"license_id" json:"license_id" validate:"required,min=3,max=250"`

	User    m.IUser    `json:"user"`
	Company m.ICompany `json:"company"`
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

func (d *Driver) SetId(id int) m.IDriver {
	d.Id = id
	return d
}

func (d Driver) GetUserId() int {
	return d.UserId
}

func (d *Driver) SetUserId(userId int) m.IDriver {
	d.UserId = userId
	return d
}

func (d Driver) GetLicenseId() string {
	return d.LicenseId
}

func (d *Driver) SetLicenseId(licenseId string) m.IDriver {
	d.LicenseId = licenseId
	return d
}

func (d Driver) GetCompanyId() int {
	return d.CompanyId
}

func (d *Driver) SetCompanyId(companyId int) m.IDriver {
	d.CompanyId = companyId
	return d
}

func (d Driver) GetUser() m.IUser {
	return d.User
}

func (d Driver) GetCompany() m.ICompany {
	return d.Company
}
