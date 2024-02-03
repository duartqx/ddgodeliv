package controllers

import (
	"encoding/json"
	"net/http"

	s "ddgodeliv/application/services"
	u "ddgodeliv/domains/user"
)

type UserController struct {
	userService *s.UserService
}

func GetNewUserController(userService *s.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (uc UserController) Create(w http.ResponseWriter, r *http.Request) {

	user := u.GetNewUser()

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if validationErrs := uc.userService.ValidateStructJson(user); validationErrs != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(*validationErrs)
		return
	}

	if err := uc.userService.Create(user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user.Clean())
	return
}

func (uc UserController) Get(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value("user").(*s.ClaimsUser)
	if !ok {
		http.Error(w, "Not Authorized", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(
		u.GetNewUser().
			SetId(user.Id).
			SetEmail(user.Name).
			SetName(user.Name).
			Clean(),
	)
}

func (uc UserController) UpdatePassword(w http.ResponseWriter, r *http.Request) {

	user, ok := r.Context().Value("user").(*s.ClaimsUser)
	if !ok {
		http.Error(w, "Not Authorized", http.StatusUnauthorized)
		return
	}

	userToUpdate := u.GetNewUser().SetId(user.Id)

	password := struct {
		Password string `json:"password"`
	}{}

	if err := json.NewDecoder(r.Body).Decode(&password); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userToUpdate.SetPassword(password.Password)

	if err := uc.userService.UpdatePassword(userToUpdate); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	return
}

// updateName or updateEmail must also update the jwt
