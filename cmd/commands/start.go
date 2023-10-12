package commands

import (
	"errors"
	"fmt"
	"go-clean-monolith/config"
	"go-clean-monolith/controller/middleware"
	v1 "go-clean-monolith/controller/v1"
	"go-clean-monolith/pkg/cli"
	"go-clean-monolith/pkg/httpserver"
	"go-clean-monolith/pkg/logger"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

// StartCommand launch web application
type StartCommand struct{}

func (s StartCommand) Setup(cmd *cli.Command) {}

func (s StartCommand) Run() ICommandRunner {
	return func(
		httpServer httpserver.HttpServer,
		middleware middleware.Middleware,
		routerV1 v1.Router,
		env config.Env,
		logger logger.Logger,
	) {
		httpServer.Use(middleware.Logger.Setup(), middleware.CORS.Setup(), middleware.Recovery.Setup())
		routerV1.Setup()

		srv := &http.Server{
			Addr:    fmt.Sprintf("%s:%d", env.ServerHost, env.ServerPort),
			Handler: httpServer,
		}

		go func() {
			if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
				logger.Printf("Server listen: %s\n", err)
			}
		}()

		quit := make(chan os.Signal)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit
	}
}

func NewStartCommand() StartCommand {
	return StartCommand{}
}
