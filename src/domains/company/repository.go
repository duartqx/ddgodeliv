package company

type ICompanyRepository interface {
	FindById(company ICompany) error

	ExistsByName(name string) bool

	Create(company ICompany, licenseId string) error
	Delete(company ICompany) error
}
