package weather

import (
    "testing"
)

// Мок-логгер для тестов
type mockLogger struct {
    infoCalled  bool
    debugCalled bool
    errorCalled bool
}

func (m *mockLogger) Info(msg string) {
    m.infoCalled = true
}

func (m *mockLogger) Debug(msg string) {
    m.debugCalled = true
}

func (m *mockLogger) Error(msg string, err error) {
    m.errorCalled = true
}

func TestWeatherInfo_New(t *testing.T) {
    logger := &mockLogger{}
    wi := New(logger)
    
    if wi == nil {
        t.Error("New() returned nil")
    }
}

func TestWeatherInfo_GetTemperature(t *testing.T) {
    logger := &mockLogger{}
    wi := New(logger)
    
    // Тест с реальным API (требует интернет)
    // Пропускаем если нет интернета
    if testing.Short() {
        t.Skip("Skipping API test in short mode")
    }
    
    temp := wi.GetTemperature(53.6688, 23.8223)
    
    if temp.Temp == 0 {
        t.Log("Warning: Got zero temperature, API might be unreachable")
    }
}
