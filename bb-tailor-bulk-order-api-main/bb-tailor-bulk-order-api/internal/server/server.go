package server

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"git.bopconsultancy.com/tejasc/bb-tailor-bulk-order-api/api/handlers"
	"git.bopconsultancy.com/tejasc/bb-tailor-bulk-order-api/api/middlewares"
	"git.bopconsultancy.com/tejasc/bb-tailor-bulk-order-api/internal/config"
)

type Server struct {
	Echo      *echo.Echo
	Config    *config.Config
	Db        *gorm.DB
	Validator *validator.Validate
	Enforcer  *casbin.Enforcer
}

func NewServer(cfg *config.Config) (*Server, error) {
	server := &Server{
		Config: cfg,
	}

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	server.Echo = e

	db, err := createGormDb(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to mysql db: %w", err)
	}

	server.Db = db

	server.Validator = validator.New(validator.WithRequiredStructEnabled())

	adapter, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		log.Fatalf("Error creating adapter: %v", err)
	}

	enforcer, err := casbin.NewEnforcer("internal/config/rbac_model.conf", adapter)
	if err != nil {
		log.Fatalf("Error creating enforcer: %v", err)
	}

	if err := enforcer.LoadPolicy(); err != nil {
		log.Fatalf("Error loading policy: %v", err)
	}

	server.Enforcer = enforcer

	server.RegisterRoutes()

	return server, nil
}

func (s *Server) StartServer() error {
	port := s.Config.Server.Port
	log.Printf("starting server on port %s\n", port)
	if err := s.Echo.Start(":" + port); err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}

	return nil
}

func (s *Server) GracefullyShutdown(ctx context.Context) error {
	return s.Echo.Shutdown(ctx)
}

func createGormDb(cfg *config.Config) (*gorm.DB, error) {
	dsn := cfg.MySqlDb.DbUser + ":" + cfg.MySqlDb.DbUserPassword + "@tcp(" + cfg.MySqlDb.DbHost + ":" + cfg.MySqlDb.DbPort + ")/" + cfg.MySqlDb.DbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil
}

func (s *Server) RegisterRoutes() error {
	healthHandler := handlers.NewHealthHandler()
	authHandler := handlers.NewAuthHandler(s.Db, s.Config, s.Validator)
	userHandler := handlers.NewUserHandler(s.Db, s.Config, s.Validator)

	s.registerHealthRoutes(s.Echo, healthHandler)
	s.registerAuthRoutes(s.Echo, authHandler)
	s.registerUserRoutes(s.Echo, userHandler)
	s.registerSwaggerRoutes(s.Echo, s.Config)

	return nil
}

func (s *Server) registerHealthRoutes(e *echo.Echo, h *handlers.HealthHandler) {
	healthGroup := e.Group("/health")

	healthGroup.GET("", h.GetHealth)
}

func (s *Server) registerAuthRoutes(e *echo.Echo, h *handlers.AuthHandler) {
	authGroup := e.Group("/auth")

	authGroup.POST("/login", h.Login)
}

func (s *Server) registerUserRoutes(e *echo.Echo, h *handlers.UserHandler) {
	userGroup := e.Group("/users")
	userGroup.Use(middlewares.JwtMiddleware(s.Config))
	userGroup.Use(
		middlewares.CasbinRBACMiddleware(
			&middlewares.CasbinRBACMiddlewareConfig{DB: s.Db, Enforcer: s.Enforcer},
		),
	)

	userGroup.GET("", h.GetAll)
}

func (s *Server) registerSwaggerRoutes(
	e *echo.Echo,
	cfg *config.Config,
) {
	e.Static("/docs", "web/static/docs")

	swaggerGroup := e.Group("/swagger")

	swaggerGroup.Use(
		middleware.BasicAuth(func(username, password string, ctx echo.Context) (bool, error) {
			if username == cfg.BasicAuth.Username && password == cfg.BasicAuth.Password {
				return true, nil
			}

			return false, nil
		}),
	)

	swaggerGroup.Use(middleware.AddTrailingSlashWithConfig(middleware.TrailingSlashConfig{
		RedirectCode: http.StatusMovedPermanently,
	}))

	swaggerGroup.Static("/", "web/static/swagger-ui")
}
