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

    cache.Set(key, value, ttl)
    time.Sleep(ttl + time.Millisecond*50)

    var result string
    err := cache.Get(key, &result)

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

func TestMemoryCache_Concurrent(t *testing.T) {
    cache := NewMemoryCache()
    defer cache.Close()

    key := "concurrent_key"
    expected := "concurrent_value"
    ttl := time.Second * 10

    cache.Set(key, expected, ttl)

    done := make(chan bool)
    for i := 0; i < 10; i++ {
        go func() {
            var result string
            err := cache.Get(key, &result)
            if err != nil {
                t.Errorf("Concurrent get failed: %v", err)
            }
            if result != expected {
                t.Errorf("Expected %s, got %s", expected, result)
            }
            done <- true
        }()
    }

    for i := 0; i < 10; i++ {
        <-done
    }
}

func TestMemoryCache_Update(t *testing.T) {
    cache := NewMemoryCache()
    defer cache.Close()

    key := "update_key"
    ttl := time.Second * 10

    cache.Set(key, "first_value", ttl)
    cache.Set(key, "updated_value", ttl)

    var result string
    cache.Get(key, &result)

    if result != "updated_value" {
        t.Errorf("Expected updated_value, got %s", result)
    }
}

// ZeroTTL в MemoryCache означает бесконечное хранение
// (если реализовано именно так) или нужно явно передавать большой TTL
func TestMemoryCache_LongTTL(t *testing.T) {
    cache := NewMemoryCache()
    defer cache.Close()

    key := "long_ttl_key"
    value := "long_ttl_value"

    // Используем большой TTL вместо нулевого
    err := cache.Set(key, value, time.Hour*24)
    if err != nil {
        t.Fatalf("Set failed: %v", err)
    }

    time.Sleep(time.Millisecond * 100)

    var result string
    err = cache.Get(key, &result)
    if err != nil {
        t.Errorf("Get after long TTL failed: %v", err)
    }

    if result != value {
        t.Errorf("Expected %s, got %s", value, result)
    }
}
