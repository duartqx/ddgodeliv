package auth

import d "ddgodeliv/domains/driver"

type sessionDriver struct {
	Id        int `json:"driver_id"`
	CompanyId int `json:"company_id"`
}

type SessionUser struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`

	Driver sessionDriver `json:"driver"`
}

func GetNewSessionUser() *SessionUser {
	return &SessionUser{
		Driver: sessionDriver{},
	}
}

func (u SessionUser) GetId() int {
	return u.Id
}

func (u *SessionUser) SetId(id int) ISessionUser {
	u.Id = id
	return u
}

func (u SessionUser) GetEmail() string {
	return u.Email
}

func (u *SessionUser) SetEmail(email string) ISessionUser {
	u.Email = email
	return u
}

func (u SessionUser) GetName() string {
	return u.Name
}

func (u *SessionUser) SetName(name string) ISessionUser {
	u.Name = name
	return u
}

func (u *SessionUser) SetDriver(driver d.IDriver) ISessionUser {
	u.Driver.Id = driver.GetId()
	u.Driver.CompanyId = driver.GetCompanyId()
	return u
}

func (u *SessionUser) ResetDriver() ISessionUser {
	u.Driver.Id = 0
	u.Driver.CompanyId = 0
	return u
}

func (u SessionUser) HasInvalidCompany() bool {
	return u.Driver.CompanyId == 0
}

func (u SessionUser) GetCompanyId() int {
	return u.Driver.CompanyId
}

func (u SessionUser) GetDriverId() int {
	return u.Driver.Id
}

func (u *SessionUser) SetFromAnother(user ISessionUser) {
	u.Id = user.GetId()
	u.Email = user.GetEmail()
	u.Name = user.GetName()
	u.Driver.Id = user.GetDriverId()
	u.Driver.CompanyId = user.GetCompanyId()
}
