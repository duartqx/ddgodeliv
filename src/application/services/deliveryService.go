package services

import (
	m "ddgodeliv/domains/models"
	r "ddgodeliv/infrastructure/repository/interfaces"
)

type DeliveryService struct {
	deliveryRepository r.IDeliveryRepository
}

func GetNewDeliveryService(deliveryRepository r.IDeliveryRepository) *DeliveryService {
	return &DeliveryService{deliveryRepository: deliveryRepository}
}

func (ds DeliveryService) Create(delivery m.IDelivery) error {
	return nil
}
