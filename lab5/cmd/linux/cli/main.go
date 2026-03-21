package main

import (
    "os"

    "github.com/Pavel90544/VSRPP/lab5/internal/adapters/weather"
    "github.com/Pavel90544/VSRPP/lab5/internal/pkg/app/cli"
    "github.com/Pavel90544/VSRPP/lab5/pkg/logger"
)

func main() {
    l := logger.New()
    wi := weather.New(l)
    app := cli.New(l, wi)
    err := app.Run()
    if err != nil {
        l.Error("Some error", err)
        os.Exit(1)
    }
    os.Exit(0)
}
