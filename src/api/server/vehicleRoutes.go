package server

import (
	"ddgodeliv/api/controllers"
	"ddgodeliv/application/services"
	repository "ddgodeliv/infrastructure/repository/postgresql"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (s server) SetupVehicleRoutes() http.Handler {

	vehicleRepository := repository.GetNewVehicleRepository(s.db)
	vehicleService := services.GetNewVehicleService(vehicleRepository)
	vehicleController := controllers.GetNewVehicleController(vehicleService, s.sessionService)

	vehicleModelRepository := repository.GetNewVehicleModelRepository(s.db)
	vehicleModelService := services.GetNewVehicleModelService(vehicleModelRepository)
	vehicleModelController := controllers.GetNewVehicleModelController(vehicleModelService)

	// Mux
	vehiclesSubRouter := chi.NewRouter()

	vehiclesSubRouter.Use(s.jwtController.AuthenticatedMiddleware)

	// Vehicle Routes
	vehiclesSubRouter.Post("/", vehicleController.CreateVehicle)

	vehiclesSubRouter.Get("/", vehicleController.GetCompanyVehicles)

	vehiclesSubRouter.Get("/{id:[0-9]+}", vehicleController.GetVehicle)

	vehiclesSubRouter.Delete("/{id:[0-9]+}", vehicleController.DeleteVehicle)

	// VehicleModel Routes
	vehiclesSubRouter.Get("/model", vehicleModelController.ListModels)

	vehiclesSubRouter.Post("/model", vehicleModelController.CreateVehicleModel)

	return vehiclesSubRouter
}
