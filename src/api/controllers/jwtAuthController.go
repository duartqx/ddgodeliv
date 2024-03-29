package controllers

import (
	"context"
	"encoding/json"
	"net/http"

	h "ddgodeliv/api/http"
	a "ddgodeliv/application/services/auth"
	e "ddgodeliv/common/errors"
	u "ddgodeliv/domains/user"
)

type JwtController struct {
	jwtService *a.JwtAuthService
}

var jwtController *JwtController

func GetJwtController(jwtService *a.JwtAuthService) *JwtController {
	if jwtController == nil {
		jwtController = &JwtController{
			jwtService: jwtService,
		}
	}
	return jwtController
}

func (jc JwtController) Login(w http.ResponseWriter, r *http.Request) {

	body := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, e.BadRequestError.Error(), http.StatusBadRequest)
		return
	}

	user := u.GetNewUser().SetEmail(body.Email).SetPassword(body.Password)

	token, expiresAt, err := jc.jwtService.Login(user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     h.AuthCookieName,
		Value:    token,
		Expires:  expiresAt,
		Secure:   true,
		HttpOnly: true,
	})

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(h.LoginResponse{
		Token:     token,
		ExpiresAt: expiresAt,
		Status:    true,
	}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (jc JwtController) Logout(w http.ResponseWriter, r *http.Request) {

	cookie, _ := r.Cookie(h.AuthCookieName)

	err := jc.jwtService.Logout(r.Header.Get("Authorization"), cookie)

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
	}

	http.SetCookie(w, h.GetInvalidCookie())
}

func (jc JwtController) AuthenticatedMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		cookie, _ := r.Cookie(h.AuthCookieName)
		claimsUser, err := jc.jwtService.ValidateAuth(r.Header.Get("Authorization"), cookie)

		if err != nil {
			http.SetCookie(w, h.GetInvalidCookie())
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

		cookie, _ := r.Cookie(h.AuthCookieName)
		claimsUser, _ := jc.jwtService.ValidateAuth(r.Header.Get("Authorization"), cookie)

		if claimsUser != nil {
			http.SetCookie(w, h.GetInvalidCookie())
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
