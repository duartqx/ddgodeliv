package server

import (
	"ddgodeliv/api/controllers"
	"ddgodeliv/application/services"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (s server) SetupDriverRoutes() http.Handler {

	driverSubRouter := chi.NewRouter()

	driverSubRouter.Use(s.jwtController.AuthenticatedMiddleware)

	userService := services.GetNewUserService(s.userRepository)
	driverService := services.GetNewDriverService(s.driverRepository, userService)
	driverController := controllers.GetNewDriverController(driverService, s.sessionService)

	driverSubRouter.Get("/", driverController.ListCompanyDrivers)

	driverSubRouter.Post("/", driverController.Create)

	driverSubRouter.Get("/{id:[0-9]+}", driverController.Get)

	driverSubRouter.Patch("/{id:[0-9]+}", driverController.Update)

	driverSubRouter.Delete("/{id:[0-9]+}", driverController.Delete)

	return driverSubRouter
}
