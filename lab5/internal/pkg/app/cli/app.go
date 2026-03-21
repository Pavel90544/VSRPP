package cli

import (
    "fmt"

    "github.com/Pavel90544/VSRPP/lab5/internal/domain/models"
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
    l  Logger
    wi WeatherInfo
}

func New(l Logger, wi WeatherInfo) *cliApp {
    return &cliApp{
        l:  l,
        wi: wi,
    }
}

func (c *cliApp) Run() error {
    fmt.Printf(
        "Температура воздуха - %.2f градусов цельсия\n",
        c.wi.GetTemperature(53.6688, 23.8223).Temp,
    )
    return nil
}
