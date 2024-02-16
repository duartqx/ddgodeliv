package server

import (
	"net/http"

	"ddgodeliv/api/controllers"
	"ddgodeliv/application/services"
	repository "ddgodeliv/infrastructure/repository/postgresql"
)

func (s server) SetupVehicleRoutes() http.Handler {

	vehicleRepository := repository.GetVehicleRepository(s.db)
	vehicleService := services.GetVehicleService(vehicleRepository)
	vehicleController := controllers.GetVehicleController(vehicleService, s.sessionService)

	vehicleModelRepository := repository.GetVehicleModelRepository(s.db)
	vehicleModelService := services.GetVehicleModelService(vehicleModelRepository)
	vehicleModelController := controllers.GetVehicleModelController(vehicleModelService)

	// Mux
	vehiclesSubRouter := http.NewServeMux()

	// Vehicle Routes
	vehiclesSubRouter.HandleFunc(
		"POST /{$}", vehicleController.CreateVehicle,
	)

	vehiclesSubRouter.HandleFunc(
		"GET /{$}", vehicleController.GetCompanyVehicles,
	)

	vehiclesSubRouter.HandleFunc(
		"GET /{id}/{$}", vehicleController.GetVehicle,
	)

	vehiclesSubRouter.HandleFunc(
		"DELETE /{id}/{$}", vehicleController.DeleteVehicle,
	)

	// VehicleModel Routes
	vehiclesSubRouter.HandleFunc(
		"GET /model/{$}", vehicleModelController.ListModels,
	)

	vehiclesSubRouter.HandleFunc(
		"POST /model/{$}", vehicleModelController.CreateVehicleModel,
	)

	return s.jwtController.AuthenticatedMiddleware(vehiclesSubRouter)
}
