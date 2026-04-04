package flags

import (
    "os"
    "testing"
)

func TestParse_DefaultConfig(t *testing.T) {
    // Сохраняем оригинальные аргументы
    oldArgs := os.Args
    defer func() { os.Args = oldArgs }()
    
    // Устанавливаем тестовые аргументы (без флага config)
    os.Args = []string{"cmd"}
    
    f := Parse()
    
    if f.Path != "./config/config.yaml" {
        t.Errorf("Expected default path './config/config.yaml', got '%s'", f.Path)
    }
}

func TestParse_CustomConfig(t *testing.T) {
    // Сохраняем оригинальные аргументы
    oldArgs := os.Args
    defer func() { os.Args = oldArgs }()
    
    // Устанавливаем тестовые аргументы
    os.Args = []string{"cmd", "-config", "./custom/path/config.yaml"}
    
    f := Parse()
    
    if f.Path != "./custom/path/config.yaml" {
        t.Errorf("Expected './custom/path/config.yaml', got '%s'", f.Path)
    }
}
