package http

import (
	"time"

	h "net/http"
)

type Response struct {
	Status int
	Body   interface{}
	Cookie *h.Cookie
}

type LoginResponse struct {
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expiresAt"`
	Status    string    `json:"status"`
}
