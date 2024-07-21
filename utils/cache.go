package utils

import (
	"sync"
	"time"
)

type ICache interface {
	Set(string, interface{}, time.Duration) error
	Get(string) (interface{}, error)
}

var DefaultMemoryCache ICache = &memoryCache{
	data: make(map[string]interface{}),
}

type memoryCache struct {
	mu   sync.RWMutex
	data map[string]interface{}
}

func (m *memoryCache) Set(key string, value interface{}, duration time.Duration) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.data == nil {
		m.data = make(map[string]interface{})
	}
	m.data[key] = value
	return nil
}

func (m *memoryCache) Get(key string) (interface{}, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	if m.data == nil {
		return nil, nil
	}
	value, ok := m.data[key]
	if ok {
		return value, nil
	}
	return nil, nil
}
