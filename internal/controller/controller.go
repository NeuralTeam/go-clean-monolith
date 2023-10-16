package controller

import (
	"go-clean-monolith/internal/controller/middleware"
	"go-clean-monolith/internal/controller/v1"
	"go.uber.org/fx"
)

// DependencyModules exports dependencies
var DependencyModules = fx.Options(
	// Middlewares
	fx.Provide(middleware.NewMiddleware),
	fx.Provide(middleware.NewLoggerMiddleware),
	fx.Provide(middleware.NewCORSMiddleware),
	fx.Provide(middleware.NewRecoveryMiddleware),

	// V1
	// Router
	fx.Provide(v1.NewRouter),
	// Controllers
	fx.Provide(v1.NewUserController),
)
