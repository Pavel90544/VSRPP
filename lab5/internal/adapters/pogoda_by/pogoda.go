package pogodaby

import (
    "encoding/json"
    "net/http"

    "github.com/Pavel90544/VSRPP/lab5/internal/domain/models"
)

const apiURL = "https://pogoda.by/api/v2/weather-fact?station=26820"

type resp struct {
    Temp float32 `json:"t"`
}

type Logger interface {
    Info(string)
    Debug(string)
    Error(string, error)
}

type pogoda struct {
    l Logger
}

func New(l Logger) *pogoda {
    return &pogoda{l: l}
}

func (p *pogoda) GetTemperature(lat, long float64) (models.TempInfo, error) {
    p.l.Debug("Getting weather from pogoda.by")

    response, err := http.Get(apiURL)
    if err != nil {
        p.l.Error("can't get data from pogoda.by", err)
        return models.TempInfo{}, err
    }
    defer func() {
        err := response.Body.Close()
        if err != nil {
            p.l.Error("can't close response body", err)
        }
    }()

    var r resp
    if err := json.NewDecoder(response.Body).Decode(&r); err != nil {
        p.l.Error("can't decode JSON", err)
        return models.TempInfo{}, err
    }

    p.l.Debug("Successfully got temperature from pogoda.by")

    return models.TempInfo{Temp: r.Temp}, nil
}
