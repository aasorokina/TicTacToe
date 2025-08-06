package web

import (
	"github.com/gin-gonic/gin"
)

// NewRouter sets up the HTTP routes for the Tic Tac Toe game API using Gin.
// It registers endpoints for retrieving games, making moves, and saving game data.
func NewRouter(h *GameHandler) *gin.Engine {
	router := gin.Default()
	router.GET("/tictactoe/games", h.GetAllGames)
	router.GET("/tictactoe/games/:id", h.GetGameByID)
	router.POST("/tictactoe/games/save", h.SaveAllGames)
	router.POST("/tictactoe/games/:id/move", h.ProcessMove)
	router.POST("/tictactoe/games/move", h.ProcessMove)

	return router
}
