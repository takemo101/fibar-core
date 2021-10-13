package boot

import (
	"testing"

	"github.com/takemo101/fibar-core/app"
	"github.com/takemo101/fibar-core/pkg"
	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"
)

// Testing test func
func Testing(t *testing.T, tests ...interface{}) {
	fxtest.New(
		t,
		pkg.Module,
		app.Module,
		fx.Invoke(tests...),
	).Done()
}
