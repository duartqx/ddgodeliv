package services

import (
	"fmt"

	v "ddgodeliv/application/validation"
	e "ddgodeliv/common/errors"
	ve "ddgodeliv/domains/vehicle"
)

type VehicleService struct {
	vehicleRepository ve.IVehicleRepository
	validator         *v.Validator
}

func GetNewVehicleService(
	vehicleRepository ve.IVehicleRepository,
) *VehicleService {
	return &VehicleService{
		vehicleRepository: vehicleRepository,
		validator:         v.NewValidator(),
	}
}

func (vs VehicleService) Create(vehicle ve.IVehicle) error {
	if validationErrs := vs.validator.ValidateStruct(vehicle); validationErrs != nil {
		return validationErrs
	}

	if !vs.vehicleRepository.ModelExists(vehicle.GetModelId()) {
		return fmt.Errorf("Model does not exists: %w", e.BadRequestError)
	}

	// TODO: Check if license is not unique

	if err := vs.vehicleRepository.Create(vehicle); err != nil {
		return err
	}

	return nil
}

func (vs VehicleService) FindById(vehicle ve.IVehicle) error {
	if err := vs.vehicleRepository.FindById(vehicle); err != nil {
		return err
	}
	return nil
}

func (vs VehicleService) FindByCompanyId(companyId int) (*[]ve.IVehicle, error) {
	if companyId == 0 {
		return nil, fmt.Errorf("Invalid Company Id: %w", e.BadRequestError)
	}

	vehicles, err := vs.vehicleRepository.FindByCompanyId(companyId)

	if err != nil {
		return nil, err
	}

	return vehicles, nil
}

func (vs VehicleService) Delete(vehicle ve.IVehicle) error {
	if vehicle.GetId() == 0 {
		return fmt.Errorf("Invalid Vehicle Id: %w", e.BadRequestError)
	}
	if vehicle.GetCompanyId() == 0 {
		return fmt.Errorf("Invalid Company Id: %w", e.BadRequestError)
	}
	return vs.vehicleRepository.Delete(vehicle)
}
