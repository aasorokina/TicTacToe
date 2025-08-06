package di

import (
	"go.uber.org/fx"

	"tictactoe/internal/datasource"
	"tictactoe/internal/domain/service"
	"tictactoe/internal/web"
)

// FxConfig defines and provides all the application dependencies using fx.Provide.
// It wires up the GameStore, GameRepository, GameService, GameHandler, and Gin router.
// This configuration is used to construct the application's dependency graph.
func FxConfig() fx.Option {
	opt := fx.Provide(
		datasource.NewGameStore,
		datasource.NewGameRepository,
		service.NewGameService,
		web.NewGameHandler,
		web.NewRouter,
	)
	return opt
}
