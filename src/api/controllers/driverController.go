package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	s "ddgodeliv/application/services"
	as "ddgodeliv/application/services/auth"
	e "ddgodeliv/common/errors"
	d "ddgodeliv/domains/driver"
	u "ddgodeliv/domains/user"
)

type DriverController struct {
	driverService  *s.DriverService
	sessionService *as.SessionService
}

func GetNewDriverController(
	driverService *s.DriverService, sessionService *as.SessionService,
) *DriverController {
	return &DriverController{
		driverService:  driverService,
		sessionService: sessionService,
	}
}

func (dc DriverController) Create(w http.ResponseWriter, r *http.Request) {

	user := dc.sessionService.GetSessionUserWithCompany(r.Context())
	if user == nil {
		http.Error(w, e.ForbiddenError.Error(), http.StatusForbidden)
		return
	}

	body := struct {
		LicenseId string `json:"license_id"`
		User      struct {
			Name  string `json:"name"`
			Email string `json:"email"`
		} `json:"user"`
	}{}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, e.BadRequestError.Error(), http.StatusBadRequest)
		return
	}

	driver := d.GetNewDriver().
		SetUser(
			u.GetNewUser().
				SetEmail(body.User.Email).
				SetName(body.User.Name),
		).
		SetLicenseId(body.LicenseId).
		SetCompanyId(user.GetCompanyId())

	if err := dc.driverService.CreateDriver(driver); err != nil {
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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(driver); err != nil {
		http.Error(w, e.InternalError.Error(), http.StatusInternalServerError)
		return
	}
}

func (dc DriverController) Delete(w http.ResponseWriter, r *http.Request) {

	user := dc.sessionService.GetSessionUserWithCompany(r.Context())
	if user == nil {
		http.Error(w, e.ForbiddenError.Error(), http.StatusForbidden)
		return
	}

	driverId, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, e.BadRequestError.Error(), http.StatusBadRequest)
		return
	}

	driver := d.GetNewDriver().SetId(driverId).SetCompanyId(user.GetCompanyId())

	if err := dc.driverService.FindById(driver); err != nil {
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

func (dc DriverController) Update(w http.ResponseWriter, r *http.Request) {
	user := dc.sessionService.GetSessionUserWithCompany(r.Context())
	if user == nil {
		http.Error(w, e.ForbiddenError.Error(), http.StatusForbidden)
		return
	}

	driverId, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, e.BadRequestError.Error(), http.StatusBadRequest)
		return
	}

	driver := d.GetNewDriver().SetId(driverId).SetCompanyId(user.GetCompanyId())

	if err := dc.driverService.FindById(driver); err != nil {
		switch {
		case errors.Is(err, e.NotFoundError):
			http.Error(w, e.NotFoundError.Error(), http.StatusNotFound)
		default:
			http.Error(w, e.InternalError.Error(), http.StatusInternalServerError)
		}
		return
	}

	body := struct {
		LicenseId string `json:"license_id"`
	}{}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid License Id", http.StatusBadRequest)
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

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(driver); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (dc DriverController) ListCompanyDrivers(w http.ResponseWriter, r *http.Request) {

	user := dc.sessionService.GetSessionUserWithCompany(r.Context())
	if user == nil {
		http.Error(w, e.ForbiddenError.Error(), http.StatusForbidden)
		return
	}

	drivers, err := dc.driverService.ListCompanyDriversById(user.GetCompanyId())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(drivers); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (dc DriverController) Get(w http.ResponseWriter, r *http.Request) {
	user := dc.sessionService.GetSessionUserWithCompany(r.Context())
	if user == nil {
		http.Error(w, e.ForbiddenError.Error(), http.StatusForbidden)
		return
	}

	driverId, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, e.BadRequestError.Error(), http.StatusBadRequest)
		return
	}

	driver := d.GetNewDriver().SetId(driverId).SetCompanyId(user.GetCompanyId())

	if err := dc.driverService.FindById(driver); err != nil {
		switch {
		case errors.Is(err, e.NotFoundError):
			http.Error(w, e.NotFoundError.Error(), http.StatusNotFound)
		default:
			http.Error(w, e.InternalError.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(driver); err != nil {
		http.Error(w, e.InternalError.Error(), http.StatusInternalServerError)
		return
	}
}
