package infrastructure

import (
	"time"

	"ddgodeliv/src/domains/company"
	"ddgodeliv/src/domains/driver"
	"ddgodeliv/src/domains/freight"
	"ddgodeliv/src/domains/user"
	"ddgodeliv/src/domains/vehicle"
)

type IUserRepository interface {
	FindById(id int) (user.IUser, error)

	FindByEmail(email string) (user.IUser, error)
	ExistsByEmail(email string) bool

	Create(user user.IUser) error
	Update(user user.IUser) error
	Delete(user user.IUser) error
}

type ICompanyRepository interface {
	FindById(id int) (company.ICompany, error)

	Create(company company.ICompany) error
	Update(company company.ICompany) error
	Delete(company company.ICompany) error
}

type IDriverRepository interface {
	FindById(id int) (driver.IDriver, error)

	FindByUserId(id int) (driver.IDriver, error)
	ExistsByUserId(id int) bool

	FindByCompanyId(id int) (*[]driver.IDriver, error)
	ExistsByCompanyId(id int) bool

	Create(driver driver.IDriver) error
	Update(driver driver.IDriver) error
	Delete(driver driver.IDriver) error
}

type IFreightRepository interface {
	FindById(id int) (freight.IFreight, error)

	FindByDriverId(id int) (*[]freight.IFreight, error)
	ExistsByDriverId(id int) bool

	FindByDeadlineDateRange(start, end time.Time) (*[]freight.IFreight, error)
	FindByDeadlineDate(deadline time.Time) (*[]freight.IFreight, error)
	ExistsByDeadlineDate(deadline time.Time) bool

	FindBySenderId(id int) (*[]freight.IFreight, error)
	ExistsBySenderId(id int) bool

	FindByCompanyId(id int) (*[]freight.IFreight, error)
	ExistsByCompanyId(id int) bool

	Create(freight freight.IFreight) error
	Update(freight freight.IFreight) error
	Delete(freight freight.IFreight) error
}

type IVehicleModelRepository interface {
	FindById(id int) (vehicle.IVehicleModel, error)

	Create(model vehicle.IVehicleModel) error
	Update(model vehicle.IVehicleModel) error
	Delete(model vehicle.IVehicleModel) error
}

type IVehicleRepository interface {
	FindById(id int) (vehicle.IVehicle, error)

	FindByCompanyId(id int) (*[]vehicle.IVehicle, error)
	ExistsByCompanyId(id int) bool

	Create(vehicle vehicle.IVehicle) error
	Update(vehicle vehicle.IVehicle) error
	Delete(vehicle vehicle.IVehicle) error
}
