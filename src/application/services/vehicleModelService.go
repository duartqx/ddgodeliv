package services

import (
	v "ddgodeliv/application/validation"
	ve "ddgodeliv/domains/vehicle"
	"fmt"
)

type VehicleModelService struct {
	vehicleModelRepository ve.IVehicleModelRepository
	*v.Validator
}

func GetNewVehicleModelService(
	vehicleModelRepository ve.IVehicleModelRepository,
	validator *v.Validator,
) *VehicleModelService {
	return &VehicleModelService{
		vehicleModelRepository: vehicleModelRepository,
		Validator:              validator,
	}
}

func (vms VehicleModelService) Create(vehicleModel ve.IVehicleModel) error {
	if err := vms.vehicleModelRepository.Create(vehicleModel); err != nil {
		return fmt.Errorf("Internal Error creating vehicle model: %v", err.Error())
	}
	return nil
}
