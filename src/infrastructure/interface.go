package infrastructure

import (
	"time"

	c "ddgodeliv/src/domains/company"
	d "ddgodeliv/src/domains/driver"
	f "ddgodeliv/src/domains/freight"
	u "ddgodeliv/src/domains/user"
	v "ddgodeliv/src/domains/vehicle"
)

type IUserRepository interface {
	FindById(id int) (u.IUser, error)

	FindByEmail(email string) (u.IUser, error)
	ExistsByEmail(email string) *bool

	Create(user u.IUser) error
	Update(user u.IUser) error
	Delete(user u.IUser) error
}

type ICompanyRepository interface {
	FindById(id int) (c.ICompany, error)

	Create(company c.ICompany) error
	Update(company c.ICompany) error
	Delete(company c.ICompany) error
}

type IDriverRepository interface {
	FindById(id int) (d.IDriver, error)

	FindByUserId(id int) (d.IDriver, error)
	ExistsByUserId(id int) *bool

	FindByCompanyId(id int) (*[]d.IDriver, error)
	ExistsByCompanyId(id int) *bool

	Create(driver d.IDriver) error
	Update(driver d.IDriver) error
	Delete(driver d.IDriver) error
}

type IFreightRepository interface {
	FindById(id int) (f.IFreight, error)

	FindByDriverId(id int) (*[]f.IFreight, error)
	ExistsByDriverId(id int) *bool

	FindByDeadlineDateRange(start, end time.Time) (*[]f.IFreight, error)
	FindByDeadlineDate(deadline time.Time) (*[]f.IFreight, error)
	ExistsByDeadlineDate(deadline time.Time) *bool

	FindBySenderId(id int) (*[]f.IFreight, error)
	ExistsBySenderId(id int) *bool

	FindByCompanyId(id int) (*[]f.IFreight, error)
	ExistsByCompanyId(id int) *bool

	Create(freight f.IFreight) error
	Update(freight f.IFreight) error
	Delete(freight f.IFreight) error
}

type IVehicleModelRepository interface {
	FindById(id int) (v.IVehicleModel, error)

	Create(model v.IVehicleModel) error
	Update(model v.IVehicleModel) error
	Delete(model v.IVehicleModel) error
}

type IVehicleRepository interface {
	FindById(id int) (v.IVehicle, error)

	FindByCompanyId(id int) (*[]v.IVehicle, error)
	ExistsByCompanyId(id int) *bool

	Create(vehicle v.IVehicle) error
	Update(vehicle v.IVehicle) error
	Delete(vehicle v.IVehicle) error
}
