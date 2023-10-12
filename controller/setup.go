package controller

import (
	"go-clean-monolith/controller/middleware"
	v1 "go-clean-monolith/controller/v1"
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
