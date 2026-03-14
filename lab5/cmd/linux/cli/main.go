package main

import (
    "flag"
    "fmt"
    "os"

    "github.com/Pavel90544/VSRPP/lab5/internal/pkg/app/cli"
    "github.com/Pavel90544/VSRPP/lab5/internal/pkg/logger"
)

func main() {
    // Парсим флаги командной строки
    debug := flag.Bool("debug", false, "enable debug mode")
    lat := flag.Float64("lat", 53.6688, "latitude coordinate")
    lon := flag.Float64("lon", 23.8223, "longitude coordinate")
    flag.Parse()

    // Создаем логгер
    log := logger.New(*debug)
    
    log.Info("Application started")
    log.Debug("Debug mode is enabled")
    log.Debug("Coordinates: lat=%.4f, lon=%.4f", *lat, *lon)

    // Создаем приложение с внедрением зависимости логгера
    app := cli.New(log, *lat, *lon)
    
    err := app.Run()
    if err != nil {
        log.Error("Application error: %s", err.Error())
        fmt.Printf("Some error - %s\n", err.Error())
        os.Exit(1)
    }
    
    log.Info("Application finished successfully")
    os.Exit(0)
}
