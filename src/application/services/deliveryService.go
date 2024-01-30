package services

import (
	v "ddgodeliv/application/validation"
	d "ddgodeliv/domains/delivery"
)

type DeliveryService struct {
	deliveryRepository d.IDeliveryRepository
	*v.Validator
}

func GetNewDeliveryService(
	deliveryRepository d.IDeliveryRepository,
	validator *v.Validator,
) *DeliveryService {
	return &DeliveryService{
		deliveryRepository: deliveryRepository,
		Validator:          validator,
	}
}

func (ds DeliveryService) Create(delivery d.IDelivery) error {
	if err := ds.ValidateStruct(delivery); err != nil {
		return err
	}
	return ds.deliveryRepository.Create(delivery)
}
