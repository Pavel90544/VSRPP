package main

import (
    "os"

    "github.com/Pavel90544/VSRPP/lab5/internal/pkg/app/gui"
    "github.com/Pavel90544/VSRPP/lab5/internal/pkg/flags"
    "github.com/Pavel90544/VSRPP/lab5/internal/pkg/gui/fyne"
    "github.com/Pavel90544/VSRPP/lab5/internal/pkg/providers"
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
    provider := providers.GetProvider(cfg, l)
    
    p := fyne.NewP()
    g := gui.New(l, p, provider, cfg)
    
    err = g.Run()
    if err != nil {
        panic(err)
    }
}
