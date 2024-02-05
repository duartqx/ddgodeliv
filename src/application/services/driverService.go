package services

import (
	"fmt"

	v "ddgodeliv/application/validation"
	e "ddgodeliv/common/errors"
	d "ddgodeliv/domains/driver"
)

type DriverService struct {
	driverRepository d.IDriverRepository
	userService      *UserService
	// emailService     *EmailService
	validator *v.Validator
}

func GetNewDriverService(
	driverRepository d.IDriverRepository,
	userService *UserService,
	// emailService *EmailService
) *DriverService {
	return &DriverService{
		driverRepository: driverRepository,
		userService:      userService,
		validator:        v.NewValidator(),
		// emailService:     emailService,
	}
}

// Asks the new user to activate it's account and warns the company owner
// of the new driver creation
func (cs DriverService) sendDriverCreationEmails(driver d.IDriver) error {
	return nil
}

// Warns the user and the company owner of the driver deletion
func (cs DriverService) sendDriverDeletionEmails(driver d.IDriver) error {
	return nil
}

// A Company owner automatically has an driver created for its user, then all
// other drivers of it's company are created by a manager
func (ds DriverService) CreateDriver(driver d.IDriver) error {
	if err := ds.validator.ValidateStruct(driver); err != nil {
		return err
	}

	if err := ds.validator.ValidateStruct(driver.GetUser()); err != nil {
		return err
	}

	driver.GetUser().SetPassword(
		"TempPasswordToChangeWhenActivatingAccount" + driver.GetUser().GetEmail(),
	)

	if err := ds.userService.Create(driver.GetUser()); err != nil {
		return fmt.Errorf("Internal Error creating the user: %w", err)
	}

	if err := ds.driverRepository.Create(driver); err != nil {
		return fmt.Errorf("Internal Error creating the driver: %w", err)
	}

	// Asks the new user to activate it's account and warns the company owner
	// of the new driver creation
	if err := ds.sendDriverCreationEmails(driver); err != nil {
		return fmt.Errorf("Internal Error sending Driver Creation Emails: %w", err)
	}

	return nil
}

func (ds DriverService) DeleteDriver(ownerId int, driver d.IDriver) error {

	switch {
	case driver.GetId() == 0 || driver.GetUserId() == 0 || driver.GetCompanyId() == 0:
		return fmt.Errorf("Invalid Driver: %w", e.ForbiddenError)
	case ownerId == 0:
		return fmt.Errorf("Invalid Owner: %w", e.ForbiddenError)
	case ownerId != driver.GetCompanyId():
		return fmt.Errorf(
			"Only the onwer can delete it's drivers: %w",
			e.ForbiddenError,
		)
	case driver.GetUserId() == driver.GetCompany().GetOwnerId():
		return fmt.Errorf(
			"The company owner can't delete it's own driver: %w",
			e.ForbiddenError,
		)
	}

	if err := ds.driverRepository.Delete(driver); err != nil {
		return fmt.Errorf("Internal Error deleting the driver: %w", err)
	}

	// Warns the user and the company owner of the driver deletion
	if err := ds.sendDriverDeletionEmails(driver); err != nil {
		return fmt.Errorf("Internal Error sending Driver Deletion Emails: %w", err)
	}

	return nil
}

func (ds DriverService) UpdateDriverLicense(userId int, driver d.IDriver) error {
	switch {
	case driver.GetId() == 0:
		return fmt.Errorf("Invalid Driver: %w", e.BadRequestError)
	case userId != driver.GetUserId() || userId != driver.GetCompany().GetOwnerId():
		return fmt.Errorf(
			"Only Company owner or the driver can change their license: %w",
			e.ForbiddenError,
		)
	}

	if err := ds.validator.Var(driver.GetLicenseId(), "required,min=3,max=250"); err != nil {
		return err
	}

	if err := ds.driverRepository.Update(driver); err != nil {
		return fmt.Errorf("Internal Error trying to update the driver")
	}

	return nil
}

func (ds DriverService) FindByUserId(id int) (d.IDriver, error) {
	driver, err := ds.driverRepository.FindByUserId(id)
	if err != nil {
		return nil, fmt.Errorf("Error trying to find user driver: %w", err)
	}
	return driver, nil
}

func (ds DriverService) FindById(id, companyId int) (d.IDriver, error) {
	driver, err := ds.driverRepository.FindById(id, companyId)
	if err != nil {
		return nil, fmt.Errorf("Error trying to find user driver: %w", err)
	}
	return driver, nil
}

func (ds DriverService) ListCompanyDriversById(id int) (*[]d.IDriver, error) {
	drivers, err := ds.driverRepository.FindByCompanyId(id)
	if err != nil {
		return nil, fmt.Errorf("Error trying to find user driver: %w", err)
	}
	return drivers, nil
}
