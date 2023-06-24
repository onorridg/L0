package inMemory

import (
	"l0/internal/env"
	"l0/internal/postgresql"
	"sync"
	"time"
)

var dataBase *InMemory

type cache struct {
	row       any
	updatedAt time.Time
}

type InMemory struct {
	mu    sync.RWMutex
	cache map[uint64]cache
}

func (m *InMemory) evict() {
	var cacheRowTime time.Time
	var cacheRowId uint64

	m.mu.Lock()
	defer m.mu.Unlock()
	for id, row := range m.cache {
		if cacheRowId == 0 {
			cacheRowId, cacheRowTime = id, row.updatedAt
		} else if cacheRowTime.Before(row.updatedAt) {
			cacheRowId, cacheRowTime = id, row.updatedAt
		}
	}
	delete(m.cache, cacheRowId)
}

func (m *InMemory) restoreDataFromPostgres() {
	db := postgresql.Conn()
	defer db.Conn.Close()

	orders := db.GetLastNOrders()
	if len(orders) == 0 {
		return
	}

	m.mu.Lock()
	defer m.mu.Unlock()
	for _, order := range orders {
		m.cache[order.Id] = cache{row: &order, updatedAt: time.Now()}
	}
}

func (m *InMemory) QueryOrder(id uint64) any {
	m.mu.RLock()
	defer m.mu.RUnlock()

	data, exist := m.cache[id]
	if exist {
		data.updatedAt = time.Now()
		return data.row
	}
	return nil
}

func (m *InMemory) Append(id uint64, data any) {
	m.mu.Lock()
	m.cache[id] = cache{row: &data, updatedAt: time.Now()}
	m.mu.Unlock()

	if float64(len(m.cache)) > float64(env.Get().CacheSize)*0.2 {
		m.evict()
	}
}

func Conn() *InMemory {
	if dataBase == nil {
		dataBase = &InMemory{}
		dataBase.cache = make(map[uint64]cache, env.Get().CacheSize)
		dataBase.restoreDataFromPostgres()
	}
	return dataBase
}
