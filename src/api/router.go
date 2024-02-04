package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"

	c "ddgodeliv/api/controllers"
	a "ddgodeliv/application/auth"
	s "ddgodeliv/application/services"
	"ddgodeliv/application/validation"
	r "ddgodeliv/infrastructure/repository/postgresql"

	lm "github.com/duartqx/ddgomiddlewares/logger"
	rm "github.com/duartqx/ddgomiddlewares/recovery"
)

type AuthHandler interface {
	AuthenticatedMiddleware(http.Handler) http.Handler
	UnauthenticatedMiddleware(http.Handler) http.Handler
}

type router struct {
	db     *sqlx.DB
	secret *[]byte
}

func NewRouterBuilder() *router {
	return &router{}
}

func (ro *router) SetDb(db *sqlx.DB) *router {
	ro.db = db
	return ro
}

func (ro *router) SetSecret(secret []byte) *router {
	ro.secret = &secret
	return ro
}

func (ro router) userRoutes(
	userService *s.UserService,
	sessionService *a.SessionService,
	authHandler AuthHandler,
) *chi.Mux {

	userController := c.GetNewUserController(userService, sessionService)

	userSubRouter := chi.NewRouter()

	// POST: Create User
	userSubRouter.
		With(authHandler.UnauthenticatedMiddleware).
		Post("/", userController.Create)

	// GET: Self (Good for checking if the user is authenticated)
	userSubRouter.
		With(authHandler.AuthenticatedMiddleware).
		Get("/", userController.Get)

	// PATCH: Password Update
	userSubRouter.
		With(authHandler.AuthenticatedMiddleware).
		Patch("/password", userController.UpdatePassword)

	return userSubRouter
}

func (ro router) vehicleRoutes(
	sessionRepository *a.SessionService,
	v *validation.Validator,
	authHandler AuthHandler,
) *chi.Mux {

	// Repositories
	vehicleRepository := r.GetNewVehicleRepository(ro.db)
	vehicleModelRepository := r.GetNewVehicleModelRepository(ro.db)

	// Services
	vehicleService := s.GetNewVehicleService(vehicleRepository, v)
	vehicleModelService := s.GetNewVehicleModelService(vehicleModelRepository, v)

	// Controllers
	vehicleController := c.GetNewVehicleController(
		vehicleService, sessionRepository,
	)
	vehicleModelController := c.GetNewVehicleModelController(vehicleModelService)

	vehiclesSubRouter := chi.NewRouter()

	// Vehicle Routes
	vehiclesSubRouter.
		With(authHandler.AuthenticatedMiddleware).
		Post("/", vehicleController.CreateVehicle)

	vehiclesSubRouter.
		With(authHandler.AuthenticatedMiddleware).
		Get("/", vehicleController.GetCompanyVehicles)

	vehiclesSubRouter.
		With(authHandler.AuthenticatedMiddleware).
		Get("/{id:[0-9]+}", vehicleController.GetVehicle)

	vehiclesSubRouter.
		With(authHandler.AuthenticatedMiddleware).
		Delete("/{id:[0-9]+}", vehicleController.DeleteVehicle)

	// VehicleModel Routes
	vehiclesSubRouter.
		With(authHandler.AuthenticatedMiddleware).
		Get("/model", vehicleModelController.ListModels)

	vehiclesSubRouter.
		With(authHandler.AuthenticatedMiddleware).
		Post("/model", vehicleModelController.CreateVehicleModel)

	return vehiclesSubRouter
}

func (ro router) Build() *chi.Mux {

	v := validation.NewValidator()

	// Repositories
	driverRepository := r.GetNewDriverRepository(ro.db)
	sessionRepository := r.GetNewSessionRepository()
	userRepository := r.GetNewUserRepository(ro.db)

	// Services
	jwtAuthService := a.GetNewJwtAuthService(
		userRepository, driverRepository, sessionRepository, ro.secret,
	)
	sessionService := a.GetNewSessionService(driverRepository, sessionRepository)
	userService := s.GetNewUserService(userRepository, v)

	// Controllers
	jwtController := c.NewJwtController(jwtAuthService)

	router := chi.NewRouter()

	router.Use(rm.RecoveryMiddleware, lm.LoggerMiddleware)

	// Auth Routes
	// POST: User Login
	router.
		With(jwtController.UnauthenticatedMiddleware).
		Post("/login", jwtController.Login)

	// DELETE: User Logout
	router.
		With(jwtController.AuthenticatedMiddleware).
		Delete("/logout", jwtController.Logout)

	// User Routes
	router.Mount(
		"/user",
		ro.userRoutes(
			userService,
			sessionService,
			jwtController,
		),
	)

	// Vehicle Routes
	router.Mount(
		"/vehicles",
		ro.vehicleRoutes(sessionService, v, jwtController),
	)

	return router
}
