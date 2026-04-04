package config

import (
    "strings"
    "testing"
    "time"
)

func TestParse(t *testing.T) {
    yamlData := `
service:
  provider:
    type: open-meteo
  location:
    lat: 53.6688
    long: 23.8223
  cache:
    type: memory
    ttl: 300
    dir: ./cache
`
    reader := strings.NewReader(yamlData)
    
    cfg, err := Parse(reader)
    if err != nil {
        t.Errorf("Failed to parse config: %v", err)
    }
    
    if cfg.P.Type != "open-meteo" {
        t.Errorf("Expected provider type 'open-meteo', got '%s'", cfg.P.Type)
    }
    
    if cfg.L.Lat != 53.6688 {
        t.Errorf("Expected lat 53.6688, got %f", cfg.L.Lat)
    }
    
    if cfg.L.Long != 23.8223 {
        t.Errorf("Expected long 23.8223, got %f", cfg.L.Long)
    }
    
    if cfg.C.Type != "memory" {
        t.Errorf("Expected cache type 'memory', got '%s'", cfg.C.Type)
    }
    
    if cfg.C.TTL != 300 {
        t.Errorf("Expected TTL 300, got %d", cfg.C.TTL)
    }
}

func TestParse_DefaultValues(t *testing.T) {
    yamlData := `
service:
  provider:
    type: open-meteo
  location:
    lat: 53.6688
    long: 23.8223
`
    reader := strings.NewReader(yamlData)
    
    cfg, err := Parse(reader)
    if err != nil {
        t.Errorf("Failed to parse config: %v", err)
    }
    
    if cfg.C.Type != "memory" {
        t.Errorf("Expected default cache type 'memory', got '%s'", cfg.C.Type)
    }
    
    if cfg.C.TTL != 300 {
        t.Errorf("Expected default TTL 300, got %d", cfg.C.TTL)
    }
}

func TestGetTTL(t *testing.T) {
    yamlData := `
service:
  provider:
    type: open-meteo
  location:
    lat: 53.6688
    long: 23.8223
  cache:
    type: memory
    ttl: 60
    dir: ./cache
`
    reader := strings.NewReader(yamlData)
    
    cfg, _ := Parse(reader)
    
    ttl := cfg.GetTTL()
    expected := time.Duration(60) * time.Second
    
    if ttl != expected {
        t.Errorf("Expected TTL %v, got %v", expected, ttl)
    }
}
