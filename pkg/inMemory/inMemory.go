package inMemory

import (
	"l0/internal/env"
	"l0/internal/postgresql"
	"sync"
)

var dataBase *InMemory

type InMemory struct {
	mu    sync.RWMutex
	Cache map[uint64]interface{}
}

func (m *InMemory) restoreDataFromPostgres() {
	db := postgresql.Conn()
	defer db.Conn.Close()

	orders := db.GetLastNOrders()
	if len(orders) == 0 {
		return
	}

	for _, order := range orders {
		m.InsertData(order.Id, order)
	}
}

func (m *InMemory) QueryOrder(id uint64) any {
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
		dataBase.Cache = make(map[uint64]interface{}, env.Get().CacheSize)
		dataBase.restoreDataFromPostgres()
	}
	return dataBase
}
