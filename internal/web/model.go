package web

import "tictactoe/internal/domain/game"

// MoveRequest represents a player's move on the game grid
type MoveRequest struct {
	Row int `json:"row"` // Row index (0-based)
	Col int `json:"col"` // Column index (0-based)
}

// GameResponse is the JSON-serializable representation of a game state
type GameResponse struct {
	ID     string                            `json:"gameID"`
	State  int                               `json:"state"`
	Grid   [game.GridSize][game.GridSize]int `json:"grid"`
	Winner int                               `json:"winner"`
}
