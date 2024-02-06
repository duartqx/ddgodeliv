package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"

	c "ddgodeliv/api/controllers"
	s "ddgodeliv/application/services"
	a "ddgodeliv/application/services/auth"
	"ddgodeliv/domains/auth"
	"ddgodeliv/domains/company"
	"ddgodeliv/domains/delivery"
	"ddgodeliv/domains/driver"
	"ddgodeliv/domains/user"
	"ddgodeliv/domains/vehicle"
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

	// Repositories
	companyRepository      company.ICompanyRepository
	deliveryRepository     delivery.IDeliveryRepository
	driverRepository       driver.IDriverRepository
	sessionRepository      auth.ISessionRepository
	userRepository         user.IUserRepository
	vehicleModelRepository vehicle.IVehicleModelRepository
	vehicleRepository      vehicle.IVehicleRepository

	// Services
	companyService      *s.CompanyService
	deliveryService     *s.DeliveryService
	driverService       *s.DriverService
	jwtAuthService      *a.JwtAuthService
	sessionService      *a.SessionService
	userService         *s.UserService
	vehicleModelService *s.VehicleModelService
	vehicleService      *s.VehicleService

	// Controllers
	companyController      *c.CompanyController
	jwtController          *c.JwtController
	deliveryController     *c.DeliveryController
	driverController       *c.DriverController
	vehicleModelController *c.VehicleModelController
	vehicleController      *c.VehicleController
	userController         *c.UserController
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

func (ro router) userRoutes() *chi.Mux {

	userSubRouter := chi.NewRouter()

	// POST: Create User
	userSubRouter.
		With(ro.jwtController.UnauthenticatedMiddleware).
		Post("/", ro.userController.Create)

	// GET: Self (Good for checking if the user is authenticated)
	userSubRouter.
		With(ro.jwtController.AuthenticatedMiddleware).
		Get("/", ro.userController.Get)

	// PATCH: Password Update
	userSubRouter.
		With(ro.jwtController.AuthenticatedMiddleware).
		Patch("/password", ro.userController.UpdatePassword)

	return userSubRouter
}

func (ro router) companyRoutes() *chi.Mux {

	companySubRouter := chi.NewRouter()

	// POST: Create company
	companySubRouter.
		With(ro.jwtController.AuthenticatedMiddleware).
		Post("/", ro.companyController.Create)

	// DELETE: Owner can delete it's company
	companySubRouter.
		With(ro.jwtController.AuthenticatedMiddleware).
		Delete("/", ro.companyController.Delete)

	return companySubRouter
}

func (ro router) driverRoutes() *chi.Mux {

	driverSubRouter := chi.NewRouter()

	driverSubRouter.
		With(ro.jwtController.AuthenticatedMiddleware).
		Get("/", ro.driverController.ListCompanyDrivers)

	driverSubRouter.
		With(ro.jwtController.AuthenticatedMiddleware).
		Post("/", ro.driverController.Create)

	driverSubRouter.
		With(ro.jwtController.AuthenticatedMiddleware).
		Get("/{id:[0-9]+}", ro.driverController.Get)

	driverSubRouter.
		With(ro.jwtController.AuthenticatedMiddleware).
		Patch("/{id:[0-9]+}", ro.driverController.Update)

	driverSubRouter.
		With(ro.jwtController.AuthenticatedMiddleware).
		Delete("/{id:[0-9]+}", ro.driverController.Delete)

	return driverSubRouter
}

func (ro router) vehicleRoutes() *chi.Mux {

	vehiclesSubRouter := chi.NewRouter()

	// Vehicle Routes
	vehiclesSubRouter.
		With(ro.jwtController.AuthenticatedMiddleware).
		Post("/", ro.vehicleController.CreateVehicle)

	vehiclesSubRouter.
		With(ro.jwtController.AuthenticatedMiddleware).
		Get("/", ro.vehicleController.GetCompanyVehicles)

	vehiclesSubRouter.
		With(ro.jwtController.AuthenticatedMiddleware).
		Get("/{id:[0-9]+}", ro.vehicleController.GetVehicle)

	vehiclesSubRouter.
		With(ro.jwtController.AuthenticatedMiddleware).
		Delete("/{id:[0-9]+}", ro.vehicleController.DeleteVehicle)

	// VehicleModel Routes
	vehiclesSubRouter.
		With(ro.jwtController.AuthenticatedMiddleware).
		Get("/model", ro.vehicleModelController.ListModels)

	vehiclesSubRouter.
		With(ro.jwtController.AuthenticatedMiddleware).
		Post("/model", ro.vehicleModelController.CreateVehicleModel)

	return vehiclesSubRouter
}

