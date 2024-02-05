package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	e "ddgodeliv/application/errors"
	s "ddgodeliv/application/services"
	as "ddgodeliv/application/services/auth"
	d "ddgodeliv/domains/driver"

	"github.com/go-chi/chi/v5"
)

type driverController struct {
	driverService  *s.DriverService
	sessionService *as.SessionService
}

func GetNewDriverController(
	driverService *s.DriverService, sessionService *as.SessionService,
) *driverController {
	return &driverController{
		driverService:  driverService,
		sessionService: sessionService,
	}
}

func (dc driverController) Create(w http.ResponseWriter, r *http.Request) {

	user := dc.sessionService.GetSessionUserWithCompany(r.Context())
	if user == nil {
		http.Error(w, e.ForbiddenError.Error(), http.StatusForbidden)
		return
	}

	driver := d.GetNewDriver()

	if err := json.NewDecoder(r.Body).Decode(&driver); err != nil {
		http.Error(w, e.BadRequestError.Error(), http.StatusBadRequest)
		return
	}

	if err := dc.driverService.CreateDriver(driver); err != nil {
		var valError *e.ValidationError
		switch {
		case errors.As(err, &valError):
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(valError.Decode())
		case errors.Is(err, e.ForbiddenError):
			http.Error(w, err.Error(), http.StatusForbidden)
			return
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
	if err := json.NewEncoder(w).Encode(driver); err != nil {
		http.Error(w, e.InternalError.Error(), http.StatusInternalServerError)
		return
	}
}

func (dc driverController) Delete(w http.ResponseWriter, r *http.Request) {

	user := dc.sessionService.GetSessionUserWithCompany(r.Context())
	if user == nil {
		http.Error(w, e.ForbiddenError.Error(), http.StatusForbidden)
		return
	}

	driverId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, e.BadRequestError.Error(), http.StatusBadRequest)
		return
	}

	driver, err := dc.driverService.FindById(driverId, user.GetCompanyId())
	if err != nil {
		http.Error(w, e.NotFoundError.Error(), http.StatusNotFound)
		return
	}

	if err := dc.driverService.DeleteDriver(user.GetId(), driver); err != nil {
		switch {
		case errors.Is(err, e.ForbiddenError):
			http.Error(w, err.Error(), http.StatusForbidden)
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
}

func (dc driverController) Update(w http.ResponseWriter, r *http.Request) {
	user := dc.sessionService.GetSessionUserWithCompany(r.Context())
	if user == nil {
		http.Error(w, e.ForbiddenError.Error(), http.StatusForbidden)
		return
	}

	driverId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, e.BadRequestError.Error(), http.StatusBadRequest)
		return
	}

	driver, err := dc.driverService.FindById(driverId, user.GetCompanyId())
	if err != nil {
		http.Error(w, e.NotFoundError.Error(), http.StatusNotFound)
		return
	}

	body := struct {
		LicenseId string `json:"license_id"`
	}{}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, e.BadRequestError.Error(), http.StatusBadRequest)
		return
	}

	if err := dc.driverService.UpdateDriverLicense(
		user.GetId(), driver.SetLicenseId(body.LicenseId),
	); err != nil {
		var valError *e.ValidationError
		switch {
		case errors.As(err, &valError):
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(valError.Decode())
		case errors.Is(err, e.ForbiddenError):
			http.Error(w, err.Error(), http.StatusForbidden)
		case errors.Is(err, e.BadRequestError):
			http.Error(w, err.Error(), http.StatusBadRequest)
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	if err := json.NewEncoder(w).Encode(driver); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
}

func (dc driverController) ListCompanyDrivers(w http.ResponseWriter, r *http.Request) {

	user := dc.sessionService.GetSessionUserWithCompany(r.Context())
	if user == nil {
		http.Error(w, e.ForbiddenError.Error(), http.StatusForbidden)
		return
	}

	drivers, err := dc.driverService.ListCompanyDriversById(user.GetCompanyId())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if err := json.NewEncoder(w).Encode(drivers); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
}