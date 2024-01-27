package repository

import (
	d "ddgodeliv/src/domains/driver"
)

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
