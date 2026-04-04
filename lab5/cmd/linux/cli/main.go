package main

import (
    "os"

    pogodaby "github.com/Pavel90544/VSRPP/lab5/internal/adapters/pogoda_by"
    "github.com/Pavel90544/VSRPP/lab5/internal/adapters/weather"
    "github.com/Pavel90544/VSRPP/lab5/internal/pkg/app/cli"
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
    wi := getProvider(cfg, l)
    app := cli.New(l, wi, cfg)

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
    case "pogoda":
        wi = pogodaby.New(l)
    default:
        wi = weather.New(l)
    }

    return wi
}
