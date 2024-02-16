package services

import (
	"fmt"

	v "ddgodeliv/application/validation"
	ve "ddgodeliv/domains/vehicle"
)

type VehicleModelService struct {
	vehicleModelRepository ve.IVehicleModelRepository
	validator              *v.Validator
}

var vehicleModelService *VehicleModelService

func GetVehicleModelService(
	vehicleModelRepository ve.IVehicleModelRepository,
) *VehicleModelService {
	if vehicleModelService == nil {
		vehicleModelService = &VehicleModelService{
			vehicleModelRepository: vehicleModelRepository,
			validator:              v.NewValidator(),
		}
	}
	return vehicleModelService
}

func (vms VehicleModelService) Create(vehicleModel ve.IVehicleModel) error {

	if validationErrs := vms.validator.ValidateStruct(vehicleModel); validationErrs != nil {
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
