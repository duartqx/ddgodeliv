package repository

import (
	m "ddgodeliv/domains/models"
)

type ICompanyRepository interface {
	FindById(id int) (m.ICompany, error)

	ExistsByName(name string) bool

	Create(company m.ICompany, licenseId string) error
	Delete(company m.ICompany) error
}
