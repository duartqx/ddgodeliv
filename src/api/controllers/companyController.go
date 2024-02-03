package controllers

import (
	"encoding/json"
	"net/http"

	h "ddgodeliv/api/http"
	as "ddgodeliv/application/auth"
	s "ddgodeliv/application/services"
	c "ddgodeliv/domains/company"
)

type CompanyController struct {
	companyService *s.CompanyService
	claimsService  *as.ClaimsService
}

func GetNewCompanyController(
	companyService *s.CompanyService, claimsService *as.ClaimsService,
) *CompanyController {
	return &CompanyController{
		companyService: companyService,
		claimsService:  claimsService,
	}
}

func (cc CompanyController) Create(w http.ResponseWriter, r *http.Request) {

	user, err := cc.claimsService.GetClaimsUserFromContext(r.Context())
	if err != nil {
		http.SetCookie(w, h.GetInvalidCookie())
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if !user.HasInvalidCompany() {
		http.Error(w, "User Already is part of a Company!", http.StatusBadRequest)
		return
	}

	tmpComp := struct {
		Name      string `json:"name"`
		LicenseId string `json:"license_id"`
	}{}
	if err := json.NewDecoder(r.Body).Decode(&tmpComp); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := cc.companyService.ValidateVarJson(
		tmpComp.LicenseId, "required,min=3,max=250",
	); err != nil {
		http.Error(w, "Invalid Driver License", http.StatusBadRequest)
		return
	}

	company := c.GetNewCompany().SetOwnerId(user.GetId()).SetName(tmpComp.Name)

	if validationsErrs := cc.companyService.ValidateStructJson(company); validationsErrs != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(*validationsErrs)
		return
	}

	if err := cc.companyService.CreateCompany(company, tmpComp.LicenseId); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(company); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (cc CompanyController) Delete(w http.ResponseWriter, r *http.Request) {

	user, err := cc.claimsService.GetClaimsUserFromContext(r.Context())
	if err != nil {
		http.SetCookie(w, h.GetInvalidCookie())
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	if err := cc.claimsService.GetWithDriverInfo(user); err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	if err := cc.companyService.DeleteCompany(
		user.GetId(), c.GetNewCompany().SetId(user.GetCompanyId()),
	); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

}
