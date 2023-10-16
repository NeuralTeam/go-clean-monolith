package commands

import (
	"errors"
	"fmt"
	"go-clean-monolith/config"
	"go-clean-monolith/internal/controller/middleware"
	"go-clean-monolith/internal/controller/v1"
	"go-clean-monolith/pkg/cli"
	"go-clean-monolith/pkg/httpserver"
	"go-clean-monolith/pkg/logger"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type StartCommand struct{}

// NewStartCommand launch web application
func NewStartCommand() StartCommand {
	return StartCommand{}
}

func (s StartCommand) Setup(cmd *cli.Command) {}

func (s StartCommand) Run() ICommandRunner {
	return func(
		httpServer httpserver.HttpServer,
		middleware middleware.Middleware,
		routerV1 v1.Router,
		env config.Env,
		logger logger.Logger,
	) {
		httpServer.EngineUse(middleware.Logger.Setup(), middleware.CORS.Setup(), middleware.Recovery.Setup())
		routerV1.Setup()

		srv := &http.Server{
			Addr:    fmt.Sprintf("%s:%d", env.ServerHost, env.ServerPort),
			Handler: httpServer.Handler(),
		}

		logger.Info().
			IPAddr("ip", env.ServerHost).
			Int("port", env.ServerPort).
			Str("environment", env.Environment).
			Str("version", env.Version).
			Msg("Overview:")
		for _, handler := range httpServer.Routes() {
			logger.Debug().Str("method", handler.Method).Str("path", handler.Path).Msg("Create handler:")
		}

		go func() {
			if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
				logger.Fatal().Msgf("Server listen: %s", err)
			}
		}()

		quit := make(chan os.Signal)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit
	}
}
