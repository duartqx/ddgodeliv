package services

import (
	d "ddgodeliv/domains/delivery"
	r "ddgodeliv/infrastructure/repository"
)

type DeliveryService struct {
	deliveryRepository r.IDeliveryRepository
}

func GetNewDeliveryService(deliveryRepository r.IDeliveryRepository) *DeliveryService {
	return &DeliveryService{deliveryRepository: deliveryRepository}
}

func (ds DeliveryService) Create(delivery d.IDelivery) error {
	return nil
}
