package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"

	"ddgodeliv/api/controllers"
	"ddgodeliv/application/services/auth"
	"ddgodeliv/domains/driver"
	"ddgodeliv/domains/user"
	repository "ddgodeliv/infrastructure/repository/postgresql"

	lm "github.com/duartqx/ddgomiddlewares/logger"
	rm "github.com/duartqx/ddgomiddlewares/recovery"
)

var s *server

type server struct {
	db     *sqlx.DB
	secret *[]byte

	// Base Repos
	userRepository   user.IUserRepository
	driverRepository driver.IDriverRepository

	// Base Services
	sessionService *auth.SessionService

	// AuthController / AuthMiddleware
	jwtController *controllers.JwtController
}

type ServerConfig struct {
	Db     *sqlx.DB
	Secret []byte
}

func GetNewServer(cfg ServerConfig) *server {
	s := &server{
		db:     cfg.Db,
		secret: &cfg.Secret,
	}
	return s.setupBase()
}

func (s *server) setupBase() *server {
	// Repositories
	s.userRepository = repository.GetNewUserRepository(s.db)
	s.driverRepository = repository.GetNewDriverRepository(s.db)

	sessionRepository := repository.GetNewSessionRepository()
	// Services
	s.sessionService = auth.GetNewSessionService(
		s.driverRepository, sessionRepository,
	)

	s.jwtController = controllers.NewJwtController(
		auth.GetNewJwtAuthService(
			s.userRepository, sessionRepository, s.secret,
		),
	)
	return s
}

func (s *server) BuildBaseServer() *chi.Mux {
	mux := chi.NewRouter()

	mux.Use(rm.RecoveryMiddleware, lm.LoggerMiddleware)

	// Auth Routes
	// POST: User Login
	mux.
		With(s.jwtController.UnauthenticatedMiddleware).
		Post("/login", s.jwtController.Login)

	// DELETE: User Logout
	mux.
		With(s.jwtController.AuthenticatedMiddleware).
		Delete("/logout", s.jwtController.Logout)

	return mux
}

func (s *server) Build() http.Handler {

	mux := s.BuildBaseServer()

	// User Routes
	mux.Mount("/user", s.SetupUserRoutes())

	// Vehicle Routes
	mux.Mount("/vehicle", s.SetupVehicleRoutes())

	// Company Routes
	mux.Mount("/company", s.SetupCompanyRoutes())

	// Driver Routes
	mux.Mount("/driver", s.SetupDriverRoutes())

	// Delivery Routes
	mux.Mount("/delivery", s.SetupDeliveryRoutes())

	return mux
}
