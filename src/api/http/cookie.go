package http

import "net/http"

const AuthCookieName = "JwtAuthToken"

func GetInvalidCookie() *http.Cookie {
	return &http.Cookie{Name: AuthCookieName, MaxAge: -1}
}
