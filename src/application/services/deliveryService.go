package services

import (
	"fmt"

	v "ddgodeliv/application/validation"
	e "ddgodeliv/common/errors"
	c "ddgodeliv/domains/company"
	de "ddgodeliv/domains/delivery"
	d "ddgodeliv/domains/driver"
	u "ddgodeliv/domains/user"
)

type DeliveryService struct {
	deliveryRepository de.IDeliveryRepository
	driverRepository   d.IDriverRepository
	validator          *v.Validator
}

func GetNewDeliveryService(
	deliveryRepository de.IDeliveryRepository,
	driverRepository d.IDriverRepository,
) *DeliveryService {
	return &DeliveryService{
		deliveryRepository: deliveryRepository,
		driverRepository:   driverRepository,
		validator:          v.NewValidator(),
	}
}

func (ds DeliveryService) Create(delivery de.IDelivery) error {
	if !delivery.DriverIsNull() {
		// Checks if the driver struct is populated and has valid .CompanyId
		if !delivery.GetDriver().HasValidCompanyId() {
			return fmt.Errorf(
				"Driver does not have a valid CompanyId: %w",
				e.BadRequestError,
			)
		}
		// If has .DriverId, then checks if the driver exists
		if !ds.driverRepository.ExistsById(
			delivery.GetDriverId(), delivery.GetDriver().GetCompanyId(),
		) {
			return fmt.Errorf("Driver does not Exists: %w", e.NotFoundError)
		}
	}

	if err := ds.validator.ValidateStruct(delivery); err != nil {
		return err
	}

	return ds.deliveryRepository.Create(delivery)
}

func (ds DeliveryService) AssignDriver(delivery de.IDelivery, driver d.IDriver) error {

	// We need the id to populate and update the delivery
	if delivery.HasInvalidId() {
		return fmt.Errorf("%w: Invalid Delivery", e.BadRequestError)
	}

	// Populates delivery information
	if err := ds.deliveryRepository.FindById(delivery); err != nil {
		return err
	}

	// If is not pending, returns an error
	if !(delivery.GetStatus() == de.StatusChoices.Pending) {
		return fmt.Errorf(
			"%w: This delivery is not pending", e.BadRequestError,
		)
	}

	if driver.HasInvalidId() || !driver.HasValidCompanyId() {
		return fmt.Errorf("%w: Invalid Driver", e.BadRequestError)
	}

	// Checks if driver.Id + driver.CompanyId are a valid driver (avoids
	// assigning to driver from different company)
	if err := ds.driverRepository.FindById(driver); err != nil {
		return err
	}

	delivery.SetDriver(driver).SetStatus(de.StatusChoices.Assigned)
	if err := ds.deliveryRepository.Update(delivery); err != nil {
		return err
	}

	return nil
}

func (ds DeliveryService) UpdateStatus(delivery de.IDelivery) error {

	status := delivery.GetStatus()

	if err := ds.validator.ValidateVar(status, "required,gte=0,lte=4"); err != nil {
		return err
	}

	if err := ds.deliveryRepository.FindById(delivery); err != nil {
		// TODO: Refactor this
		return fmt.Errorf("%w: Invalid Delivery", err)
	}

	if !delivery.GetDriver().HasInvalidId() && status == de.StatusChoices.Pending {
		return fmt.Errorf(
			"%w: Delivery has driver, can't be set to pending",
			e.BadRequestError,
		)
	} else if delivery.GetDriver().HasInvalidId() && status > de.StatusChoices.Pending {
		return fmt.Errorf(
			"%w: Delivery has no driver, can only be set to pending",
			e.BadRequestError,
		)
	}

	if err := ds.deliveryRepository.Update(delivery); err != nil {
		return fmt.Errorf("%w: %s", e.InternalError, err.Error())
	}

	return nil
}

func (ds DeliveryService) FindById(user u.IUser, delivery de.IDelivery) error {
	if delivery.HasInvalidId() {
		return fmt.Errorf("%w: Invalid Delivery", e.BadRequestError)
	}

	if err := ds.deliveryRepository.FindById(delivery); err != nil {
		return e.NotFoundError
	}

	// The sender can always see their delivery
	if delivery.GetSenderId() == user.GetId() {
		return nil
	}

	driver, _ := ds.driverRepository.FindByUserId(user.GetId())

	// Allows all drivers to see the delivery if it's pending
	// Otherwise only allows drivers from same company
	if driver != nil && (delivery.IsPending() || delivery.GetDriver().GetCompanyId() == driver.GetCompanyId()) {
		return nil
	}

	return e.NotFoundError
}

func (ds DeliveryService) FindByCompanyId(company c.ICompany) (*[]de.IDelivery, error) {
	if company.HasInvalidId() {
		return nil, fmt.Errorf("%w: Invalid Company", e.BadRequestError)
	}
	deliveries, err := ds.deliveryRepository.FindByCompanyId(company.GetId())
	if err != nil {
		return nil, fmt.Errorf("%w: %v", e.InternalError, err.Error())
	}
	return deliveries, nil
}

func (ds DeliveryService) Delete(user u.IUser, delivery de.IDelivery) error {
	if delivery.HasInvalidId() {
		return fmt.Errorf("%w: Invalid Delivery", e.BadRequestError)
	}

	// TODO: Validate if user is delivery.Sender or if the delivery has
	// DriverId, check if this user is from the same company as the driver
	// assigned to this delivery

	if err := ds.deliveryRepository.Delete(delivery); err != nil {
		return fmt.Errorf("%w: %v", e.InternalError, err.Error())
	}

	return nil
}

func (ds DeliveryService) FindPendingWithoutDriver() (*[]de.IDelivery, error) {
	deliveries, err := ds.deliveryRepository.FindPendingWithNoDriver()
	if err != nil {
		return nil, fmt.Errorf("%w: %v", e.InternalError, err.Error())
	}
	return deliveries, nil
}

func (ds DeliveryService) FindBySenderId(user u.IUser) (*[]de.IDelivery, error) {
	deliveries, err := ds.deliveryRepository.FindBySenderId(user.GetId())
	if err != nil {
		return nil, fmt.Errorf("%w: %v", e.InternalError, err.Error())
	}
	return deliveries, nil
}
