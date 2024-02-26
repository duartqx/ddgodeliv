package server

import (
	"net/http"

	"github.com/jmoiron/sqlx"

	"ddgodeliv/api/controllers"
	"ddgodeliv/application/services/auth"
	"ddgodeliv/domains/driver"
	"ddgodeliv/domains/user"
	repository "ddgodeliv/infrastructure/repository/postgresql"

	cm "github.com/duartqx/ddgomiddlewares/cors"
	lm "github.com/duartqx/ddgomiddlewares/logger"
	rm "github.com/duartqx/ddgomiddlewares/recovery"
	tm "github.com/duartqx/ddgomiddlewares/trailling"
)

var s *server

type server struct {
	db     *sqlx.DB
	secret *[]byte
	index  *[]byte
	assets *string

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
	Index  []byte
	Assets string
}

func GetNewServer(cfg ServerConfig) *server {
	s := &server{
		db:     cfg.Db,
		secret: &cfg.Secret,
		index:  &cfg.Index,
		assets: &cfg.Assets,
	}
	return s.setupBase()
}

func (s *server) setupBase() *server {
	// Repositories
	s.userRepository = repository.GetUserRepository(s.db)
	s.driverRepository = repository.GetDriverRepository(s.db)

	sessionRepository := repository.GetSessionRepository()
	// Services
	s.sessionService = auth.GetSessionService(
		s.driverRepository, sessionRepository,
	)

	s.jwtController = controllers.GetJwtController(
		auth.GetJwtAuthService(
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

func (s *server) Use(
	mux http.Handler, middlewares ...func(http.Handler) http.Handler,
) http.Handler {
	for _, m := range middlewares {
		mux = m(mux)
	}
	return mux
}

func (s *server) Build() http.Handler {

	mux := s.BuildBaseServer()
	// User Routes
	mux.Handle("/api/user/", http.StripPrefix("/api/user", s.SetupUserRoutes()))

	// Vehicle Routes
	mux.Handle("/api/vehicle/", http.StripPrefix("/api/vehicle", s.SetupVehicleRoutes()))

	// Company Routes
	mux.Handle("/api/company/", http.StripPrefix("/api/company", s.SetupCompanyRoutes()))

	// Driver Routes
	mux.Handle("/api/driver/", http.StripPrefix("/api/driver", s.SetupDriverRoutes()))

	// Delivery Routes
	mux.Handle("/api/delivery/", http.StripPrefix("/api/delivery", s.SetupDeliveryRoutes()))

	mux.Handle(
		"/assets/",
		http.StripPrefix("/assets/", http.FileServer(http.Dir(*s.assets))),
	)

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Write(*s.index)
	})

	return s.Use(
		mux,
		cm.CorsMiddleware,
		tm.TrailingSlashMiddleware,
		lm.LoggerMiddleware,
		rm.RecoveryMiddleware,
	)
}
