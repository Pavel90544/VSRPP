package config

import (
    "io"
    "time"
    "gopkg.in/yaml.v3"
)

type ConfigFile struct {
    C Config `yaml:"service"`
}

type Provider struct {
    Type string `yaml:"type"`
}

type Location struct {
    Lat  float64 `yaml:"lat"`
    Long float64 `yaml:"long"`
}

type Cache struct {
    Type string `yaml:"type"`
    TTL  int    `yaml:"ttl"`
    Dir  string `yaml:"dir"`
}

type Config struct {
    P Provider `yaml:"provider"`
    L Location `yaml:"location"`
    C Cache    `yaml:"cache"`
}

func Parse(r io.Reader) (Config, error) {
    var c ConfigFile
    if err := yaml.NewDecoder(r).Decode(&c); err != nil {
        return Config{}, err
    }
    
    if c.C.C.TTL == 0 {
        c.C.C.TTL = 300
    }
    if c.C.C.Type == "" {
        c.C.C.Type = "memory"
    }
    if c.C.C.Dir == "" {
        c.C.C.Dir = "./cache"
    }
    
    return c.C, nil
}

func (c *Config) GetTTL() time.Duration {
    return time.Duration(c.C.TTL) * time.Second
}
