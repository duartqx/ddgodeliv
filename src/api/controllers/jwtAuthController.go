package controllers

import (
	"context"
	"encoding/json"
	"net/http"

	h "ddgodeliv/api/http"
	s "ddgodeliv/application/services"
	u "ddgodeliv/domains/user"
)

type JwtController struct {
	jwtService *s.JwtAuthService
	cookieName string
}

func NewJwtController(jwtService *s.JwtAuthService) *JwtController {
	return &JwtController{
		jwtService: jwtService,
		cookieName: "JwtAuthToken",
	}
}

func (jc JwtController) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	var response *h.Response

	switch r.Method {
	case http.MethodPost:
		response = jc.Login(r)
	case http.MethodDelete:
		response = jc.Logout(r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if body, err := json.Marshal(response.Body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	} else {
		if response.Cookie != nil {
			http.SetCookie(w, response.Cookie)
		}
		w.WriteHeader(response.Status)
		w.Write(body)
	}
}

func (jc JwtController) Login(r *http.Request) *h.Response {

	user := u.GetNewUser()

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		return &h.Response{
			Status: http.StatusBadRequest,
			Body:   "Bad Request",
		}
	}

	token, expiresAt, err := jc.jwtService.Login(user)

	if err != nil {
		return &h.Response{
			Status: http.StatusUnauthorized,
			Body:   "Unauthorized",
		}
	}

	return &h.Response{
		Status: http.StatusOK,
		Body: h.LoginResponse{
			Token:     token,
			ExpiresAt: expiresAt,
			Status:    "Valid",
		},
		Cookie: &http.Cookie{
			Name:     jc.cookieName,
			Value:    token,
			Expires:  expiresAt,
			Secure:   true,
			HttpOnly: true,
		},
	}
}

func (jc JwtController) Logout(r *http.Request) *h.Response {

	cookie, _ := r.Cookie(jc.cookieName)

	err := jc.jwtService.Logout(r.Header.Get("Authorization"), cookie)

	if err != nil {
		return &h.Response{
			Status: http.StatusUnauthorized,
			Body:   err.Error(),
		}
	}

	return &h.Response{
		Status: http.StatusOK,
		Body:   "Logged out",
		Cookie: &http.Cookie{
			Name:   jc.cookieName,
			MaxAge: -1,
		},
	}
}

func (jc JwtController) AuthenticatedMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		cookie, _ := r.Cookie(jc.cookieName)
		claimsUser, err := jc.jwtService.ValidateAuth(r.Header.Get("Authorization"), cookie)

		if err != nil {
			http.SetCookie(w, &http.Cookie{Name: jc.cookieName, MaxAge: -1})
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// Injects the User Information into the request context
		ctx := context.WithValue(r.Context(), "user", claimsUser)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (jc JwtController) UnauthenticatedMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		cookie, _ := r.Cookie(jc.cookieName)
		claimsUser, _ := jc.jwtService.ValidateAuth(r.Header.Get("Authorization"), cookie)

		if claimsUser != nil {
			http.SetCookie(w, &http.Cookie{Name: jc.cookieName, MaxAge: -1})
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
