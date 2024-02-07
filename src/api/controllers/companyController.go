package controllers

import (
	"encoding/json"
	"errors"
	"net/http"

	s "ddgodeliv/application/services"
	as "ddgodeliv/application/services/auth"
	e "ddgodeliv/common/errors"
	c "ddgodeliv/domains/company"
)

type CompanyController struct {
	companyService *s.CompanyService
	sessionService *as.SessionService
}

func GetNewCompanyController(
	companyService *s.CompanyService, sessionService *as.SessionService,
) *CompanyController {
	return &CompanyController{
		companyService: companyService,
		sessionService: sessionService,
	}
}

func (cc CompanyController) Create(w http.ResponseWriter, r *http.Request) {

	user, err := cc.sessionService.GetSessionUserWithoutCompany(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	body := struct {
		Name      string `json:"name"`
		LicenseId string `json:"license_id"`
	}{}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, e.BadRequestError.Error(), http.StatusBadRequest)
		return
	}

	company := c.GetNewCompany().SetOwnerId(user.GetId()).SetName(body.Name)

	if err := cc.companyService.CreateCompany(company, body.LicenseId); err != nil {
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

	if err := json.NewEncoder(w).Encode(company); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// Sets driver/company info on sessionUser
	go cc.sessionService.SetSessionUserCompany(user)
}

func (cc CompanyController) Delete(w http.ResponseWriter, r *http.Request) {

	user := cc.sessionService.GetSessionUserWithCompany(r.Context())
	if user == nil {
		http.Error(w, e.ForbiddenError.Error(), http.StatusForbidden)
		return
	}

	company := c.GetNewCompany().
		SetId(user.GetCompanyId()).
		SetOwnerId(user.GetId())

	if err := cc.companyService.DeleteCompany(user.GetId(), company); err != nil {
		var valError *e.ValidationError
		switch {
		case errors.As(err, &valError):
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(valError.Decode())
		case errors.Is(err, e.ForbiddenError):
			http.Error(w, err.Error(), http.StatusForbidden)
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	// Removes driver/company information on sessionUser
	go cc.sessionService.SetSessionUserNoCompany(user)
}
