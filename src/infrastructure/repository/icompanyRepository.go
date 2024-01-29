package repository

import (
	c "ddgodeliv/domains/company"
)

type ICompanyRepository interface {
	FindById(id int) (c.ICompany, error)

	ExistsByName(name string) bool

	Create(company c.ICompany, licenseId string) error
	Delete(company c.ICompany) error
}
