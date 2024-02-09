package controllers

import (
	"encoding/json"
	"errors"
	"net/http"

	s "ddgodeliv/application/services"
	as "ddgodeliv/application/services/auth"
	e "ddgodeliv/common/errors"
	u "ddgodeliv/domains/user"
)

type UserController struct {
	userService    *s.UserService
	sessionService *as.SessionService
}

func GetNewUserController(
	userService *s.UserService, sessionService *as.SessionService,
) *UserController {
	return &UserController{
		userService: userService, sessionService: sessionService,
	}
}

func (uc UserController) Create(w http.ResponseWriter, r *http.Request) {

	body := struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, e.BadRequestError.Error(), http.StatusBadRequest)
		return
	}

	user := u.GetNewUser().
		SetName(body.Name).
		SetEmail(body.Email).
		SetPassword(body.Password)

	if err := uc.userService.Create(user); err != nil {
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
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (uc UserController) Get(w http.ResponseWriter, r *http.Request) {

	user := uc.sessionService.GetSessionUserWithCompany(r.Context())
	if user == nil {
		http.Error(w, e.ForbiddenError.Error(), http.StatusForbidden)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (uc UserController) UpdatePassword(w http.ResponseWriter, r *http.Request) {

	user := uc.sessionService.GetSessionUser(r.Context())
	if user == nil {
		http.Error(w, e.ForbiddenError.Error(), http.StatusForbidden)
		return
	}

	p := struct {
		Password string `json:"password"`
	}{}

	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userToUpdate := u.GetNewUser().SetId(user.GetId()).SetPassword(p.Password)

	if err := uc.userService.UpdatePassword(userToUpdate); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

// updateName or updateEmail must also update the jwt
