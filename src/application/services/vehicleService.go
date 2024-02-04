package services

import (
	"fmt"

	e "ddgodeliv/application/errors"
	v "ddgodeliv/application/validation"
	ve "ddgodeliv/domains/vehicle"
)

type VehicleService struct {
	vehicleRepository ve.IVehicleRepository
	*v.Validator
}

func GetNewVehicleService(
	vehicleRepository ve.IVehicleRepository,
	validator *v.Validator,
) *VehicleService {
	return &VehicleService{
		vehicleRepository: vehicleRepository,
		Validator:         validator,
	}
}

func (vs VehicleService) Create(vehicle ve.IVehicle) error {
	if validationErrs := vs.ValidateStruct(vehicle); validationErrs != nil {
		return validationErrs
	}

	if !vs.vehicleRepository.ModelExists(vehicle.GetModelId()) {
		return fmt.Errorf("Model does not exists: %w", e.BadRequestError)
	}

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
