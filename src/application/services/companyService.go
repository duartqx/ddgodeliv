package services

import (
	e "ddgodeliv/application/errors"
	v "ddgodeliv/application/validation"
	c "ddgodeliv/domains/company"
	re "ddgodeliv/infrastructure/repository"
	"fmt"
)

type CompanyService struct {
	companyRepository re.ICompanyRepository
	*v.Validator
}

func GetNewCompanyService(companyRepository re.ICompanyRepository, validator *v.Validator) *CompanyService {
	return &CompanyService{companyRepository: companyRepository, Validator: validator}
}

func (cs CompanyService) Validate(company c.ICompany) error {
	if errs := cs.Struct(company); errs != nil {
		return fmt.Errorf(string(*cs.JSON(errs)))
	}
	return nil
}

func (cs CompanyService) CreateCompany(userId int, company c.ICompany) error {

	if err := cs.Validate(company); err != nil {
		return err
	}

	if cs.companyRepository.ExistsByName(company.GetName()) {
		return e.BadRequestError
	}

	if err := cs.companyRepository.Create(userId, company); err != nil {
		return e.InternalError
	}

	return nil
}

func (cs CompanyService) DeleteCompany(userId int, company c.ICompany) error {
	// Only the owner can Delete it's company
	if userId != company.GetOwnerId() || company.GetId() == 0 {
		return e.BadRequestError
	}

	return cs.companyRepository.Delete(company)
}
