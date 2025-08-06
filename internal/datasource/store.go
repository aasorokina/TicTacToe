package datasource

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"tictactoe/internal/domain/game"

	"github.com/google/uuid"
)

// GamesListFile is the path to the file that stores the list of saved games.
const GamesListFile = "internal/datasource/games_list.json"

// GameStore is a concurrent in-memory storage for active games.
// It uses sync.Map to safely handle concurrent read/write operations.
type GameStore struct {
	games sync.Map
}

// NewGameStore creates a new instance of GameStore.
func NewGameStore() GameStore {
	return GameStore{}
}

// SaveGame stores the given game in the GameStore using its ID as the key.
func (s *GameStore) SaveGame(g *game.Game) {
	s.games.Store(g.ID, g)
}

// GetGame retrieves a game by its UUID from the GameStore.
// Returns an error if the game is not found or if the stored value has an unexpected type.
func (s *GameStore) GetGame(id uuid.UUID) (*game.Game, error) {
	value, ok := s.games.Load(id)
	if !ok {
		return nil, fmt.Errorf("game not found")
	}

	gm, ok := value.(*game.Game)
	if !ok {
		return nil, fmt.Errorf("invalid type in store")
	}

	return gm, nil
}

// GetAllGames returns a slice of all stored games from the GameStore.
// Returns an error if the store is empty.
func (s *GameStore) GetAllGames() ([]*game.Game, error) {
	var games []*game.Game

	s.games.Range(func(key, value any) bool {
		if gameValue, ok := value.(*game.Game); ok && gameValue != nil {
			games = append(games, gameValue)
		}
		return true
	})

	if len(games) == 0 {
		return nil, fmt.Errorf("storage is empty")
	}

	return games, nil
}

// LoadGamesFromJSON reads games from a JSON file and loads them into the GameStore.
// Returns an error if the file can't be read or JSON unmarshalling fails.
func (s *GameStore) LoadGamesFromJSON() error {
	data, err := os.ReadFile(GamesListFile)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	var dtos []*GameDTO
	if err := json.Unmarshal(data, &dtos); err != nil {
		return fmt.Errorf("failed to unmarshal json: %w", err)
	}

	for _, dto := range dtos {
		id, err := uuid.Parse(dto.ID)
		if err != nil {
			continue
		}
		game, err := GameFromDTO(dto)
		if err == nil {
			s.games.Store(id, game)
		}
	}

	return nil
}

// SaveGamesToJSON serializes all games from the GameStore into JSON format
// and writes them to the file specified by GamesListFile.
// Returns an error if JSON marshalling or file writing fails.
func (s *GameStore) SaveGamesToJSON() error {
	var dtos []*GameDTO

	s.games.Range(func(key, value any) bool {
		gameValue := value.(*game.Game)
		dto := GameToDTO(gameValue)
		if dto != nil {
			dtos = append(dtos, dto)
		}
		return true
	})

	data, err := json.MarshalIndent(dtos, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(GamesListFile, data, 0644)
}
