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
		var body interface{}
		json.Unmarshal([]byte(err.Error()), &body)

		return &h.Response{
			Body:   body,
			Status: http.StatusBadRequest,
		}
	}

	return &h.Response{
		Body:   user,
		Status: http.StatusCreated,
	}
}