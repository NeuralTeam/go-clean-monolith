package app

import (
	"go-clean-monolith/config"
	"go-clean-monolith/controller"
	"go-clean-monolith/gateway"
	"go-clean-monolith/pkg/httpserver"
	"go-clean-monolith/pkg/logger"
	"go-clean-monolith/repository"
	"go-clean-monolith/service"
	"go.uber.org/fx"
)

// DependencyModules exports dependency
var DependencyModules = fx.Options(
	controller.DependencyModules,
	service.DependencyModules,
	repository.DependencyModules,
	gateway.DependencyModules,
	config.DependencyModules,
	logger.DependencyModules,
	httpserver.DependencyModules,
)
