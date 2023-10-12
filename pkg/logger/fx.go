package logger

import "go.uber.org/fx"

// DependencyModules exports dependencies
var DependencyModules = fx.Options(
	fx.Provide(NewLogger),
)
