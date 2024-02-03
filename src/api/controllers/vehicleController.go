package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	h "ddgodeliv/api/http"
	as "ddgodeliv/application/auth"
	s "ddgodeliv/application/services"
	v "ddgodeliv/domains/vehicle"

	"github.com/go-chi/chi/v5"
)

type VehicleController struct {
	vehicleModelService *s.VehicleModelService
	vehicleService      *s.VehicleService
	claimsService       *as.ClaimsService
}

func GetNewVehicleController(
	vehicleModelService *s.VehicleModelService,
	vehicleService *s.VehicleService,
	claimsService *as.ClaimsService,
) *VehicleController {

	return &VehicleController{
		vehicleModelService: vehicleModelService,
		vehicleService:      vehicleService,
		claimsService:       claimsService,
	}
}

func (vc VehicleController) CreateVehicleModel(w http.ResponseWriter, r *http.Request) {

	vehicleModel := v.GetNewVehicleModel()

	if err := json.NewDecoder(r.Body).Decode(&vehicleModel); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if validationErrs := vc.vehicleModelService.ValidateStructJson(vehicleModel); validationErrs != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(*validationErrs)
		return
	}

	if err := vc.vehicleModelService.Create(vehicleModel); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
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

	w.Header().Set("Content-Type", "application/json")

	if validationErrs := vc.vehicleService.ValidateStructJson(vehicle); validationErrs != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(*validationErrs)
		return
	}

	if err := vc.vehicleService.Create(vehicle); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(vehicle); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (vc VehicleController) GetCompanyVehicles(w http.ResponseWriter, r *http.Request) {

	user, err := vc.claimsService.GetClaimsUserFromContext(r.Context())
	if err != nil {
		http.SetCookie(w, h.GetInvalidCookie())
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	if err := vc.claimsService.GetWithDriverInfo(user); err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
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

	user, err := vc.claimsService.GetClaimsUserFromContext(r.Context())
	if err != nil {
		http.Error(w, "Not Authorized", http.StatusUnauthorized)
		return
	}
	if err := vc.claimsService.GetWithDriverInfo(user); err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
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
