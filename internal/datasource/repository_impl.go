package datasource

import (
	game "tictactoe/internal/domain/game"

	"github.com/google/uuid"
)

type gameRepository struct {
	Storage GameStore
}

// NewGameRepository creates a new instance of GameRepository.
func NewGameRepository() GameRepository {
	repo := gameRepository{
		Storage: NewGameStore(),
	}
	repo.Storage.LoadGamesFromJSON()
	return &repo
}

// SaveGame is saving game in storage
func (r *gameRepository) SaveGame(g *game.Game) {
	r.Storage.SaveGame(g)
}

// GetGame is loading game by uuid from storage
func (r *gameRepository) GetGame(id uuid.UUID) (*game.Game, error) {
	return r.Storage.GetGame(id)
}

// GetAllGames returns slice of all saved in storage games.
// Returns an error if the store is empty.
func (r *gameRepository) GetAllGames() ([]*game.Game, error) {
	return r.Storage.GetAllGames()
}

// LoadGames is loading all games from json to storage
func (r *gameRepository) LoadGames() error {
	return r.Storage.LoadGamesFromJSON()
}

// SaveGames is saving all games from storage to json
func (r *gameRepository) SaveGames() error {
	return r.Storage.SaveGamesToJSON()
}
