package server

import (
	"net/http"

	"ddgodeliv/api/controllers"
	"ddgodeliv/application/services"
	repository "ddgodeliv/infrastructure/repository/postgresql"
)

func (s server) SetupDeliveryRoutes() http.Handler {

	deliverySubRouter := http.NewServeMux()

	deliveryRepository := repository.GetNewDeliveryRepository(s.db)
	deliveryService := services.GetNewDeliveryService(
		deliveryRepository, s.driverRepository,
	)
	deliveryController := controllers.GetNewDeliveryController(
		deliveryService, s.sessionService,
	)

	deliverySubRouter.HandleFunc(
		"POST /{$}", deliveryController.Create,
	)

	deliverySubRouter.HandleFunc(
		"GET /{$}", deliveryController.ListAllForSender,
	)

	deliverySubRouter.HandleFunc("GET /{id}/{$}", deliveryController.Get)

	deliverySubRouter.HandleFunc(
		"PATCH /{id}/status/{$}", deliveryController.UpdateStatus,
	)

	deliverySubRouter.HandleFunc(
		"PATCH /{id}/assign/{$}", deliveryController.AssignDriver,
	)

	deliverySubRouter.HandleFunc("DELETE /{id}/{$}", deliveryController.Delete)

	deliverySubRouter.HandleFunc("GET /company/{$}", deliveryController.ListByCompany)

	deliverySubRouter.HandleFunc(
		"GET /pending/{$}", deliveryController.ListAllPendingsWithoutDriver,
	)

	return s.jwtController.AuthenticatedMiddleware(deliverySubRouter)
}
