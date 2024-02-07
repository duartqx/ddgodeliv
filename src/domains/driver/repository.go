package driver

type IDriverRepository interface {
	FindById(driver IDriver) error
	ExistsById(id, companyId int) bool

	FindByUserId(id int) (IDriver, error)
	ExistsByUserId(id int) bool

	FindByCompanyId(id int) (*[]IDriver, error)
	ExistsByCompanyId(id int) bool

	Create(driver IDriver) error
	Update(driver IDriver) error
	Delete(driver IDriver) error
}
