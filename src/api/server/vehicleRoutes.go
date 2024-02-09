package server

import (
	"net/http"

	"ddgodeliv/api/controllers"
	"ddgodeliv/application/services"
	repository "ddgodeliv/infrastructure/repository/postgresql"
)

func (s server) SetupVehicleRoutes() http.Handler {

	vehicleRepository := repository.GetNewVehicleRepository(s.db)
	vehicleService := services.GetNewVehicleService(vehicleRepository)
	vehicleController := controllers.GetNewVehicleController(vehicleService, s.sessionService)

	vehicleModelRepository := repository.GetNewVehicleModelRepository(s.db)
	vehicleModelService := services.GetNewVehicleModelService(vehicleModelRepository)
	vehicleModelController := controllers.GetNewVehicleModelController(vehicleModelService)

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
