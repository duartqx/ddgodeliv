package controllers

import (
	"encoding/json"
	"net/http"

	h "ddgodeliv/api/http"
	s "ddgodeliv/application/services"
	e "ddgodeliv/common/errors"
	v "ddgodeliv/domains/vehicle"
)

type VehicleModelController struct {
	vehicleModelService *s.VehicleModelService
}

func GetNewVehicleModelController(
	vehicleModelService *s.VehicleModelService,
) *VehicleModelController {
	return &VehicleModelController{vehicleModelService: vehicleModelService}
}

func (vmc VehicleModelController) CreateVehicleModel(
	w http.ResponseWriter, r *http.Request,
) {

	vehicleModel := v.GetNewVehicleModel()

	if err := json.NewDecoder(r.Body).Decode(&vehicleModel); err != nil {
		http.Error(w, e.BadRequestError.Error(), http.StatusBadRequest)
		return
	}

	if err := vmc.vehicleModelService.Create(vehicleModel); err != nil {
		h.ErrorResponse(w, err)
		return
	}

	h.JsonResponse(w, http.StatusCreated, vehicleModel)
}

func (vmc VehicleModelController) ListModels(w http.ResponseWriter, r *http.Request) {
	models, err := vmc.vehicleModelService.All()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	h.JsonResponse(w, http.StatusOK, models)
}
