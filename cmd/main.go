package main

import (
	"log"
	"log/slog"

	"github.com/joho/godotenv"
	"github.com/wiliamvj/go-vagas/internal/bot"
)

func main() {
	slog.Info("Starting bot")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	err = bot.Websocket()
	if err != nil {
		log.Fatal(err)
	}
}
