package services

import (
	"fmt"

	v "ddgodeliv/application/validation"
	m "ddgodeliv/domains/models"
	r "ddgodeliv/infrastructure/repository/interfaces"
)

type CompanyService struct {
	companyRepository r.ICompanyRepository
	*v.Validator
}

func GetNewCompanyService(companyRepository r.ICompanyRepository, validator *v.Validator) *CompanyService {
	return &CompanyService{companyRepository: companyRepository, Validator: validator}
}

func (cs CompanyService) Validate(company m.ICompany) error {
	if errs := cs.Struct(company); errs != nil {
		return fmt.Errorf(string(*cs.JSON(errs)))
	}
	return nil
}

func (cs CompanyService) CreateCompany(company m.ICompany, licenseId string) error {

	if err := cs.Validate(company); err != nil {
		return err
	}

	if cs.companyRepository.ExistsByName(company.GetName()) {
		return fmt.Errorf("Company with this name already exists!")
	}

	if err := cs.companyRepository.Create(company, licenseId); err != nil {
		return fmt.Errorf("Internal: %v", err.Error())
	}

	return nil
}

func (cs CompanyService) DeleteCompany(userId int, company m.ICompany) error {

	if company.GetId() == 0 {
		return fmt.Errorf("Invalid Company Id")
	}

	if userId != company.GetOwnerId() {
		return fmt.Errorf("Only the owner can Delete a company!")
	}

	return cs.companyRepository.Delete(company)
}
