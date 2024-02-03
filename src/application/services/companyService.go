package services

import (
	"fmt"

	v "ddgodeliv/application/validation"
	c "ddgodeliv/domains/company"
)

type CompanyService struct {
	companyRepository c.ICompanyRepository
	*v.Validator
}

func GetNewCompanyService(companyRepository c.ICompanyRepository, validator *v.Validator) *CompanyService {
	return &CompanyService{companyRepository: companyRepository, Validator: validator}
}

func (cs CompanyService) CreateCompany(company c.ICompany, licenseId string) error {

	if cs.companyRepository.ExistsByName(company.GetName()) {
		return fmt.Errorf("Company with this name already exists!")
	}

	if err := cs.companyRepository.Create(company, licenseId); err != nil {
		return fmt.Errorf("Internal: %v", err.Error())
	}

	return nil
}

func (cs CompanyService) DeleteCompany(userId int, company c.ICompany) error {

	if company.HasInvalidId() {
		return fmt.Errorf("Invalid Company Id")
	}

	if userId != company.GetOwnerId() {
		return fmt.Errorf("Only the owner can Delete a company!")
	}

	return cs.companyRepository.Delete(company)
}

func (cs CompanyService) FindById(company c.ICompany) error {

	if company.HasInvalidId() {
		return fmt.Errorf("Invalid Company Id")
	}

	if err := cs.companyRepository.FindById(company); err != nil {
		return fmt.Errorf("Internal Error trying to find company")
	}
	return nil
}
