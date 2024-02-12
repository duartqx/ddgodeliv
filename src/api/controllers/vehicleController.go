package controllers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	s "ddgodeliv/application/services"
	as "ddgodeliv/application/services/auth"
	e "ddgodeliv/common/errors"
	v "ddgodeliv/domains/vehicle"
)

type VehicleController struct {
	vehicleService *s.VehicleService
	sessionService *as.SessionService
}

func GetNewVehicleController(
	vehicleService *s.VehicleService,
	sessionService *as.SessionService,
) *VehicleController {

	return &VehicleController{
		vehicleService: vehicleService,
		sessionService: sessionService,
	}
}

func (vc VehicleController) CreateVehicle(w http.ResponseWriter, r *http.Request) {

	user := vc.sessionService.GetSessionUserWithCompany(r.Context())
	if user == nil {
		http.Error(w, e.ForbiddenError.Error(), http.StatusForbidden)
		return
	}

	body := struct {
		LicenseId string `json:"license_id"`
		ModelId   int    `json:"model_id"`
	}{}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, e.BadRequestError.Error(), http.StatusBadRequest)
		return
	}

	vehicle := v.GetNewVehicle().
		SetCompanyId(user.GetCompanyId()).
		SetLicenseId(body.LicenseId).
		SetModelId(body.ModelId)

	if err := vc.vehicleService.Create(vehicle); err != nil {
		var valError *e.ValidationError
		switch {
		case errors.As(err, &valError):
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(valError.Decode())
		case errors.Is(err, e.BadRequestError):
			http.Error(w, err.Error(), http.StatusBadRequest)
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(vehicle); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (vc VehicleController) GetCompanyVehicles(w http.ResponseWriter, r *http.Request) {

	user := vc.sessionService.GetSessionUserWithCompany(r.Context())
	if user == nil {
		http.Error(w, e.ForbiddenError.Error(), http.StatusForbidden)
		return
	}

	vehicles, err := vc.vehicleService.FindByCompanyId(user.GetCompanyId())
	if err != nil {
		log.Println(vehicles, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(vehicles); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (vc VehicleController) GetVehicle(w http.ResponseWriter, r *http.Request) {
	user := vc.sessionService.GetSessionUserWithCompany(r.Context())
	if user == nil {
		http.Error(w, e.ForbiddenError.Error(), http.StatusForbidden)
		return
	}

	vehicleId, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, e.BadRequestError.Error(), http.StatusBadRequest)
		return
	}

	vehicle := v.GetNewVehicle().SetId(vehicleId).SetCompanyId(user.GetCompanyId())

	if err := vc.vehicleService.FindById(vehicle); err != nil {
		switch {
		case errors.Is(err, e.NotFoundError):
			http.Error(w, e.NotFoundError.Error(), http.StatusNotFound)
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(vehicle); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (vc VehicleController) DeleteVehicle(w http.ResponseWriter, r *http.Request) {

	user := vc.sessionService.GetSessionUserWithCompany(r.Context())
	if user == nil {
		http.Error(w, e.ForbiddenError.Error(), http.StatusForbidden)
		return
	}

	vehicleId, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, e.BadRequestError.Error(), http.StatusBadRequest)
		return
	}

	vehicle := v.GetNewVehicle().SetId(vehicleId).SetCompanyId(user.GetCompanyId())

	if err := vc.vehicleService.FindById(vehicle); err != nil {
		switch {
		case errors.Is(err, e.NotFoundError):
			http.Error(w, e.NotFoundError.Error(), http.StatusNotFound)
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	if err := vc.vehicleService.Delete(vehicle); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
