package service

import (
	"fmt"
	"tictactoe/internal/datasource"
	"tictactoe/internal/domain/game"

	"github.com/google/uuid"
)

type gameService struct {
	repo datasource.GameRepository
}

// NewGameService creates a new instance of GameService with GameRepository
func NewGameService(r datasource.GameRepository) GameService {
	return &gameService{
		repo: r,
	}
}

// GetNextMove calculates and performs the next move for the given player using the minimax algorithm.
// Returns an error if the move cannot be determined or applied
func (s *gameService) GetNextMove(g *game.Game, currentPlayer game.Mark) error {
	coord, err := g.NextMove(currentPlayer)

	if err != nil {
		return fmt.Errorf("%v", err)
	}

	g.Grid[coord.Row][coord.Col] = currentPlayer
	return nil
}

// IsOver checks whether the game is over and sets the final state and winner.
// Saves the game to the repository and returns true if the game is over
func (s *gameService) IsOver(g *game.Game) bool {
	isOver, winner := g.IsOver()
	if isOver {
		g.State = game.Completed
		g.Winner = winner
	}
	s.SaveGame(g)
	return isOver
}

// ValidateField compares two game states and ensures that exactly one cell is different,
// indicating a valid move. Returns an error if the move is invalid
func (s *gameService) ValidateField(old, updated *game.Game) error {
	diffCount := 0

	for i := 0; i < game.GridSize; i++ {
		for j := 0; j < game.GridSize; j++ {
			if old.Grid[i][j] != updated.Grid[i][j] {
				diffCount++
			}
		}
	}

	if diffCount != 1 {
		return fmt.Errorf("invalid move on field")
	}
	return nil
}

// SaveGame saves the given game to the repository
func (s *gameService) SaveGame(g *game.Game) {
	s.repo.SaveGame(g)
}

// GetAllGames retrieves all games from the repository
func (s *gameService) GetGame(strID string) (*game.Game, error) {
	id, err := uuid.Parse(strID)

	if err != nil {
		return nil, fmt.Errorf("failed to parse uuid")
	}

	return s.repo.GetGame(id)
}
func (s *gameService) GetAllGames() ([]*game.Game, error) {
	return s.repo.GetAllGames()
}

// SaveGames save all games from repository to json file
func (s *gameService) SaveGames() error {
	return s.repo.SaveGames()
}
