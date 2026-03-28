package flags

import "flag"

type Flags struct {
    Path string // путь к конфигурационному файлу
}

// Parse - парсит аргументы командной строки
func Parse() *Flags {
    // Определяем флаг -config с значением по умолчанию
    configPath := flag.String("config", "./config/config.yaml", "path to config file")
    
    // Парсим аргументы
    flag.Parse()
    
    // Возвращаем структуру с путем к конфигу
    return &Flags{
        Path: *configPath, // разименовываем указатель
    }
}
