package controllers

import (
	"encoding/json"
	"errors"
	"net/http"

	e "ddgodeliv/application/errors"
	s "ddgodeliv/application/services"
	v "ddgodeliv/domains/vehicle"
)

type vehicleModelController struct {
	vehicleModelService *s.VehicleModelService
}

func GetNewVehicleModelController(vehicleModelService *s.VehicleModelService) *vehicleModelController {
	return &vehicleModelController{vehicleModelService: vehicleModelService}
}

func (vmc vehicleModelController) CreateVehicleModel(w http.ResponseWriter, r *http.Request) {

	vehicleModel := v.GetNewVehicleModel()

	if err := json.NewDecoder(r.Body).Decode(&vehicleModel); err != nil {
		http.Error(w, e.BadRequestError.Error(), http.StatusBadRequest)
		return
	}

	if err := vmc.vehicleModelService.Create(vehicleModel); err != nil {
		var valError *e.ValidationError
		switch {
		case errors.As(err, &valError):
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(valError.Decode())
		case errors.Is(err, e.BadRequestError):
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(vehicleModel); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (vmc vehicleModelController) ListModels(w http.ResponseWriter, r *http.Request) {
	models, err := vmc.vehicleModelService.All()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(models); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
