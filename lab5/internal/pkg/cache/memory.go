package cache

import (
    "encoding/json"
    "errors"
    "sync"
    "time"
)

type memoryItem struct {
    Value      []byte
    Expiration time.Time
}

type MemoryCache struct {
    data map[string]memoryItem
    mu   sync.RWMutex
}

func NewMemoryCache() *MemoryCache {
    return &MemoryCache{
        data: make(map[string]memoryItem),
    }
}

func (c *MemoryCache) Set(key string, value interface{}, ttl time.Duration) error {
    c.mu.Lock()
    defer c.mu.Unlock()

    data, err := json.Marshal(value)
    if err != nil {
        return err
    }

    c.data[key] = memoryItem{
        Value:      data,
        Expiration: time.Now().Add(ttl),
    }
    return nil
}

func (c *MemoryCache) Get(key string, value interface{}) error {
    c.mu.RLock()
    defer c.mu.RUnlock()

    item, exists := c.data[key]
    if !exists {
        return errors.New("key not found")
    }

    if time.Now().After(item.Expiration) {
        delete(c.data, key)
        return errors.New("key expired")
    }

    return json.Unmarshal(item.Value, value)
}

func (c *MemoryCache) Delete(key string) error {
    c.mu.Lock()
    defer c.mu.Unlock()
    delete(c.data, key)
    return nil
}

func (c *MemoryCache) Clear() error {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.data = make(map[string]memoryItem)
    return nil
}

func (c *MemoryCache) Exists(key string) bool {
    c.mu.RLock()
    defer c.mu.RUnlock()

    item, exists := c.data[key]
    if !exists {
        return false
    }
    return !time.Now().After(item.Expiration)
}

func (c *MemoryCache) Close() error {
    c.Clear()
    return nil
}
