package main

import (
	"log"
	"log/slog"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/wiliamvj/go-vagas/internal/bot"
)

func main() {
	slog.Info("Starting bot")
	_ = godotenv.Load()

	go func() {
		http.HandleFunc("/health", bot.HealthCheck)
		slog.Info("Starting health check server on :8080")

		if err := http.ListenAndServe(":8080", nil); err != nil {
			log.Fatal("Failed to start health check server:", err)
		}
	}()

	err := bot.Websocket()
	if err != nil {
		log.Fatal(err)
	}
}
