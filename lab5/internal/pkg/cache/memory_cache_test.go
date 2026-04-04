package cache

import (
    "testing"
    "time"
)

func TestMemoryCache_SetAndGet(t *testing.T) {
    cache := NewMemoryCache()

    key := "test_key"
    value := "test_value"
    ttl := time.Second * 10

    err := cache.Set(key, value, ttl)
    if err != nil {
        t.Errorf("Failed to set cache: %v", err)
    }

    var result string
    err = cache.Get(key, &result)
    if err != nil {
        t.Errorf("Failed to get cache: %v", err)
    }

    if result != value {
        t.Errorf("Expected %s, got %s", value, result)
    }
}

func TestMemoryCache_Expiration(t *testing.T) {
    cache := NewMemoryCache()

    key := "expire_key"
    value := "expire_value"
    ttl := time.Millisecond * 100

    cache.Set(key, value, ttl)
    time.Sleep(ttl + time.Millisecond*50)

    var result string
    err := cache.Get(key, &result)
    if err == nil {
        t.Error("Expected error for expired key, got nil")
    }
}

func TestMemoryCache_Delete(t *testing.T) {
    cache := NewMemoryCache()

    key := "delete_key"
    cache.Set(key, "value", time.Second*10)

    if !cache.Exists(key) {
        t.Error("Key should exist before deletion")
    }

    cache.Delete(key)

    if cache.Exists(key) {
        t.Error("Key should not exist after deletion")
    }
}

func TestMemoryCache_Clear(t *testing.T) {
    cache := NewMemoryCache()

    cache.Set("key1", "value1", time.Second*10)
    cache.Set("key2", "value2", time.Second*10)
    cache.Set("key3", "value3", time.Second*10)

    cache.Clear()

    if cache.Exists("key1") || cache.Exists("key2") || cache.Exists("key3") {
        t.Error("Cache should be empty after Clear()")
    }
}

func TestMemoryCache_Exists(t *testing.T) {
    cache := NewMemoryCache()

    key := "exists_key"
    cache.Set(key, "value", time.Second*10)

    if !cache.Exists(key) {
        t.Error("Exists() should return true for existing key")
    }

    cache.Delete(key)

    if cache.Exists(key) {
        t.Error("Exists() should return false after deletion")
    }
}
