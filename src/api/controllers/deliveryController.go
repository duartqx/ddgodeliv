package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	h "ddgodeliv/api/http"
	s "ddgodeliv/application/services"
	a "ddgodeliv/application/services/auth"
	e "ddgodeliv/common/errors"
	"ddgodeliv/domains/company"
	de "ddgodeliv/domains/delivery"
	d "ddgodeliv/domains/driver"
)

type DeliveryController struct {
	deliveryService *s.DeliveryService
	driverService   *s.DriverService
	sessionService  *a.SessionService
}

func GetDeliveryController(
	deliveryService *s.DeliveryService,
	driverService *s.DriverService,
	sessionService *a.SessionService,
) *DeliveryController {
	return &DeliveryController{
		deliveryService: deliveryService,
		driverService:   driverService,
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

	body := struct {
		Loadout     string    `json:"loadout"`
		Weight      int       `json:"weight"`
		Origin      string    `json:"origin"`
		Destination string    `json:"destination"`
		Deadline    time.Time `json:"deadline"`
	}{}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	delivery := de.GetNewDelivery().
		SetSenderId(user.GetId()).
		SetStatus(de.StatusChoices.Pending).
		SetLoadout(body.Loadout).
		SetWeight(body.Weight).
		SetOrigin(body.Origin).
		SetDestination(body.Destination).
		SetDeadline(body.Deadline)

	if err := dc.deliveryService.Create(delivery); err != nil {
		h.ErrorResponse(w, err)
		return
	}

	h.JsonResponse(w, http.StatusCreated, delivery)
}

func (dc DeliveryController) AssignDriver(w http.ResponseWriter, r *http.Request) {
	user := dc.sessionService.GetSessionUserWithCompany(r.Context())
	if user == nil || user.HasInvalidCompany() {
		http.Error(w, e.ForbiddenError.Error(), http.StatusForbidden)
		return
	}

	deliveryId, err := strconv.Atoi(r.PathValue("id"))
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
		h.ErrorResponse(w, err)
		return
	}

	h.JsonResponse(w, http.StatusOK, delivery)
}

func (dc DeliveryController) UpdateStatus(w http.ResponseWriter, r *http.Request) {
	user := dc.sessionService.GetSessionUserWithCompany(r.Context())
	if user == nil || user.HasInvalidCompany() {
		http.Error(w, e.ForbiddenError.Error(), http.StatusForbidden)
		return
	}

	deliveryId, err := strconv.Atoi(r.PathValue("id"))
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
		h.ErrorResponse(w, err)
		return
	}

	h.JsonResponse(w, http.StatusOK, delivery)
}

func (dc DeliveryController) Get(w http.ResponseWriter, r *http.Request) {
	user := dc.sessionService.GetSessionUser(r.Context())
	if user == nil {
		http.Error(w, e.ForbiddenError.Error(), http.StatusForbidden)
		return
	}

	deliveryId, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, e.BadRequestError.Error(), http.StatusBadRequest)
		return
	}

	delivery := de.GetNewDelivery().SetId(deliveryId)

	if err := dc.deliveryService.FindById(user.ToUser(), delivery); err != nil {
		h.ErrorResponse(w, err)
		return
	}

	h.JsonResponse(w, http.StatusOK, delivery)
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
		h.ErrorResponse(w, err)
		return
	}

	h.JsonResponse(w, http.StatusOK, deliveries)
}

func (dc DeliveryController) Delete(w http.ResponseWriter, r *http.Request) {
	user := dc.sessionService.GetSessionUserWithCompany(r.Context())
	if user == nil {
		http.Error(w, e.ForbiddenError.Error(), http.StatusForbidden)
		return
	}

	deliveryId, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, e.BadRequestError.Error(), http.StatusBadRequest)
		return
	}

	delivery := de.GetNewDelivery().SetId(deliveryId)

	if err := dc.deliveryService.Delete(user.ToUser(), delivery); err != nil {
		h.ErrorResponse(w, err)
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
		h.ErrorResponse(w, err)
		return
	}

	h.JsonResponse(w, http.StatusOK, deliveries)
}

func (dc DeliveryController) ListAllForSender(w http.ResponseWriter, r *http.Request) {
	user := dc.sessionService.GetSessionUser(r.Context())
	if user == nil {
		http.Error(w, e.ForbiddenError.Error(), http.StatusForbidden)
		return
	}

	deliveries, err := dc.deliveryService.FindBySenderId(user.ToUser())
	if err != nil {
		h.ErrorResponse(w, err)
		return
	}

	h.JsonResponse(w, http.StatusOK, deliveries)
}

func (dc DeliveryController) ListAllForDriver(w http.ResponseWriter, r *http.Request) {
	user := dc.sessionService.GetSessionUserWithCompany(r.Context())
	if user == nil {
		http.Error(w, e.ForbiddenError.Error(), http.StatusForbidden)
		return
	}

	driverId, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, e.BadRequestError.Error(), http.StatusBadRequest)
		return
	}

	driver := d.GetNewDriver().SetId(driverId).SetCompanyId(user.GetCompanyId())

	if err := dc.driverService.FindById(driver); err != nil {
		h.ErrorResponse(w, err)
		return
	}

	deliveries, err := dc.deliveryService.FindByDriverId(user.ToUser(), driver)
	if err != nil {
		h.ErrorResponse(w, err)
		return
	}

	h.JsonResponse(w, http.StatusOK, deliveries)
}

func (dc DeliveryController) GetDriverCurrentDelivery(w http.ResponseWriter, r *http.Request) {
	user := dc.sessionService.GetSessionUserWithCompany(r.Context())
	if user == nil {
		http.Error(w, e.ForbiddenError.Error(), http.StatusForbidden)
		return
	}

	driverId, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, e.BadRequestError.Error(), http.StatusBadRequest)
		return
	}

	driver := d.GetNewDriver().SetId(driverId).SetCompanyId(user.GetCompanyId())

	if err := dc.driverService.FindById(driver); err != nil {
		h.ErrorResponse(w, err)
		return
	}

	delivery, err := dc.deliveryService.FindCurrentByDriverId(user.ToUser(), driver)
	if err != nil {
		h.ErrorResponse(w, err)
		return
	}

	h.JsonResponse(w, http.StatusOK, delivery)
}
