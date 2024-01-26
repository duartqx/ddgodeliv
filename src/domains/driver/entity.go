package driver

import (
	"ddgodeliv/src/domains/company"
	"ddgodeliv/src/domains/user"
)

type IDriver interface {
	GetId() int
	SetId(id int) IDriver

	GetUserId() int
	SetUserId(userId int) IDriver

	GetLicenseId() string
	SetLicenseId(licenseId string) IDriver

	GetCompanyId() int
	SetCompanyId(companyId int) IDriver

	GetUser() user.IUser
	GetCompany() company.ICompany
}

type Driver struct {
	Id        int    `db:"id" json:"id"`
	UserId    int    `db:"user_id" json:"user_id"`
	CompanyId int    `db:"company_id" json:"company_id"`
	LicenseId string `db:"license_id" json:"license_id"`

	User    user.IUser       `json:"user"`
	Company company.ICompany `json:"company"`
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

func (d Driver) GetUser() user.IUser {
	return d.User
}

func (d Driver) GetCompany() company.ICompany {
	return d.Company
}
