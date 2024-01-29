package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"

	c "ddgodeliv/api/controllers"
	s "ddgodeliv/application/services"
	"ddgodeliv/application/validation"
	r "ddgodeliv/infrastructure/postgresql"

	lm "github.com/duartqx/ddgomiddlewares/logger"
	rm "github.com/duartqx/ddgomiddlewares/recovery"
)

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

func (ro router) userRoutes(userService *s.UserService, jwtController *c.JwtController) *chi.Mux {

	userController := c.GetNewUserController(userService)

	userSubrouter := chi.NewRouter()

	// POST: Create User
	userSubrouter.
		With(jwtController.UnauthenticatedMiddleware).
		Method(http.MethodPost, "/", userController)

	// GET: Self (Good for checking if the user is authenticated)
	userSubrouter.
		With(jwtController.AuthenticatedMiddleware).
		Method(http.MethodGet, "/", userController)

	// PATCH: Password Update
	userSubrouter.
		With(jwtController.AuthenticatedMiddleware).
		Method(http.MethodPatch, "/password", userController)

	return userSubrouter
}

func (ro router) Build() *chi.Mux {

	v := validation.NewValidator()

	userRepository := r.GetNewUserRepository(ro.db)
	userService := s.GetNewUserService(userRepository, v)

	jwtAuthService := s.GetNewJwtAuthService(
		userRepository, ro.secret, r.GetNewSessionRepository(),
	)
	jwtController := c.NewJwtController(jwtAuthService)

	router := chi.NewRouter()

	router.Use(rm.RecoveryMiddleware, lm.LoggerMiddleware)

	// Auth Routes
	// POST: User Login
	router.
		With(jwtController.UnauthenticatedMiddleware).
		Method(http.MethodPost, "/login", jwtController)

	// DELETE: User Logout
	router.
		With(jwtController.AuthenticatedMiddleware).
		Method(http.MethodDelete, "/logout", jwtController)

	// User Routes
	router.Mount("/user", ro.userRoutes(userService, jwtController))

	return router
}
