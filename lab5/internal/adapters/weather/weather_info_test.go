package weather

import (
    "os"
    "testing"
)

// mockLogger - мок-логгер для тестов
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

func TestWeatherInfo_GetTemperature_RealAPI(t *testing.T) {
    // Пропускаем тест в CI окружении, так как API может быть недоступен
    if os.Getenv("CI") != "" {
        t.Skip("Skipping real API test in CI environment")
    }

    logger := &mockLogger{}
    wi := New(logger)

    // Тест с реальным API (требует интернет)
    temp, err := wi.GetTemperature(53.6688, 23.8223)

    if err != nil {
        t.Logf("API error (may be network issue): %v", err)
        // Не считаем ошибкой, так как сеть может быть недоступна
        return
    }

    if temp.Temp == 0 {
        t.Log("Warning: Got zero temperature, API might be unreachable")
    }

    t.Logf("Temperature from real API: %.2f°C", temp.Temp)
}

func TestWeatherInfo_GetTemperature_Cached(t *testing.T) {
    logger := &mockLogger{}
    wi := New(logger)

    // Первый запрос - загружает данные
    temp1, err1 := wi.GetTemperature(53.6688, 23.8223)
    if err1 != nil {
        t.Logf("First request error (may be network): %v", err1)
    }

    // Второй запрос - должен вернуть кэшированные данные
    temp2, err2 := wi.GetTemperature(53.6688, 23.8223)

    if err1 == nil && err2 == nil {
        if temp1.Temp != temp2.Temp {
            t.Errorf("Cached temperature mismatch: got %.2f, expected %.2f", temp2.Temp, temp1.Temp)
        }
    }
}
