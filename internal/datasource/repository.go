package datasource

import (
	game "tictactoe/internal/domain/game"

	"github.com/google/uuid"
)

// GameRepository is interface for interacting with the games storage structure
type GameRepository interface {
	SaveGame(g *game.Game)
	GetGame(id uuid.UUID) (*game.Game, error)
	SaveGames() error
	GetAllGames() ([]*game.Game, error)
}
