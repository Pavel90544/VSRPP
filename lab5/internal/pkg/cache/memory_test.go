package cache

import (
    "testing"
    "time"
)

func TestMemoryCache_SetAndGet(t *testing.T) {
    cache := NewMemoryCache()
    defer cache.Close()

    key := "test_key"
    expected := "test_value"
    ttl := time.Second * 10

    err := cache.Set(key, expected, ttl)
    if err != nil {
        t.Fatalf("Set failed: %v", err)
    }

    var result string
    err = cache.Get(key, &result)
    if err != nil {
        t.Fatalf("Get failed: %v", err)
    }

    if result != expected {
        t.Errorf("Expected %s, got %s", expected, result)
    }
}

func TestMemoryCache_GetNotFound(t *testing.T) {
    cache := NewMemoryCache()
    defer cache.Close()

    var result string
    err := cache.Get("nonexistent", &result)

    if err == nil {
        t.Error("Expected error for nonexistent key")
    }
}

func TestMemoryCache_Expiration(t *testing.T) {
    cache := NewMemoryCache()
    defer cache.Close()

    key := "expire_key"
    value := "expire_value"
    ttl := time.Millisecond * 100

    err := cache.Set(key, value, ttl)
    if err != nil {
        t.Fatalf("Set failed: %v", err)
    }

    time.Sleep(ttl + time.Millisecond*50)

    var result string
    err = cache.Get(key, &result)

    if err == nil {
        t.Error("Expected error for expired key")
    }
}

func TestMemoryCache_Delete(t *testing.T) {
    cache := NewMemoryCache()
    defer cache.Close()

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
    defer cache.Close()

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
    defer cache.Close()

    key := "exists_key"
    cache.Set(key, "value", time.Second*10)

    if !cache.Exists(key) {
        t.Error("Exists should return true for existing key")
    }

    cache.Delete(key)

    if cache.Exists(key) {
        t.Error("Exists should return false after deletion")
    }
}
