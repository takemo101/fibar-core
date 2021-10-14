package boot

import (
	"testing"

	"github.com/takemo101/fibar-core/pkg"
	"github.com/takemo101/fibar-core/pkg/contract"
	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"
)

// TestOptions app boot options
type TestOptions struct {
	ConfigPath contract.ConfigPath
	FXOption   fx.Option
}

// Testing test func
func Testing(t *testing.T, options TestOptions, tests ...interface{}) {
	fxtest.New(
		t,
		fx.Provide(func() contract.ConfigPath {
			return options.ConfigPath
		}),
		options.FXOption,
		pkg.Module,
		fx.Invoke(tests...),
	).Done()
}
