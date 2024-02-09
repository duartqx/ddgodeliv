package server

import (
	"net/http"

	"ddgodeliv/api/controllers"
	"ddgodeliv/application/services"
	repository "ddgodeliv/infrastructure/repository/postgresql"
)

func (s *server) SetupCompanyRoutes() http.Handler {

	companyRepository := repository.GetNewCompanyRepository(s.db)
	companyService := services.GetNewCompanyService(companyRepository)
	companyController := controllers.GetNewCompanyController(
		companyService, s.sessionService,
	)

	companySubRouter := http.NewServeMux()

	// POST: Create company
	companySubRouter.HandleFunc("POST /{$}", companyController.Create)

	// DELETE: Owner can delete it's company
	companySubRouter.HandleFunc("DELETE /{$}", companyController.Delete)

	return s.jwtController.AuthenticatedMiddleware(companySubRouter)
}
