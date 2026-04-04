package flags

import (
    "flag"
    "os"
)

type Flags struct {
    Path string
}

// Parse - парсит аргументы командной строки
func Parse() *Flags {
    // Создаем новый набор флагов для каждого вызова
    flagSet := flag.NewFlagSet("weather-app", flag.ContinueOnError)
    
    configPath := flagSet.String("config", "./config/config.yaml", "path to config file")
    
    // Парсим аргументы
    flagSet.Parse(os.Args[1:])
    
    return &Flags{
        Path: *configPath,
    }
}
