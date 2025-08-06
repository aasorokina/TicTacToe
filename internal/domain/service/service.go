package service

import (
	"tictactoe/internal/domain/game"
)

// GameService defines the interface for operations with game logic
type GameService interface {
	GetNextMove(game *game.Game, currentPlayer game.Mark) error
	ValidateField(old, updated *game.Game) error
	IsOver(game *game.Game) bool
	SaveGame(g *game.Game)
	GetGame(id string) (*game.Game, error)
	SaveGames() error
	GetAllGames() ([]*game.Game, error)
}
