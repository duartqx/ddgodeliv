package http

import (
	"encoding/json"
	"errors"
	"net/http"

	e "ddgodeliv/common/errors"
)

func ErrorResponse(w http.ResponseWriter, err error) {
	var valError *e.ValidationError
	switch {
	case errors.As(err, &valError):
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(valError.Decode())
	case errors.Is(err, e.NotFoundError):
		http.Error(w, err.Error(), http.StatusNotFound)
	case errors.Is(err, e.BadRequestError):
		http.Error(w, err.Error(), http.StatusBadRequest)
	case errors.Is(err, e.ForbiddenError):
		http.Error(w, err.Error(), http.StatusForbidden)
	default:
		panic(err.Error())
	}
}
