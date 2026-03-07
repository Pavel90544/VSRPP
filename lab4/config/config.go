package config

type Config struct {
    Port   string
    DBPath string
    Env    string
}

func New() *Config {
    return &Config{
        Port:   ":8080",
        DBPath: "products.db",
        Env:    "development",
    }
}
