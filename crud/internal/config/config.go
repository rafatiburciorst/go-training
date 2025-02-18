package config

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort string
}

func Load() *Config {
	err := godotenv.Load()
	if err != nil {
		slog.Error("Erro ao carregar o arquivo .env")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	return &Config{
		ServerPort: "localhost:" + port,
	}
}
