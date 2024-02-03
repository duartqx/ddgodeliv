package services

import (
	"fmt"

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
	if err := vs.vehicleRepository.Create(vehicle); err != nil {
		return fmt.Errorf("Internal Error creating vehicle: %v", err.Error())
	}

	return nil
}

func (vs VehicleService) FindById(vehicle ve.IVehicle) error {
	if err := vs.vehicleRepository.FindById(vehicle); err != nil {
		return fmt.Errorf("Internal Error trying to locate vehicle: %v", err.Error())
	}
	return nil
}

func (vs VehicleService) FindByCompanyId(companyId int) (*[]ve.IVehicle, error) {
	if companyId == 0 {
		return nil, fmt.Errorf("Invalid Company Id")
	}

	vehicles, err := vs.vehicleRepository.FindByCompanyId(companyId)

	if err != nil {
		return nil, fmt.Errorf("Internal Error trying to locate vehicles: %v", err.Error())
	}

	return vehicles, nil
}

func (vs VehicleService) Delete(vehicle ve.IVehicle) error {
	if vehicle.GetId() == 0 {
		return fmt.Errorf("Invalid Vehicle Id")
	}
	if vehicle.GetCompanyId() == 0 {
		return fmt.Errorf("Invalid Company Id")
	}
	return vs.vehicleRepository.Delete(vehicle)
}
