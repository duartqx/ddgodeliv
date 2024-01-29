package repository

import (
	c "ddgodeliv/domains/company"
)

type ICompanyRepository interface {
	FindById(id int) (c.ICompany, error)

	ExistsByName(name string) bool

	Create(ownerId int, company c.ICompany) error
	Delete(company c.ICompany) error
}
