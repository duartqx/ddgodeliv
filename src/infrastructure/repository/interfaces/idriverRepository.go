package repository

import (
	m "ddgodeliv/domains/models"
)

type IDriverRepository interface {
	FindById(id int) (m.IDriver, error)

	FindByUserId(id int) (m.IDriver, error)
	ExistsByUserId(id int) *bool

	FindByCompanyId(id int) (*[]m.IDriver, error)
	ExistsByCompanyId(id int) *bool

	Create(driver m.IDriver) error
	Update(driver m.IDriver) error
	Delete(driver m.IDriver) error
}
