package services

import (
	"fmt"

	v "ddgodeliv/application/validation"
	c "ddgodeliv/domains/company"
	de "ddgodeliv/domains/delivery"
	d "ddgodeliv/domains/driver"
	u "ddgodeliv/domains/user"
)

type DeliveryService struct {
	deliveryRepository de.IDeliveryRepository
	*v.Validator
}

func GetNewDeliveryService(
	deliveryRepository de.IDeliveryRepository,
	validator *v.Validator,
) *DeliveryService {
	return &DeliveryService{
		deliveryRepository: deliveryRepository,
		Validator:          validator,
	}
}

func (ds DeliveryService) Create(delivery de.IDelivery) error {
	if err := ds.ValidateStruct(delivery); err != nil {
		return err
	}
	return ds.deliveryRepository.Create(delivery)
}

func (ds DeliveryService) AssignDriver(delivery de.IDelivery, driver d.IDriver) error {
	if delivery.HasInvalidId() {
		return fmt.Errorf("Invalid Delivery!")
	}

	if driver.HasInvalidId() {
		return fmt.Errorf("Invalid Driver!")
	}

	err := ds.deliveryRepository.Update(delivery.SetDriverId(driver.GetId()))
	if err != nil {
		return fmt.Errorf("Internal Error while trying to update driver: %v", err.Error())
	}

	return nil
}

func (ds DeliveryService) UpdateStatus(delivery de.IDelivery) error {
	if err := ds.ValidateVar(delivery.GetStatus(), "required,gte=0,lte=4"); err != nil {
		return err
	}

	if err := ds.deliveryRepository.Update(delivery); err != nil {
		return fmt.Errorf("Internal Error while trying to update status: %v", err.Error())
	}

	return nil
}

func (ds DeliveryService) FindById(user u.IUser, delivery de.IDelivery) error {
	if delivery.HasInvalidId() {
		return fmt.Errorf("Invalid Delivery!")
	}

	// TODO: Validate if user is delivery.Sender or if the delivery has
	// DriverId, check if this user is from the same company as the driver
	// assigned to this delivery

	if err := ds.deliveryRepository.FindById(delivery); err != nil {
		return fmt.Errorf("Internal Error while trying to find delivery: %v", err.Error())
	}

	return nil
}

func (ds DeliveryService) FindByCompany(company c.ICompany) (*[]de.IDelivery, error) {
	if company.HasInvalidId() {
		return nil, fmt.Errorf("Invalid Company!")
	}
	deliveries, err := ds.deliveryRepository.FindByCompanyId(company.GetId())
	if err != nil {
		return nil, fmt.Errorf("Internal Error trying to find by company: %v", err.Error())
	}
	return deliveries, nil
}

func (ds DeliveryService) Delete(user u.IUser, delivery de.IDelivery) error {
	if delivery.HasInvalidId() {
		return fmt.Errorf("Invalid Delivery!")
	}

	// TODO: Validate if user is delivery.Sender or if the delivery has
	// DriverId, check if this user is from the same company as the driver
	// assigned to this delivery

	if err := ds.deliveryRepository.Delete(delivery); err != nil {
		return fmt.Errorf("Internal Error trying to delete delivery: %v", err.Error())
	}

	return nil
}
