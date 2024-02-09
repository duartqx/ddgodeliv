package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"

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

	companySubRouter := chi.NewRouter()

	companySubRouter.Use(s.jwtController.AuthenticatedMiddleware)

	// POST: Create company
	companySubRouter.Post("/", companyController.Create)

	// DELETE: Owner can delete it's company
	companySubRouter.Delete("/", companyController.Delete)

	return companySubRouter
}
