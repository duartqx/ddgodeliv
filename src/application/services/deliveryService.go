package services

import (
	d "ddgodeliv/domains/delivery"
)

type DeliveryService struct {
	deliveryRepository d.IDeliveryRepository
}

func GetNewDeliveryService(deliveryRepository d.IDeliveryRepository) *DeliveryService {
	return &DeliveryService{deliveryRepository: deliveryRepository}
}

func (ds DeliveryService) Create(delivery d.IDelivery) error {
	return nil
}
