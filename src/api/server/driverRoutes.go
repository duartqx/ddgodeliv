package server

import (
	"net/http"

	"ddgodeliv/api/controllers"
	"ddgodeliv/application/services"
)

func (s server) SetupDriverRoutes() http.Handler {

	driverSubRouter := http.NewServeMux()

	userService := services.GetUserService(s.userRepository)
	driverService := services.GetDriverService(s.driverRepository, userService)
	driverController := controllers.GetDriverController(driverService, s.sessionService)

	driverSubRouter.HandleFunc("GET /{$}", driverController.ListCompanyDrivers)

	driverSubRouter.HandleFunc("POST /{$}", driverController.Create)

	driverSubRouter.HandleFunc("GET /{id}/{$}", driverController.Get)

	driverSubRouter.HandleFunc("PATCH /{id}/{$}", driverController.Update)

	driverSubRouter.HandleFunc("DELETE /{id}/{$}", driverController.Delete)

	return s.jwtController.AuthenticatedMiddleware(driverSubRouter)
}
