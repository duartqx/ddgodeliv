package services

import (
	"fmt"

	v "ddgodeliv/application/validation"
	e "ddgodeliv/common/errors"
	c "ddgodeliv/domains/company"
)

type CompanyService struct {
	companyRepository c.ICompanyRepository
	validator         *v.Validator
}

func GetNewCompanyService(companyRepository c.ICompanyRepository) *CompanyService {
	return &CompanyService{
		companyRepository: companyRepository, validator: v.NewValidator(),
	}
}

func (cs CompanyService) CreateCompany(company c.ICompany, licenseId string) error {

	if err := cs.validator.ValidateVar(licenseId, "required,min=3,max=250"); err != nil {
		return fmt.Errorf("Invalid Driver License: %w", e.BadRequestError)
	}

	if validationsErrs := cs.validator.ValidateStruct(company); validationsErrs != nil {
		return validationsErrs
	}

	if cs.companyRepository.ExistsByName(company.GetName()) {
		return fmt.Errorf(
			"Company with this name already exists: %w", e.BadRequestError,
		)
	}

	if err := cs.companyRepository.Create(company, licenseId); err != nil {
		return fmt.Errorf("Internal: %v - %w", err.Error(), e.InternalError)
	}

	return nil
}

func (cs CompanyService) DeleteCompany(userId int, company c.ICompany) error {

	if company.HasInvalidId() {
		return fmt.Errorf("Invalid Company Id: %w", e.BadRequestError)
	}

	if userId != company.GetOwnerId() {
		return fmt.Errorf(
			"Only the owner can Delete a company: %w", e.ForbiddenError,
		)
	}

	return cs.companyRepository.Delete(company)
}

func (cs CompanyService) FindById(company c.ICompany) error {

	if company.HasInvalidId() {
		return fmt.Errorf("Invalid Company Id: %w", e.BadRequestError)
	}

	if err := cs.companyRepository.FindById(company); err != nil {
		return fmt.Errorf("Internal Error trying to find company: %w", e.InternalError)
	}
	return nil
}

func (cs CompanyService) ValidateDriverLicense(license string) error {
	return cs.validator.ValidateVar(license, "required,min=3,max=250")
}
