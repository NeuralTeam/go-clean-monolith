package app

import (
	"go-clean-monolith/cmd/commands"
	"go-clean-monolith/pkg/cli"
)

// SetupServer : create application
func SetupServer() error {
	rootCmd := commands.GetSubCommands(DependencyModules)
	cli.Parse(rootCmd)
	return nil
}
