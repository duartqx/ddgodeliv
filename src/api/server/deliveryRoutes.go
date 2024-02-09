package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"ddgodeliv/api/controllers"
	"ddgodeliv/application/services"
	repository "ddgodeliv/infrastructure/repository/postgresql"
)

func (s server) SetupDeliveryRoutes() http.Handler {

	deliverySubRouter := chi.NewRouter()

	deliverySubRouter.Use(s.jwtController.AuthenticatedMiddleware)

	deliveryRepository := repository.GetNewDeliveryRepository(s.db)
	deliveryService := services.GetNewDeliveryService(
		deliveryRepository, s.driverRepository,
	)
	deliveryController := controllers.GetNewDeliveryController(
		deliveryService, s.sessionService,
	)

	deliverySubRouter.Post("/", deliveryController.Create)

	deliverySubRouter.Get("/", deliveryController.ListAllForSender)

	deliverySubRouter.Get("/{id:[0-9]+}", deliveryController.Get)

	deliverySubRouter.Patch("/{id:[0-9]+}/status", deliveryController.UpdateStatus)

	deliverySubRouter.Patch("/{id:[0-9]+}/assign", deliveryController.AssignDriver)

	deliverySubRouter.Delete("/{id:[0-9]+}", deliveryController.Delete)

	deliverySubRouter.Get("/company", deliveryController.ListByCompany)

	deliverySubRouter.Get("/pending", deliveryController.ListAllPendingsWithoutDriver)

	return deliverySubRouter
}
