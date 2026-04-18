package gui

import (
    "github.com/Pavel90544/VSRPP/lab5/internal/domain/gui_settings"
    "github.com/Pavel90544/VSRPP/lab5/internal/pkg/app/cli"
    "github.com/Pavel90544/VSRPP/lab5/pkg/config"
)

type guiApp struct {
    l        cli.Logger
    provider guisettings.WindowProvider
    wi       cli.WeatherInfo
    cfg      config.Config
}

func New(l cli.Logger, provider guisettings.WindowProvider, wi cli.WeatherInfo, cfg config.Config) *guiApp {
    return &guiApp{
        l:        l,
        provider: provider,
        wi:       wi,
        cfg:      cfg,
    }
}

func (g *guiApp) Run() error {
    // Создаем окно
    windowSize := guisettings.NewWS(400, 300)
    win, err := g.provider.CreateWindow("Weather Informer", windowSize)
    if err != nil {
        return err
    }

    // Создаем виджет текста
    textWidget := g.provider.GetTextWidget("Загрузка...")
    
    // Устанавливаем виджет в окно
    if err := win.SetTemperatureWidget(textWidget); err != nil {
        return err
    }

    // Получаем температуру
    tempInfo, err := g.wi.GetTemperature(g.cfg.L.Lat, g.cfg.L.Long)
    if err != nil {
        g.l.Error("Ошибка получения погоды", err)
        textWidget.SetText("Ошибка получения данных")
    } else {
        win.UpdateTemperature(tempInfo.Temp)
    }

    // Отображаем окно
    win.Render()
    
    // Запускаем приложение
    appRunner := g.provider.GetAppRunner()
    appRunner.Run()
    
    return nil
}
