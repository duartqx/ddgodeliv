package services

import (
	v "ddgodeliv/application/validation"
	ve "ddgodeliv/domains/vehicle"
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
	if err := vms.ValidateStruct(vehicleModel); err != nil {
		return err
	}
	return vms.vehicleModelRepository.Create(vehicleModel)
}
