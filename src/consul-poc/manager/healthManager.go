package manager

import (
	"sync"
)

type HealthManager struct {
	health bool
	mutex  *sync.RWMutex
}

func (manager *HealthManager) SetHealth(health bool) {
	manager.mutex.Lock()
	defer manager.mutex.Unlock()

	manager.health = health
}

func (manager *HealthManager) IsHealthy() bool {
	manager.mutex.RLock()
	defer manager.mutex.RUnlock()

	return manager.health
}

var healthManager *HealthManager
var once sync.Once

func GetHealthManager() *HealthManager {
	once.Do(func() {
		healthManager = &HealthManager{true, &sync.RWMutex{}}
	})

	return healthManager
}
