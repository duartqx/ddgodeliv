package server

import (
	"net/http"

	"ddgodeliv/api/controllers"
	"ddgodeliv/application/services"
)

func (s server) SetupUserRoutes() http.Handler {

	userSubRouter := http.NewServeMux()

	userService := services.GetNewUserService(s.userRepository)
	userController := controllers.GetNewUserController(userService, s.sessionService)

	// POST: Create User
	userSubRouter.Handle(
		"POST /{$}",
		s.jwtController.UnauthenticatedMiddleware(
			http.HandlerFunc(userController.Create),
		),
	)

	// GET: Self (Good for checking if the user is authenticated)
	userSubRouter.Handle(
		"GET /{$}",
		s.jwtController.AuthenticatedMiddleware(
			http.HandlerFunc(userController.Get),
		),
	)

	// PATCH: Password Update
	userSubRouter.Handle(
		"POST /password/{$}",
		s.jwtController.AuthenticatedMiddleware(
			http.HandlerFunc(userController.UpdatePassword),
		),
	)

	return userSubRouter
}
