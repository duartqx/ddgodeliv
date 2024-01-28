package driver

import (
	"ddgodeliv/domains/company"
	"ddgodeliv/domains/user"
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
