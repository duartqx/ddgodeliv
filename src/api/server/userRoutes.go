package server

import (
	"ddgodeliv/api/controllers"
	"ddgodeliv/application/services"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (s server) SetupUserRoutes() http.Handler {

	userSubRouter := chi.NewRouter()

	userService := services.GetNewUserService(s.userRepository)
	userController := controllers.GetNewUserController(userService, s.sessionService)

	// POST: Create User
	userSubRouter.
		With(s.jwtController.UnauthenticatedMiddleware).
		Post("/", userController.Create)

	// GET: Self (Good for checking if the user is authenticated)
	userSubRouter.
		With(s.jwtController.AuthenticatedMiddleware).
		Get("/", userController.Get)

	// PATCH: Password Update
	userSubRouter.
		With(s.jwtController.AuthenticatedMiddleware).
		Patch("/password", userController.UpdatePassword)

	return userSubRouter
}
