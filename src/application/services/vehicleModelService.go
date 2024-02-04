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

	if validationErrs := vms.ValidateStruct(vehicleModel); validationErrs != nil {
		return validationErrs
	}

	if err := vms.vehicleModelRepository.Create(vehicleModel); err != nil {
		return fmt.Errorf("Internal Error creating vehicle model: %v", err.Error())
	}
	return nil
}

func (vms VehicleModelService) All() (*[]ve.IVehicleModel, error) {
	models, err := vms.vehicleModelRepository.All()
	if err != nil {
		return nil, fmt.Errorf("Internal Error trying to list all models: %v", err.Error())
	}
	return models, nil
}
