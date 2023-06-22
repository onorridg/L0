package inMemory

import (
	"sync"
)

var dataBase *InMemory

type InMemory struct {
	mu    sync.RWMutex
	Cache map[uint64]interface{}
}

func (m *InMemory) restoreDataFromPostgres() {
	// PASS
}

func (m *InMemory) QueryData(id uint64) any {
	m.mu.RLock()
	defer m.mu.RUnlock()

	data, exist := m.Cache[id]
	if exist {
		return data
	}
	return nil
}

func (m *InMemory) InsertData(id uint64, data any) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.Cache[id] = data
}

func Conn() *InMemory {
	if dataBase == nil {
		dataBase = &InMemory{}
		dataBase.Cache = make(map[uint64]interface{}, 1000)
		//dataBase.restoreDataFromPostgres()
	}
	return dataBase
}
