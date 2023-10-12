package commands

import (
	"context"
	"go-clean-monolith/pkg/cli"
	"go-clean-monolith/pkg/logger"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"log"
)

type ICommandRunner interface{}

// ICommand interface is used to implement sub-commands in the system.
type ICommand interface {
	Setup(cmd *cli.Command)
	Run() ICommandRunner
}

// additionalList of sub commands
var additionalList = map[string]ICommand{
	"start": NewStartCommand(),
}

// GetSubCommands gives a list of sub commands
func GetSubCommands(opt fx.Option) []*cli.Command {
	var subCommands []*cli.Command
	for name, cmd := range additionalList {
		subCommands = append(subCommands, RegisterSubCommands(name, cmd, opt))
	}
	return subCommands
}

// RegisterSubCommands register a list of sub commands
func RegisterSubCommands(name string, cmd ICommand, opt fx.Option) *cli.Command {
	subCmd := &cli.Command{
		Name: name,
		Run: func() {
			opts := fx.Options(
				fx.WithLogger(func() fxevent.Logger {
					return logger.NewFxLogger()
				}),
				fx.Invoke(cmd.Run()),
			)
			ctx := context.Background()
			app := fx.New(opt, opts)
			err := app.Start(ctx)
			defer func(app *fx.App, ctx context.Context) {
				_ = app.Stop(ctx)
			}(app, ctx)
			if err != nil {
				log.Fatalln(err)
			}
		},
	}
	cmd.Setup(subCmd)
	return subCmd
}
