package config

import "os"

type AppConfig struct {
	Port string
	Mode string
	CORS string
}

type DBConf struct {
	Dialect  string
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
}

type Config struct {
	App *AppConfig
	DB  *DBConf
	// TODO: add other stuff later as configs for technology stack
}

func GetConfig() *Config {
	return &Config{
		App: &AppConfig{
			Port: getEnvAsStr("APP_PORT", "8080"),
			Mode: getEnvAsStr("APP_MODE", "release"),
		},
		DB: &DBConf{
			Dialect:  os.Getenv("POSTGRES_DIALECT"),
			Host:     os.Getenv("POSTGRES_URI"),
			Port:     os.Getenv("POSTGRES_PORT"),
			Username: os.Getenv("POSTGRES_USER"),
			Password: os.Getenv("POSTGRES_PASSWORD"),
			DBName:   os.Getenv("POSTGRES_DB"),
		},
	}
}
