package app

import (
	"tictactoe/internal/di"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

// NewFxApp creates and returns new fx.App instance.
// It initializes the application's dependencies using di.FxConfig()
// and starts the HTTP server on port 8080 with fx.Invoke.
func NewFxApp() *fx.App {
	app := fx.New(
		di.FxConfig(),
		fx.Invoke(func(router *gin.Engine) {
			router.Run(":8080")
		}),
	)
	return app
}
