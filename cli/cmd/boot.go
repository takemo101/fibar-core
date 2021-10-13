package cmd

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/takemo101/fibar-core/pkg/contract"
	"go.uber.org/fx"
)

// Module export
var Module = fx.Options(
	fx.Provide(NewStorageLinkCommand),
	fx.Provide(NewMigrateCommand),
	fx.Provide(NewRollbackCommand),
	fx.Provide(NewAutoMigrateCommand),
	fx.Provide(NewCommandRoot),
	fx.Provide(NewBaseCommand),
)

// CommandOptions is base commands options
type CommandOptions struct {
	Migrations []*gormigrate.Migration
	Models     []interface{}
}

// BaseCommands is base commands slice
type BaseCommands []contract.Command

// NewBaseCommand is setup command
func NewBaseCommand(
	storageLinkCommand StorageLinkCommand,
	migrateCommand MigrateCommand,
	rollbackCommand RollbackCommand,
	autoMigrateCommand AutoMigrateCommand,
) BaseCommands {
	return BaseCommands{
		storageLinkCommand,
		migrateCommand,
		rollbackCommand,
		autoMigrateCommand,
	}
}

// Setup all the command
func (commands BaseCommands) Setup() {
	for _, cmd := range commands {
		cmd.Setup()
	}
}
