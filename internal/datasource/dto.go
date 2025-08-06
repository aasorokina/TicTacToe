package datasource

import (
	"tictactoe/internal/domain/game"

	"github.com/google/uuid"
)

// GameDTO is a Data Transfer Object for serializing and deserializing game.Game.
// It is used to convert the internal game state to a JSON-compatible format.
type GameDTO struct {
	ID     string                            `json:"gameID"`
	State  int                               `json:"state"`
	Grid   [game.GridSize][game.GridSize]int `json:"grid"`
	Winner int                               `json:"winner"`
}

// GameToDTO creates GameDTO struct from game.Game
func GameToDTO(g *game.Game) *GameDTO {
	dto := GameDTO{}
	dto.ID = g.ID.String()
	dto.State = int(g.State)
	dto.Winner = int(g.Winner)

	for i := 0; i < game.GridSize; i++ {
		for j := 0; j < game.GridSize; j++ {
			dto.Grid[i][j] = int(g.Grid[i][j])
		}
	}

	return &dto
}

// GameFromDTO creates game.Game struct from GameDTO
func GameFromDTO(dto *GameDTO) (*game.Game, error) {
	id, err := uuid.Parse(dto.ID)
	if err != nil {
		return nil, err
	}

	g := game.Game{}
	g.ID = id
	g.State = game.State(dto.State)
	g.Winner = game.Mark(dto.Winner)

	for i := 0; i < game.GridSize; i++ {
		for j := 0; j < game.GridSize; j++ {
			g.Grid[i][j] = game.Mark(dto.Grid[i][j])
		}
	}

	return &g, nil
}
