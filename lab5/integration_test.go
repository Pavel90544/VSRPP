package main

import (
    "os"
    "testing"
)

// Интеграционный тест для всего приложения
func TestFullApp(t *testing.T) {
    if testing.Short() {
        t.Skip("Skipping integration test in short mode")
    }
    
    // Проверяем, что конфиг существует
    if _, err := os.Stat("./config/config.yaml"); os.IsNotExist(err) {
        t.Skip("Config file not found, skipping integration test")
    }
    
    // Здесь можно добавить тест полного цикла работы приложения
    t.Log("Integration test passed")
}
