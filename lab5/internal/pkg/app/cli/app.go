package cli

import (
    "fmt"

    "github.com/Pavel90544/VSRPP/lab5/internal/domain/models"
    "github.com/Pavel90544/VSRPP/lab5/pkg/config"
)

type Logger interface {
    Info(string)
    Debug(string)
    Error(string, error)
}

type WeatherInfo interface {
    GetTemperature(float64, float64) models.TempInfo
}

type cliApp struct {
    l      Logger
    wi     WeatherInfo
    config config.Config
}

func New(l Logger, wi WeatherInfo, cfg config.Config) *cliApp {
    return &cliApp{
        l:      l,
        wi:     wi,
        config: cfg,
    }
}

func (c *cliApp) Run() error {
    // Используем координаты из конфига
    lat := c.config.L.Lat
    lon := c.config.L.Long
    
    c.l.Debug(fmt.Sprintf("Using coordinates: lat=%.4f, lon=%.4f", lat, lon))
    
    fmt.Printf(
        "Температура воздуха - %.2f градусов цельсия\n",
        c.wi.GetTemperature(lat, lon).Temp,
    )
    return nil
}
