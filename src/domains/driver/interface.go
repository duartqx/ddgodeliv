package driver

import (
	c "ddgodeliv/domains/company"
	u "ddgodeliv/domains/user"
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

	GetUser() u.IUser
	SetUser(user u.IUser) IDriver
	GetCompany() c.ICompany

	HasInvalidId() bool
	HasValidCompanyId() bool
}
