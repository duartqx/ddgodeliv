package controllers

import (
	"encoding/json"
	"net/http"

	h "ddgodeliv/api/http"
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

func (uc UserController) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	var response *h.Response

	switch r.Method {
	case http.MethodPost:
		response = uc.create(r)
	case http.MethodGet:
		response = uc.get(r)
	case http.MethodPatch:
		response = uc.updatePassword(r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if body, err := json.Marshal(response.Body); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(response.Status)
		w.Write(body)
	}
}

func (uc UserController) create(r *http.Request) *h.Response {

	user := u.GetNewUser()

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		return &h.Response{
			Body:   map[string]string{"error": "Json Decode Error"},
			Status: http.StatusBadRequest,
		}
	}

	if err := uc.userService.Create(user); err != nil {

		creationErr := err.Error()
		var errs interface{}

		err = json.Unmarshal([]byte(creationErr), &errs)
		if err != nil {
			errs = creationErr
		}

		return &h.Response{
			Body:   map[string]interface{}{"error": errs},
			Status: http.StatusBadRequest,
		}
	}

	return &h.Response{
		Body:   user.Clean(),
		Status: http.StatusCreated,
	}
}

func (uc UserController) get(r *http.Request) *h.Response {
	user, ok := r.Context().Value("user").(*s.ClaimsUser)
	if !ok {
		return &h.Response{Status: http.StatusUnauthorized}
	}
	return &h.Response{
		Body:   user,
		Status: http.StatusOK,
	}
}

func (uc UserController) updatePassword(r *http.Request) *h.Response {

	user, ok := r.Context().Value("user").(*s.ClaimsUser)
	if !ok {
		return &h.Response{Status: http.StatusUnauthorized}
	}

	userToUpdate := u.GetNewUser().SetId(user.Id)

	password := struct {
		Password string `json:"password"`
	}{}

	if err := json.NewDecoder(r.Body).Decode(&password); err != nil {
		return &h.Response{Status: http.StatusBadRequest}
	}

	userToUpdate.SetPassword(password.Password)

	if err := uc.userService.UpdatePassword(userToUpdate); err != nil {
		return &h.Response{Status: http.StatusBadRequest}
	}

	return &h.Response{Status: http.StatusOK}
}

// updateName or updateEmail must also update the jwt
