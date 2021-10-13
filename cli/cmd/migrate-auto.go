package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/takemo101/fibar-core/pkg"
)

// AutoMigrateCommand is struct
type AutoMigrateCommand struct {
	logger pkg.Logger
	root   RootCommand
	db     pkg.Database
	models []interface{}
}

// Setup is setup command
func (c AutoMigrateCommand) Setup() {
	c.logger.Info("setup migrate:auto-command")

	c.root.Cmd.AddCommand(&cobra.Command{
		Use:   "migrate:auto",
		Short: "auto migrate from model",
		Run: func(cmd *cobra.Command, args []string) {

			c.db.GormDB.AutoMigrate(c.models...)

			fmt.Println("finish migrate")
		},
	})
}

// NewAutoMigrateCommand create migrate command
func NewAutoMigrateCommand(
	root RootCommand,
	logger pkg.Logger,
	db pkg.Database,
	options CommandOptions,
) AutoMigrateCommand {
	return AutoMigrateCommand{
		root:   root,
		logger: logger,
		db:     db,
		models: options.Models,
	}
}
