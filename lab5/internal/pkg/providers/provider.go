package providers

import (
    pogodaby "github.com/Pavel90544/VSRPP/lab5/internal/adapters/pogoda_by"
    "github.com/Pavel90544/VSRPP/lab5/internal/adapters/weather"
    "github.com/Pavel90544/VSRPP/lab5/internal/pkg/app/cli"
    "github.com/Pavel90544/VSRPP/lab5/pkg/config"
)

func GetProvider(c config.Config, l cli.Logger) cli.WeatherInfo {
    var wi cli.WeatherInfo
    switch c.P.Type {
    case "open-meteo":
        wi = weather.New(l)
    case "pogoda":
        wi = pogodaby.New(l)
    default:
        wi = weather.New(l)
    }
    return wi
}
