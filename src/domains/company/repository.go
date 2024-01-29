package company

type ICompanyRepository interface {
	FindById(id int) (ICompany, error)

	ExistsByName(name string) bool

	Create(company ICompany, licenseId string) error
	Delete(company ICompany) error
}
