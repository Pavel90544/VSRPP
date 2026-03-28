package main

import (
    "os"

    "github.com/Pavel90544/VSRPP/lab5/internal/adapters/weather"
    "github.com/Pavel90544/VSRPP/lab5/internal/pkg/app/cli"
    "github.com/Pavel90544/VSRPP/lab5/internal/pkg/cache"
    "github.com/Pavel90544/VSRPP/lab5/internal/pkg/flags"
    "github.com/Pavel90544/VSRPP/lab5/pkg/config"
    "github.com/Pavel90544/VSRPP/lab5/pkg/logger"
)

func main() {
    arguments := flags.Parse()

    r, err := os.Open(arguments.Path)
    if err != nil {
        panic(err)
    }
    defer r.Close()

    cfg, err := config.Parse(r)
    if err != nil {
        panic(err)
    }

    l := logger.New()

    // Создаем кэш в зависимости от типа
    var cacheInstance cache.Cache
    
    switch cfg.C.Type {
    case "file":
        l.Info("Using file cache")
        cacheInstance, err = cache.NewFileCache(cfg.C.Dir)
        if err != nil {
            l.Error("Failed to create file cache", err)
            panic(err)
        }
    default:
        l.Info("Using memory cache")
        cacheInstance = cache.NewMemoryCache()
    }

    // Получаем базового провайдера
    baseWeather := getProvider(cfg, l)
    
    // Оборачиваем в кэширующий адаптер
    cachedWeather := weather.NewCachedWeatherInfo(baseWeather, cacheInstance, cfg.GetTTL(), l)

    app := cli.New(l, cachedWeather, cfg)

    err = app.Run()
    if err != nil {
        l.Error("Some error", err)
        os.Exit(1)
    }

    os.Exit(0)
}

func getProvider(cfg config.Config, l cli.Logger) cli.WeatherInfo {
    var wi cli.WeatherInfo

    switch cfg.P.Type {
    case "open-meteo":
        wi = weather.New(l)
    default:
        wi = weather.New(l)
    }

    return wi
}
