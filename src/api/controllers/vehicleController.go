package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	as "ddgodeliv/application/auth"
	e "ddgodeliv/application/errors"
	s "ddgodeliv/application/services"
	v "ddgodeliv/domains/vehicle"

	"github.com/go-chi/chi/v5"
)

type VehicleController struct {
	vehicleModelService *s.VehicleModelService
	vehicleService      *s.VehicleService
	sessionService      *as.SessionService
}

func GetNewVehicleController(
	vehicleModelService *s.VehicleModelService,
	vehicleService *s.VehicleService,
	sessionService *as.SessionService,
) *VehicleController {

	return &VehicleController{
		vehicleModelService: vehicleModelService,
		vehicleService:      vehicleService,
		sessionService:      sessionService,
	}
}

func (vc VehicleController) CreateVehicleModel(w http.ResponseWriter, r *http.Request) {

	vehicleModel := v.GetNewVehicleModel()

	if err := json.NewDecoder(r.Body).Decode(&vehicleModel); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := vc.vehicleModelService.Create(vehicleModel); err != nil {
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

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(vehicleModel); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (vc VehicleController) CreateVehicle(w http.ResponseWriter, r *http.Request) {

	vehicle := v.GetNewVehicle()

	if err := json.NewDecoder(r.Body).Decode(&vehicle); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := vc.vehicleService.Create(vehicle); err != nil {
		switch {
		case errors.Is(err, &e.ValidationError{}):
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(vc.vehicleModelService.Decode(err))
			return
		case errors.Is(err, e.BadRequestError):
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(vehicle); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (vc VehicleController) GetCompanyVehicles(w http.ResponseWriter, r *http.Request) {

	user := vc.sessionService.GetSessionUserWithCompany(r.Context())
	if user == nil {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	vehicles, err := vc.vehicleService.FindByCompanyId(user.GetCompanyId())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(vehicles); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	return
}

func (vc VehicleController) DeleteVehicle(w http.ResponseWriter, r *http.Request) {

	user := vc.sessionService.GetSessionUserWithCompany(r.Context())
	if user == nil {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	vehicleId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	vehicle := v.GetNewVehicle().SetId(vehicleId).SetCompanyId(user.GetCompanyId())

	if err := vc.vehicleService.FindById(vehicle); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if err := vc.vehicleService.Delete(vehicle); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	return
}
