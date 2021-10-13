package main

import (
	"github.com/takemo101/fibar-core/app"
	"github.com/takemo101/fibar-core/cli/cmd"
	"github.com/takemo101/fibar-core/cli/kernel"
	"github.com/takemo101/fibar-core/database"
	"github.com/takemo101/fibar-core/pkg/contract"
	"go.uber.org/fx"
)

// CLIBooter is command module root struct
type CLIBooter struct{}

// CLIBooter all setup
func (booter CLIBooter) CLIBoot() {
	//
}

// NewAppBooter app create
func NewCLIBooter() contract.CLIBooter {
	return CLIBooter{}
}

func main() {
	// boot cobra application
	kernel.Run(
		kernel.CLIOptions{
			ConfigPath: "config.yml",
			CommandOptions: cmd.CommandOptions{
				Models:     database.Models,
				Migrations: database.Migrations,
			},
			CLIBooterConstructor: NewCLIBooter,
			FXOption: fx.Options(
				app.Module,
			),
		},
	)
}
