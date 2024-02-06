package auth

import (
	d "ddgodeliv/domains/driver"
	u "ddgodeliv/domains/user"
)

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
	ResetDriver() ISessionUser
	SetFromAnother(user ISessionUser)

	HasInvalidCompany() bool
	ToUser() u.IUser
}
