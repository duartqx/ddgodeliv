package services

import (
	"fmt"

	v "ddgodeliv/application/validation"
	m "ddgodeliv/domains/models"
	r "ddgodeliv/infrastructure/repository/interfaces"
)

type DriverService struct {
	driverRepository r.IDriverRepository
	userService      *UserService
	// emailService     *EmailService
	*v.Validator
}

func GetNewDriverService(
	driverRepository r.IDriverRepository,
	userService *UserService,
	validator *v.Validator,
	// emailService *EmailService
) *DriverService {
	return &DriverService{
		driverRepository: driverRepository,
		userService:      userService,
		Validator:        validator,
		// emailService:     emailService,
	}
}

// Asks the new user to activate it's account and warns the company owner
// of the new driver creation
func (cs DriverService) sendDriverCreationEmails(driver m.IDriver) error {
	return nil
}

// Warns the user and the company owner of the driver deletion
func (cs DriverService) sendDriverDeletionEmails(driver m.IDriver) error {
	return nil
}

// s: Driver | User
func (cs DriverService) Validate(s interface{}) error {
	if errs := cs.Struct(s); errs != nil {
		return fmt.Errorf(string(*cs.JSON(errs)))
	}
	return nil
}

func (ds DriverService) CreateDriver(driver m.IDriver) error {
	if err := ds.Struct(driver); err != nil {
		return err
	}

	if err := ds.Struct(driver.GetUser()); err != nil {
		return err
	}

	driver.GetUser().SetPassword("TempPasswordToChangeWhenActivatingAccount")

	if err := ds.userService.Create(driver.GetUser()); err != nil {
		return fmt.Errorf("Internal Error creating the user: %v", err.Error())
	}

	if err := ds.driverRepository.Create(driver); err != nil {
		return fmt.Errorf("Internal Error creating the driver: %v", err.Error())
	}

	// Asks the new user to activate it's account and warns the company owner
	// of the new driver creation
	if err := ds.sendDriverCreationEmails(driver); err != nil {
		return fmt.Errorf("Internal Error sending Driver Creation Emails: %v", err.Error())
	}

	return nil
}

func (ds DriverService) DeleteDriver(driver m.IDriver) error {
	if driver.GetId() == 0 || driver.GetUserId() == 0 || driver.GetCompanyId() == 0 {
		return fmt.Errorf("Invalid Driver")
	}

	if driver.GetUserId() == driver.GetCompany().GetOwnerId() {
		return fmt.Errorf("The company owner can't delete it's own driver")
	}

	if err := ds.driverRepository.Delete(driver); err != nil {
		return fmt.Errorf("Internal Error deleting the driver: %v", err.Error())
	}

	// Warns the user and the company owner of the driver deletion
	if err := ds.sendDriverDeletionEmails(driver); err != nil {
		return fmt.Errorf("Internal Error sending Driver Deletion Emails: %v", err.Error())
	}

	return nil
}

func (ds DriverService) UpdateDriverLicense(driver m.IDriver) error {
	if err := ds.Validator.Var(driver.GetLicenseId(), "required,min=3,max=250"); err != nil {
		return fmt.Errorf("Invalid Driver License")
	}

	if err := ds.driverRepository.Update(driver); err != nil {
		return fmt.Errorf("Internal Error trying to update the driver")
	}

	return nil
}
