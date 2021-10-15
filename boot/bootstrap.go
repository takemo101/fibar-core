package boot

import (
	"context"

	"github.com/takemo101/fibar-core/pkg"
	"github.com/takemo101/fibar-core/pkg/contract"
	"go.uber.org/fx"
)

// AppOptions app boot options
type AppOptions struct {
	ConfigPath           contract.ConfigPath
	AppBooterConstructor interface{}
	FXOption             fx.Option
}

// boot is initialize application
func boot(
	lifecycle fx.Lifecycle,
	app pkg.Application,
	logger pkg.Logger,
	database pkg.Database,
	booter contract.AppBooter,
) {
	sql, err := database.DB()
	if err != nil {
		logger.Info("database connection sql failed : %v", err)
	}

	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			logger.Info("-- start application --")

			sql.SetMaxOpenConns(app.Config.DB.Connection.Max)
			go func() {
				booter.AppBoot()
				app.Run()
			}()
			return nil
		},
		OnStop: func(context.Context) error {
			logger.Info("-- stop application --")
			sql.Close()
			return nil
		},
	})
}

// app run
func Run(options AppOptions) {
	opts := []fx.Option{
		options.FXOption,
		pkg.Module,
		fx.Provide(
			func() contract.ConfigPath {
				return options.ConfigPath
			},
			options.AppBooterConstructor,
		),
		fx.Invoke(boot),
	}

	fx.New(opts...).Run()
}
