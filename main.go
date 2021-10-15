package main

import (
	"github.com/takemo101/fibar-core/app"
	"github.com/takemo101/fibar-core/boot"
	"github.com/takemo101/fibar-core/pkg"
	"github.com/takemo101/fibar-core/pkg/contract"
	"go.uber.org/fx"
)

// AppBooter is module root struct
type AppBooter struct {
	app pkg.Application
}

// AppBoot all setup
func (booter AppBooter) AppBoot() {
	//
}

// NewAppBooter app create
func NewAppBooter(app pkg.Application) contract.AppBooter {
	return AppBooter{
		app: app,
	}
}

func main() {

	// boot gin application
	boot.Run(
		boot.AppOptions{
			ConfigPath:           "config.yml",
			AppBooterConstructor: NewAppBooter,
			FXOption: fx.Options(
				app.Module,
			),
		},
	)
}
