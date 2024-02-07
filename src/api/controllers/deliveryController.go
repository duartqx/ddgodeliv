package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	s "ddgodeliv/application/services"
	a "ddgodeliv/application/services/auth"
	e "ddgodeliv/common/errors"
	"ddgodeliv/domains/company"
	de "ddgodeliv/domains/delivery"
	d "ddgodeliv/domains/driver"
)

type DeliveryController struct {
	deliveryService *s.DeliveryService
	sessionService  *a.SessionService
}

func GetNewDeliveryController(
	deliveryService *s.DeliveryService, sessionService *a.SessionService,
) *DeliveryController {
	return &DeliveryController{
		deliveryService: deliveryService,
		sessionService:  sessionService,
	}
}

// At this first version a user actively creates deliveries, on a more
// realistic scenario the system would create this after an event, then it
// would trigger another event to assign a driver/partner company to start
// the delivery asyncronously
func (dc DeliveryController) Create(w http.ResponseWriter, r *http.Request) {

	user := dc.sessionService.GetSessionUser(r.Context())
	if user == nil {
		http.Error(w, e.ForbiddenError.Error(), http.StatusForbidden)
		return
	}

	delivery := de.GetNewDelivery()

	if err := json.NewDecoder(r.Body).Decode(&delivery); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := dc.deliveryService.Create(delivery); err != nil {
		var valError *e.ValidationError
		switch {
		case errors.As(err, &valError):
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(valError.Decode())
		case errors.Is(err, e.BadRequestError):
			http.Error(w, err.Error(), http.StatusBadRequest)
		case errors.Is(err, e.NotFoundError):
			http.Error(w, err.Error(), http.StatusNotFound)
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(delivery); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (dc DeliveryController) AssignDriver(w http.ResponseWriter, r *http.Request) {
	user := dc.sessionService.GetSessionUserWithCompany(r.Context())
	if user == nil || user.HasInvalidCompany() {
		http.Error(w, e.ForbiddenError.Error(), http.StatusForbidden)
		return
	}

	deliveryId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, e.BadRequestError.Error(), http.StatusBadRequest)
		return
	}

	body := struct {
		DriverId int `json:"driver_id"`
	}{}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, e.BadRequestError.Error(), http.StatusBadRequest)
		return
	}

	driver := d.GetNewDriver().
		SetId(body.DriverId).
		SetCompanyId(user.GetCompanyId())

	delivery := de.GetNewDelivery().SetId(deliveryId)

	if err := dc.deliveryService.AssignDriver(delivery, driver); err != nil {
		switch {
		case errors.Is(err, e.BadRequestError):
			http.Error(w, err.Error(), http.StatusBadRequest)
		case errors.Is(err, e.NotFoundError):
			http.Error(w, err.Error(), http.StatusNotFound)
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(delivery); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (dc DeliveryController) UpdateStatus(w http.ResponseWriter, r *http.Request) {
	user := dc.sessionService.GetSessionUserWithCompany(r.Context())
	if user == nil || user.HasInvalidCompany() {
		http.Error(w, e.ForbiddenError.Error(), http.StatusForbidden)
		return
	}

	deliveryId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, e.BadRequestError.Error(), http.StatusBadRequest)
		return
	}

	body := struct {
		Status uint8 `json:"status"`
	}{}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, e.BadRequestError.Error(), http.StatusBadRequest)
		return
	}

	delivery := de.GetNewDelivery().SetId(deliveryId).SetStatus(body.Status)

	if err := dc.deliveryService.UpdateStatus(delivery); err != nil {
		var valError *e.ValidationError
		switch {
		case errors.As(err, &valError):
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(valError.Decode())
		case errors.Is(err, e.BadRequestError):
			http.Error(w, err.Error(), http.StatusBadRequest)
		case errors.Is(err, e.NotFoundError):
			http.Error(w, err.Error(), http.StatusNotFound)
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(delivery); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (dc DeliveryController) Get(w http.ResponseWriter, r *http.Request) {
	user := dc.sessionService.GetSessionUserWithCompany(r.Context())
	if user == nil {
		http.Error(w, e.ForbiddenError.Error(), http.StatusForbidden)
		return
	}

	deliveryId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, e.BadRequestError.Error(), http.StatusBadRequest)
		return
	}

	delivery := de.GetNewDelivery().SetId(deliveryId)

	if err := dc.deliveryService.FindById(user.ToUser(), delivery); err != nil {
		switch {
		case errors.Is(err, e.BadRequestError):
			http.Error(w, err.Error(), http.StatusBadRequest)
		case errors.Is(err, e.NotFoundError):
			http.Error(w, err.Error(), http.StatusNotFound)
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(delivery); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (dc DeliveryController) ListByCompany(w http.ResponseWriter, r *http.Request) {
	user := dc.sessionService.GetSessionUserWithCompany(r.Context())
	if user == nil || user.HasInvalidCompany() {
		http.Error(w, e.ForbiddenError.Error(), http.StatusForbidden)
		return
	}

	company := company.GetNewCompany().SetId(user.GetCompanyId())

	deliveries, err := dc.deliveryService.FindByCompanyId(company)
	if err != nil {
		switch {
		case errors.Is(err, e.BadRequestError):
			http.Error(w, err.Error(), http.StatusBadRequest)
		case errors.Is(err, e.NotFoundError):
			http.Error(w, err.Error(), http.StatusNotFound)
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(deliveries); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (dc DeliveryController) Delete(w http.ResponseWriter, r *http.Request) {
	user := dc.sessionService.GetSessionUserWithCompany(r.Context())
	if user == nil {
		http.Error(w, e.ForbiddenError.Error(), http.StatusForbidden)
		return
	}

	deliveryId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, e.BadRequestError.Error(), http.StatusBadRequest)
		return
	}

	delivery := de.GetNewDelivery().SetId(deliveryId)

	if err := dc.deliveryService.Delete(user.ToUser(), delivery); err != nil {
		switch {
		case errors.Is(err, e.BadRequestError):
			http.Error(w, err.Error(), http.StatusBadRequest)
		case errors.Is(err, e.NotFoundError):
			http.Error(w, err.Error(), http.StatusNotFound)
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(delivery); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (dc DeliveryController) ListAllPendingsWithoutDriver(w http.ResponseWriter, r *http.Request) {

	user := dc.sessionService.GetSessionUser(r.Context())
	if user == nil {
		http.Error(w, e.ForbiddenError.Error(), http.StatusForbidden)
		return
	}

	deliveries, err := dc.deliveryService.FindPendingWithoutDriver()
	if err != nil {
		switch {
		case errors.Is(err, e.BadRequestError):
			http.Error(w, err.Error(), http.StatusBadRequest)
		case errors.Is(err, e.NotFoundError):
			http.Error(w, err.Error(), http.StatusNotFound)
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(deliveries); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (dc DeliveryController) ListAllForSender(w http.ResponseWriter, r *http.Request) {
	user := dc.sessionService.GetSessionUser(r.Context())
	if user == nil {
		http.Error(w, e.ForbiddenError.Error(), http.StatusForbidden)
		return
	}

	deliveries, err := dc.deliveryService.FindBySenderId(user.ToUser())
	if err != nil {
		switch {
		case errors.Is(err, e.BadRequestError):
			http.Error(w, err.Error(), http.StatusBadRequest)
		case errors.Is(err, e.NotFoundError):
			http.Error(w, err.Error(), http.StatusNotFound)
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(deliveries); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