func (ro router) deliveryRoutes() *chi.Mux {

	deliverySubRouter := chi.NewRouter()

	deliverySubRouter.
		With(ro.jwtController.AuthenticatedMiddleware).
		Post("/", ro.deliveryController.Create)

	deliverySubRouter.
		With(ro.jwtController.AuthenticatedMiddleware).
		Get("/{id:[0+9]+}", ro.deliveryController.Get)

	deliverySubRouter.
		With(ro.jwtController.AuthenticatedMiddleware).
		Patch("/{id:[0+9]+}/status", ro.deliveryController.UpdateStatus)

	deliverySubRouter.
		With(ro.jwtController.AuthenticatedMiddleware).
		Patch("/{id:[0+9]+}/assign", ro.deliveryController.AssignDriver)

	deliverySubRouter.
		With(ro.jwtController.AuthenticatedMiddleware).
		Delete("/{id:[0+9]+}", ro.deliveryController.Delete)

	// deliverySubRouter.
	// 	With(ro.jwtController.AuthenticatedMiddleware).
	// 	Get("/", ro.deliveryController.ListAllForSender)

	deliverySubRouter.
		With(ro.jwtController.AuthenticatedMiddleware).
		Get("/company", ro.deliveryController.ListByCompany)

	deliverySubRouter.
		With(ro.jwtController.AuthenticatedMiddleware).
		Get("/pending", ro.deliveryController.ListAllPendingsWithoutDriver)

	return deliverySubRouter
}

func (ro router) SetupRepositories() router {

	ro.companyRepository = r.GetNewCompanyRepository(ro.db)
	ro.deliveryRepository = r.GetNewDeliveryRepository(ro.db)
	ro.driverRepository = r.GetNewDriverRepository(ro.db)
	ro.sessionRepository = r.GetNewSessionRepository()
	ro.userRepository = r.GetNewUserRepository(ro.db)
	ro.vehicleModelRepository = r.GetNewVehicleModelRepository(ro.db)
	ro.vehicleRepository = r.GetNewVehicleRepository(ro.db)

	return ro
}

func (ro router) SetupServices() router {

	ro.companyService = s.GetNewCompanyService(ro.companyRepository)
	ro.deliveryService = s.GetNewDeliveryService(
		ro.deliveryRepository, ro.driverRepository,
	)
	ro.jwtAuthService = a.GetNewJwtAuthService(
		ro.userRepository, ro.driverRepository, ro.sessionRepository, ro.secret,
	)
	ro.sessionService = a.GetNewSessionService(
		ro.driverRepository, ro.sessionRepository,
	)
	ro.userService = s.GetNewUserService(ro.userRepository)
	ro.driverService = s.GetNewDriverService(ro.driverRepository, ro.userService)
	ro.vehicleModelService = s.GetNewVehicleModelService(ro.vehicleModelRepository)
	ro.vehicleService = s.GetNewVehicleService(ro.vehicleRepository)

	return ro
}

func (ro router) SetupControllers() router {

	ro.jwtController = c.NewJwtController(ro.jwtAuthService)
	ro.companyController = c.GetNewCompanyController(
		ro.companyService, ro.sessionService,
	)
	ro.deliveryController = c.GetNewDeliveryController(
		ro.deliveryService, ro.sessionService,
	)
	ro.driverController = c.GetNewDriverController(
		ro.driverService, ro.sessionService,
	)
	ro.vehicleController = c.GetNewVehicleController(
		ro.vehicleService, ro.sessionService,
	)
	ro.vehicleModelController = c.GetNewVehicleModelController(
		ro.vehicleModelService,
	)
	ro.userController = c.GetNewUserController(ro.userService, ro.sessionService)

	return ro
}

func (ro router) Build() *chi.Mux {

	router := chi.NewRouter()

	router.Use(rm.RecoveryMiddleware, lm.LoggerMiddleware)

	// Auth Routes
	// POST: User Login
	router.
		With(ro.jwtController.UnauthenticatedMiddleware).
		Post("/login", ro.jwtController.Login)

	// DELETE: User Logout
	router.
		With(ro.jwtController.AuthenticatedMiddleware).
		Delete("/logout", ro.jwtController.Logout)

	// User Routes
	router.Mount("/user", ro.userRoutes())

	// Vehicle Routes
	router.Mount("/vehicle", ro.vehicleRoutes())

	// Company Routes
	router.Mount("/company", ro.companyRoutes())

	// Driver Routes
	router.Mount("/driver", ro.driverRoutes())

	// Delivery Routes
	router.Mount("/delivery", ro.deliveryRoutes())

	return router
}
