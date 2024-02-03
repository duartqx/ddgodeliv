package services

import (
	"fmt"

	v "ddgodeliv/application/validation"
	d "ddgodeliv/domains/driver"
)

type DriverService struct {
	driverRepository d.IDriverRepository
	userService      *UserService
	// emailService     *EmailService
	*v.Validator
}

func GetNewDriverService(
	driverRepository d.IDriverRepository,
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
	if err := ds.ValidateStruct(driver); err != nil {
		return err
	}

	if err := ds.ValidateStruct(driver.GetUser()); err != nil {
		return err
	}

	driver.GetUser().SetPassword(
		"TempPasswordToChangeWhenActivatingAccount" + driver.GetUser().GetEmail(),
	)

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

func (ds DriverService) DeleteDriver(driver d.IDriver) error {
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

func (ds DriverService) UpdateDriverLicense(driver d.IDriver) error {
	if err := ds.Validator.Var(driver.GetLicenseId(), "required,min=3,max=250"); err != nil {
		return fmt.Errorf("Invalid Driver License")
	}

	if err := ds.driverRepository.Update(driver); err != nil {
		return fmt.Errorf("Internal Error trying to update the driver")
	}

	return nil
}

func (ds DriverService) FindByUserId(id int) (d.IDriver, error) {
	driver, err := ds.FindByUserId(id)
	if err != nil {
		return nil, fmt.Errorf("Error trying to find user driver: %v", err.Error())
	}
	return driver, nil
}
