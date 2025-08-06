package web

import (
	"net/http"
	"tictactoe/internal/domain/game"
	"tictactoe/internal/domain/service"

	"github.com/gin-gonic/gin"
)

// GameHandler handles HTTP requests related to tic-tac-toe games
type GameHandler struct {
	gameService service.GameService
}

// NewGameHandler creates a new GameHandler instance with GameService
func NewGameHandler(s service.GameService) *GameHandler {
	return &GameHandler{gameService: s}
}

// ProcessMove handles a POST request to make a move in a game.
// If no game ID is provided, it creates a new game.
// It validates the player's move, performs the opponent's move,
// checks for game over, and returns the updated game state.
func (h *GameHandler) ProcessMove(c *gin.Context) {
	strID := c.Param("id")
	if strID == "" {
		g := game.NewGame()
		h.gameService.SaveGame(g)
		strID = g.ID.String()
	}

	oldGame, err := h.gameService.GetGame(strID)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	var move MoveRequest
	if err := c.BindJSON(&move); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newGame := *oldGame
	errMove := newGame.SetPlayerMove(game.Coord{Row: move.Row, Col: move.Col}, game.Cross)
	if errMove != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": errMove.Error()})
		return
	}

	errValidate := h.gameService.ValidateField(oldGame, &newGame)
	if errValidate != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": errValidate.Error()})
		return
	}

	if isOver := h.gameService.IsOver(&newGame); isOver == true {
		c.IndentedJSON(http.StatusOK, ToGameResponse(&newGame))
		return
	}

	errNextMove := h.gameService.GetNextMove(&newGame, game.Nought)
	if errNextMove != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": errNextMove.Error()})
		return
	}

	h.gameService.IsOver(&newGame)
	h.gameService.SaveGame(&newGame)
	c.IndentedJSON(http.StatusOK, ToGameResponse(&newGame))
}

// GetAllGames handles a GET request to retrieve all saved games.
// Returns a list of game states or an error if no games are found
func (h *GameHandler) GetAllGames(c *gin.Context) {
	games, err := h.gameService.GetAllGames()
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	var res []GameResponse
	for i := 0; i < len(games); i++ {
		gr := ToGameResponse(games[i])
		res = append(res, gr)
	}

	c.IndentedJSON(http.StatusOK, res)
}

// GetGameByID handles a GET request to retrieve a specific game by its ID.
// Returns the game's current state or an error if not found
func (h *GameHandler) GetGameByID(c *gin.Context) {
	strID := c.Param("id")
	game, err := h.gameService.GetGame(strID)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusOK, ToGameResponse(game))
}

// SaveAllGames handles a POST request to persist all games currently stored in memory.
// Returns a success message or an error if saving fails
func (h *GameHandler) SaveAllGames(c *gin.Context) {
	err := h.gameService.SaveGames()

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "data saved succsessfully",
	})
}
