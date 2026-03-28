package config

import (
    "io"
    "gopkg.in/yaml.v3"
)

// ConfigFile - структура корневого элемента yaml файла
type ConfigFile struct {
    C Config `yaml:"service"`
}

// Provider - структура для хранения типа провайдера
type Provider struct {
    Type string `yaml:"type"`
}

// Location - структура для хранения координат
type Location struct {
    Lat  float64 `yaml:"lat"`
    Long float64 `yaml:"long"`
}

// Config - основная структура конфигурации
type Config struct {
    P Provider `yaml:"provider"`
    L Location `yaml:"location"`
}

// Parse - функция парсинга yaml из io.Reader
func Parse(r io.Reader) (Config, error) {
    var c ConfigFile
    if err := yaml.NewDecoder(r).Decode(&c); err != nil {
        return Config{}, err
    }
    return c.C, nil
}
