package server

import (
	"net/http"

	"github.com/jmoiron/sqlx"

	"ddgodeliv/api/controllers"
	"ddgodeliv/application/services/auth"
	"ddgodeliv/domains/driver"
	"ddgodeliv/domains/user"
	repository "ddgodeliv/infrastructure/repository/postgresql"

	lm "github.com/duartqx/ddgomiddlewares/logger"
	rm "github.com/duartqx/ddgomiddlewares/recovery"
	tm "github.com/duartqx/ddgomiddlewares/trailling"
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

func (s *server) BuildBaseServer() *http.ServeMux {

	mux := http.NewServeMux()

	// Auth Routes
	// POST: User Login
	mux.Handle(
		"POST /login/{$}",
		s.jwtController.UnauthenticatedMiddleware(
			http.HandlerFunc(s.jwtController.Login),
		),
	)

	// DELETE: User Logout
	mux.Handle(
		"DELETE /logout/{$}",
		s.jwtController.AuthenticatedMiddleware(
			http.HandlerFunc(s.jwtController.Logout),
		),
	)

	return mux
}

func (s *server) Build() http.Handler {

	mux := s.BuildBaseServer()

	// User Routes
	mux.Handle("/user/", http.StripPrefix("/user", s.SetupUserRoutes()))

	// Vehicle Routes
	mux.Handle("/vehicle/", http.StripPrefix("/vehicle", s.SetupVehicleRoutes()))

	// Company Routes
	mux.Handle("/company/", http.StripPrefix("/company", s.SetupCompanyRoutes()))

	// Driver Routes
	mux.Handle("/driver/", http.StripPrefix("/driver", s.SetupDriverRoutes()))

	// Delivery Routes
	mux.Handle("/delivery/", http.StripPrefix("/delivery", s.SetupDeliveryRoutes()))

	// Recovery and Logger middleware
	wrapped := tm.TrailingSlashMiddleware(rm.RecoveryMiddleware(lm.LoggerMiddleware(mux)))

	return wrapped
}
