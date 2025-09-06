package storage

import (
	"sync"
	"time"

	"ROMAX/internal/models"
)

var (
	players = map[string]*models.Player{}
	mu      sync.RWMutex
)

func Init() {
	CreatePlayer("test_player_1")
}

func CreatePlayer(id string) *models.Player {
	mu.Lock()
	defer mu.Unlock()
	if p, ok := players[id]; ok {
		return p
	}
	p := &models.Player{
		ID:                 id,
		Balance:            1000.0,
		FreeSpinsRemaining: 0,
		History:            []models.Round{},
		CreatedAt:          time.Now(),
	}
	players[id] = p
	return p
}

func GetPlayer(id string) (*models.Player, bool) {
	mu.RLock()
	defer mu.RUnlock()
	p, ok := players[id]
	return p, ok
}

func SavePlayer(p *models.Player) {
	mu.Lock()
	defer mu.Unlock()
	players[p.ID] = p
}
