package web

import "tictactoe/internal/domain/game"

// ToMoveRequest creates a MoveRequest from a game state and coordinate
func ToMoveRequest(g *game.Game, coord game.Coord) MoveRequest {
	return MoveRequest{
		Row: coord.Row,
		Col: coord.Col,
	}
}

// ToGameResponse converts a game.Game instance into a GameResponse.
func ToGameResponse(g *game.Game) GameResponse {
	gr := GameResponse{}
	gr.ID = g.ID.String()
	gr.State = int(g.State)
	gr.Winner = int(g.Winner)

	for i := 0; i < game.GridSize; i++ {
		for j := 0; j < game.GridSize; j++ {
			gr.Grid[i][j] = int(g.Grid[i][j])
		}
	}

	return gr
}
