package main

import (
	"os"

	"github.com/Pavel90544/VSRPP/lab5/internal/adapters/weather"
	"github.com/Pavel90544/VSRPP/lab5/internal/pkg/app/cli"
	"github.com/Pavel90544/VSRPP/lab5/internal/pkg/flags"
	"github.com/Pavel90544/VSRPP/lab5/pkg/config"
	"github.com/Pavel90544/VSRPP/lab5/pkg/logger"
)

func main() {
	// Шаг 1: Парсим аргументы командной строки
	arguments := flags.Parse()

	// Шаг 2: Открываем конфигурационный файл
	r, err := os.Open(arguments.Path)
	if err != nil {
		panic(err) // Если файл не найден - паникуем
	}
	defer r.Close()

	// Шаг 3: Парсим конфигурационный файл
	cfg, err := config.Parse(r)
	if err != nil {
		panic(err)
	}

	// Шаг 4: Создаем логгер
	l := logger.New()

	// Шаг 5: Получаем провайдера погоды
	wi := getProvider(cfg, l)

	// Шаг 6: Создаем приложение с зависимостями
	app := cli.New(l, wi, cfg)

	// Шаг 7: Запускаем приложение
	err = app.Run()
	if err != nil {
		l.Error("Some error", err)
		os.Exit(1)
	}

	os.Exit(0)
}

// getProvider - фабричная функция для выбора источника данных о погоде
func getProvider(cfg config.Config, l cli.Logger) cli.WeatherInfo {
	var wi cli.WeatherInfo

	switch cfg.P.Type {
	case "open-meteo":
		wi = weather.New(l)
	default:
		// По умолчанию используем open-meteo
		wi = weather.New(l)
	}

	return wi
}
