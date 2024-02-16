package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	h "ddgodeliv/api/http"
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

var driverController *DriverController

func GetDriverController(
	driverService *s.DriverService, sessionService *as.SessionService,
) *DriverController {
	if driverController == nil {
		driverController = &DriverController{
			driverService:  driverService,
			sessionService: sessionService,
		}
	}
	return driverController
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
		h.ErrorResponse(w, err)
		return
	}

	h.JsonResponse(w, http.StatusCreated, driver)
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
		h.ErrorResponse(w, err)
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
		h.ErrorResponse(w, err)
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
		h.ErrorResponse(w, err)
		return
	}

	h.JsonResponse(w, http.StatusOK, driver)
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

	h.JsonResponse(w, http.StatusOK, drivers)
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
		h.ErrorResponse(w, err)
		return
	}

	h.JsonResponse(w, http.StatusOK, driver)
}
