package main

import (
	"crud/internal/config"
	"crud/internal/models"
	"crud/internal/routes"
	"log/slog"
	"net/http"
)

func main() {
	db := make(map[int64]models.User)
	slog.Info("HTTP server is running")
	cfg := config.Load()
	router := routes.Init(db)

	if err := http.ListenAndServe(cfg.ServerPort, router); err != nil {
		panic(err)
	}
}
