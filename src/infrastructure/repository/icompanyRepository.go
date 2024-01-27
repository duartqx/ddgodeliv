package repository

import (
	c "ddgodeliv/src/domains/company"
)

type ICompanyRepository interface {
	FindById(id int) (c.ICompany, error)

	Create(company c.ICompany) error
	Delete(company c.ICompany) error
}
