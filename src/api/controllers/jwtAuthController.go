package controllers

import (
	"context"
	"encoding/json"
	"net/http"

	h "ddgodeliv/api/http"
	a "ddgodeliv/application/auth"
	e "ddgodeliv/application/errors"
	u "ddgodeliv/domains/user"
)

type jwtController struct {
	jwtService *a.JwtAuthService
}

func NewJwtController(jwtService *a.JwtAuthService) *jwtController {
	return &jwtController{
		jwtService: jwtService,
	}
}

func (jc jwtController) Login(w http.ResponseWriter, r *http.Request) {

	user := u.GetNewUser()

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, e.BadRequestError.Error(), http.StatusBadRequest)
		return
	}

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
		Status:    "Valid",
	}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (jc jwtController) Logout(w http.ResponseWriter, r *http.Request) {

	cookie, _ := r.Cookie(h.AuthCookieName)

	err := jc.jwtService.Logout(r.Header.Get("Authorization"), cookie)

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
	}

	http.SetCookie(w, h.GetInvalidCookie())
	w.WriteHeader(http.StatusUnauthorized)
}

func (jc jwtController) AuthenticatedMiddleware(next http.Handler) http.Handler {
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

func (jc jwtController) UnauthenticatedMiddleware(next http.Handler) http.Handler {
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
