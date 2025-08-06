package game

import (
	"fmt"

	"github.com/google/uuid"
)

// GridSize - default numbers of rows and cols in TicTacToe game
const GridSize = 3

// Mark represents a player's mark or state on the game board
type Mark int

// Constants representing the possible marks on the game board
const (
	Empty  Mark = iota // Empty cell on the board
	Cross              // Player's mark: Cross (X)
	Nought             // Player's mark: Nought (O)
)

// State represents the current status of the game
type State int

// Constants representing the possible game states
const (
	InProgress State = iota // game in progress
	Completed               // game ended
)

// Coord represents a coordinate on the game board with row and column
type Coord struct {
	Row, Col int
}

// NoCoord represents an invalid or undefined coordinate
var NoCoord = Coord{Row: -1, Col: -1}

// Grid represents the game board as a 2D array of Marks
type Grid [GridSize][GridSize]Mark

// Game represents data about specified TicTacToe game instance
type Game struct {
	Grid   Grid      // Current state of the board
	ID     uuid.UUID // Unique identifier of the game
	State  State     // Current state of the game
	Winner Mark      // The winner mark
}

// NewGame returns a new Game instance with initialized values
func NewGame() *Game {
	return &Game{
		ID:     uuid.New(),
		Grid:   Grid{},
		State:  InProgress,
		Winner: Empty,
	}
}

// IsOver checks if the game is finished (there is a horizontal, vertical or diagonal row of same symbols),
// returns the finish status and the winner if there is one.
// If there is no winner (a draw or the game is not finished yet), returns Empty mark
func (g *Game) IsOver() (bool, Mark) {
	for i := 0; i < GridSize; i++ {
		if g.Grid[i][0] != Empty && g.Grid[i][0] == g.Grid[i][1] && g.Grid[i][1] == g.Grid[i][2] {
			return true, g.Grid[i][0]
		}

		if g.Grid[0][i] != Empty && g.Grid[0][i] == g.Grid[1][i] && g.Grid[1][i] == g.Grid[2][i] {
			return true, g.Grid[0][i]
		}
	}

	if g.Grid[0][0] != Empty && g.Grid[0][0] == g.Grid[1][1] && g.Grid[1][1] == g.Grid[2][2] {
		return true, g.Grid[0][0]
	}

	if g.Grid[0][2] != Empty && g.Grid[0][2] == g.Grid[1][1] && g.Grid[1][1] == g.Grid[2][0] {
		return true, g.Grid[0][2]
	}

	for i := 0; i < GridSize; i++ {
		for j := 0; j < GridSize; j++ {
			if g.Grid[i][j] == Empty {
				return false, Empty
			}
		}
	}

	return true, Empty
}

// NextMove returns next move from computer, calculated by minimax algorithm,
// and error if next move is not possible
func (g *Game) NextMove(currentPlayer Mark) (Coord, error) {
	_, bestCoord := Minimax(g, currentPlayer, 0)

	if bestCoord.Row == NoCoord.Row || bestCoord.Col == NoCoord.Col {
		return NoCoord, fmt.Errorf("could not find a valid move")
	}

	return bestCoord, nil
}

// SetPlayerMove checks player move coord and set it to game
// board if move is correct, else returns error
func (g *Game) SetPlayerMove(move Coord, currentPlayer Mark) error {
	gameOver, _ := g.IsOver()
	if gameOver {
		g.State = Completed
		return fmt.Errorf("no move possible: game is over")
	}

	if (move.Row < 0 || move.Row >= GridSize) || (move.Col < 0 || move.Col >= GridSize) {
		return fmt.Errorf("no move possible: no such cell")
	}

	if g.Grid[move.Row][move.Col] != Empty {
		return fmt.Errorf("no move possible: cell is occupied")
	}

	g.Grid[move.Row][move.Col] = currentPlayer
	return nil
}
