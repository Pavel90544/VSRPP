package cli

import (
    "encoding/json"
    "errors"
    "fmt"
    "io"
    "net/http"

    "github.com/Pavel90544/VSRPP/lab5/internal/pkg/logger"
)

type cliApp struct {
    logger logger.Logger
    lat    float64
    lon    float64
}

func New(logger logger.Logger, lat, lon float64) *cliApp {
    return &cliApp{
        logger: logger,
        lat:    lat,
        lon:    lon,
    }
}

func (c *cliApp) Run() error {
    c.logger.Info("Starting weather informer")

    type Current struct {
        Temp float32 `json:"temperature_2m"`
    }
    type Response struct {
        Curr Current `json:"current"`
    }
    var response Response

    params := fmt.Sprintf(
        "latitude=%f&longitude=%f&current=temperature_2m",
        c.lat,
        c.lon,
    )
    url := fmt.Sprintf("https://api.open-meteo.com/v1/forecast?%s", params)
    
    c.logger.Debug("Requesting URL: %s", url)

    resp, err := http.Get(url)
    if err != nil {
        c.logger.Error("HTTP request failed: %v", err)
        customErr := errors.New("can't get weather data from openmeteo")
        return errors.Join(customErr, err)
    }
    defer func() {
        if err := resp.Body.Close(); err != nil {
            c.logger.Error("can't close body: %v", err)
        }
    }()

    if resp.StatusCode != http.StatusOK {
        c.logger.Error("API returned non-200 status: %d", resp.StatusCode)
        return fmt.Errorf("API returned status: %d", resp.StatusCode)
    }

    data, err := io.ReadAll(resp.Body)
    if err != nil {
        c.logger.Error("Failed to read response body: %v", err)
        customErr := errors.New("can't read data from response")
        return errors.Join(customErr, err)
    }

    c.logger.Debug("Response data: %s", string(data))

    if err := json.Unmarshal(data, &response); err != nil {
        c.logger.Error("Failed to unmarshal JSON: %v", err)
        customErr := errors.New("can't unmarshal data from response")
        return errors.Join(customErr, err)
    }

    result := fmt.Sprintf("Температура воздуха - %.2f градусов цельсия", response.Curr.Temp)
    fmt.Println(result)
    c.logger.Info("Successfully retrieved weather data: %.2f°C", response.Curr.Temp)
    
    return nil
}
