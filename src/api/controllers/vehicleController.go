package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	h "ddgodeliv/api/http"
	s "ddgodeliv/application/services"
	as "ddgodeliv/application/services/auth"
	e "ddgodeliv/common/errors"
	v "ddgodeliv/domains/vehicle"
)

type VehicleController struct {
	vehicleService *s.VehicleService
	sessionService *as.SessionService
}

var vehicleController *VehicleController

func GetVehicleController(
	vehicleService *s.VehicleService,
	sessionService *as.SessionService,
) *VehicleController {
	if vehicleController == nil {
		vehicleController = &VehicleController{
			vehicleService: vehicleService,
			sessionService: sessionService,
		}
	}
	return vehicleController
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
		h.ErrorResponse(w, err)
		return
	}

	h.JsonResponse(w, http.StatusCreated, vehicle)
}

func (vc VehicleController) GetCompanyVehicles(w http.ResponseWriter, r *http.Request) {

	user := vc.sessionService.GetSessionUserWithCompany(r.Context())
	if user == nil {
		http.Error(w, e.ForbiddenError.Error(), http.StatusForbidden)
		return
	}

	vehicles, err := vc.vehicleService.FindByCompanyId(user.GetCompanyId())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	h.JsonResponse(w, http.StatusOK, vehicles)
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
		h.ErrorResponse(w, err)
		return
	}

	h.JsonResponse(w, http.StatusOK, vehicle)
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
		h.ErrorResponse(w, err)
		return
	}

	if err := vc.vehicleService.Delete(vehicle); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
