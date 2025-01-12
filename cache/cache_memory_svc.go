package cache

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"
)

type memoryCache struct {
	data map[string]cacheItem
	mu   sync.RWMutex
}

type cacheItem struct {
	value      []byte
	expiration int64
}

func NewMemoryCache() Cache {
	return &memoryCache{
		data: make(map[string]cacheItem),
	}
}

func (m *memoryCache) Set(key string, value interface{}, duration int) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	jsonData, err := json.Marshal(value)
	if err != nil {
		return err
	}

	m.data[key] = cacheItem{
		value:      jsonData,
		expiration: time.Now().Add(time.Duration(duration) * time.Second).Unix(),
	}
	return nil
}

func (m *memoryCache) Get(key string, receiver interface{}) error {
	m.mu.RLock()
	defer m.mu.RUnlock()

	item, found := m.data[key]
	if !found || time.Now().Unix() > item.expiration {
		return fmt.Errorf("key not found or expired")
	}

	return json.Unmarshal(item.value, receiver)
}
