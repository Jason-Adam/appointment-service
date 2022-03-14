package config

import "os"

type Config struct {
	DBConfig     DatabaseConfig
	ServerConfig ServerConfig
}

type DatabaseConfig struct {
	URL string
}

type ServerConfig struct {
	Port string
}

func Local() Config {
	return Config{
		DBConfig: DatabaseConfig{
			URL: os.Getenv("DATABASE_URL"),
		},
		ServerConfig: ServerConfig{
			Port: "8080",
		},
	}
}
