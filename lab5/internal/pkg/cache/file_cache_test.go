package cache

import (
    "os"
    "testing"
    "time"
)

func TestFileCache_SetAndGet(t *testing.T) {
    tmpDir := "./test_cache"
    defer os.RemoveAll(tmpDir)
    
    cache, err := NewFileCache(tmpDir)
    if err != nil {
        t.Fatalf("Failed to create file cache: %v", err)
    }
    
    key := "test_key"
    value := map[string]string{"data": "test_value"}
    ttl := time.Second * 10
    
    err = cache.Set(key, value, ttl)
    if err != nil {
        t.Errorf("Failed to set cache: %v", err)
    }
    
    var result map[string]string
    err = cache.Get(key, &result)
    if err != nil {
        t.Errorf("Failed to get cache: %v", err)
    }
    
    if result["data"] != value["data"] {
        t.Errorf("Expected %s, got %s", value["data"], result["data"])
    }
}

func TestFileCache_Expiration(t *testing.T) {
    tmpDir := "./test_cache_expire"
    defer os.RemoveAll(tmpDir)
    
    cache, _ := NewFileCache(tmpDir)
    
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

func TestFileCache_Delete(t *testing.T) {
    tmpDir := "./test_cache_delete"
    defer os.RemoveAll(tmpDir)
    
    cache, _ := NewFileCache(tmpDir)
    
    key := "delete_key"
    cache.Set(key, "value", time.Second*10)
    
    cache.Delete(key)
    
    var result string
    err := cache.Get(key, &result)
    if err == nil {
        t.Error("Key should not exist after deletion")
    }
}
