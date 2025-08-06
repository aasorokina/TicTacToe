package game

import "math"

// Minimax implements the Minimax algorithm to evaluate the best move for the current player.
// It recursively simulates all possible game outcomes and chooses the move that leads to the best result.
// Returns the best score and the coordinate of the optimal move.
func Minimax(g *Game, currentPlayer Mark, depth int) (int, Coord) {
	opponent := GetOpponent(currentPlayer)
	bestScore := GetBestScore(currentPlayer)
	bestCoord := NoCoord

	gameOver, winner := g.IsOver()
	if gameOver {
		return CalculateWinPoints(winner, depth), bestCoord
	}

	depth++

	for i := 0; i < GridSize; i++ {
		for j := 0; j < GridSize; j++ {
			if g.Grid[i][j] == Empty {
				currCoord := Coord{i, j}
				g.Grid[i][j] = currentPlayer
				score, _ := Minimax(g, opponent, depth)
				g.Grid[i][j] = Empty

				if UpdateBestScore(currentPlayer, score, &bestScore) == true {
					bestCoord = currCoord
				}
			}
		}
	}

	return bestScore, bestCoord
}

// CalculateWinPoints returns a score based on the winner and current depth.
// Used in the Minimax algorithm to evaluate terminal game states.
func CalculateWinPoints(winner Mark, depth int) int {
	if winner == Cross {
		return 10 - depth
	} else if winner == Nought {
		return depth - 10
	} else {
		return 0
	}
}

// GetBestScore initializes and returns the starting best score depending on the current player.
// For Cross (maximizing), it returns the smallest integer value; for Nought (minimizing), the largest.
func GetBestScore(currentPlayer Mark) int {
	var bestScore int
	if currentPlayer == Cross {
		bestScore = math.MinInt
	} else {
		bestScore = math.MaxInt
	}
	return bestScore
}

// GetOpponent returns the mark opposite to the current player
func GetOpponent(currentPlayer Mark) Mark {
	var opponent Mark
	if currentPlayer == Cross {
		opponent = Nought
	} else {
		opponent = Cross
	}
	return opponent
}

// UpdateBestScore compares the given score with the current bestScore based on the player's role.
// For Cross (maximizing player), it updates if the score is greater.
// For Nought (minimizing player), it updates if the score is lower.
// Returns true if the bestScore was updated, false otherwise.
func UpdateBestScore(currentPlayer Mark, score int, bestScore *int) bool {
	if currentPlayer == Cross {
		if score > *bestScore {
			*bestScore = score
			return true
		}
	} else {
		if score < *bestScore {
			*bestScore = score
			return true
		}
	}
	return false
}
