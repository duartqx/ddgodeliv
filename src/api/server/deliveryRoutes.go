package server

import (
	"net/http"

	"ddgodeliv/api/controllers"
	"ddgodeliv/application/services"
	repository "ddgodeliv/infrastructure/repository/postgresql"
)

func (s server) SetupDeliveryRoutes() http.Handler {

	deliverySubRouter := http.NewServeMux()

	deliveryRepository := repository.GetDeliveryRepository(s.db)
	deliveryService := services.GetDeliveryService(
		deliveryRepository, s.driverRepository,
	)
	deliveryController := controllers.GetDeliveryController(
		deliveryService,
		services.GetDriverService(
			s.driverRepository, services.GetUserService(s.userRepository),
		),
		s.sessionService,
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

	deliverySubRouter.HandleFunc(
		"GET /company/{$}", deliveryController.ListByCompany,
	)

	deliverySubRouter.HandleFunc(
		"GET /company/{id}/{$}", deliveryController.Get,
	)

	deliverySubRouter.HandleFunc(
		"GET /company/driver/{id}/{$}", deliveryController.ListAllForDriver,
	)

	deliverySubRouter.HandleFunc(
		"GET /company/driver/{id}/current/{$}",
		deliveryController.GetDriverCurrentDelivery,
	)

	deliverySubRouter.HandleFunc(
		"GET /pending/{$}", deliveryController.ListAllPendingsWithoutDriver,
	)

	return s.jwtController.AuthenticatedMiddleware(deliverySubRouter)
}
