package app

import (
	"go-clean-monolith/config"
	"go-clean-monolith/internal/controller"
	"go-clean-monolith/internal/gateway"
	"go-clean-monolith/internal/repository"
	"go-clean-monolith/internal/service"
	"go-clean-monolith/pkg/httpserver"
	"go-clean-monolith/pkg/logger"
	"go.uber.org/fx"
)

// DependencyModules exports dependency
var DependencyModules = fx.Options(
	fx.Provide(config.New),
	fx.Provide(logger.New),
	fx.Provide(httpserver.New),

	controller.DependencyModules,
	service.DependencyModules,
	repository.DependencyModules,
	gateway.DependencyModules,
)
