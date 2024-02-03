package auth

import d "ddgodeliv/domains/driver"

type ISessionUser interface {
	GetId() int
	GetEmail() string
	GetName() string

	GetDriverId() int
	GetCompanyId() int

	SetId(id int) ISessionUser
	SetEmail(email string) ISessionUser
	SetName(name string) ISessionUser

	SetDriver(driver d.IDriver) ISessionUser
	SetFromAnother(user ISessionUser)

	HasInvalidCompany() bool
}
