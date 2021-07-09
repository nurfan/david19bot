package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/nurfan/david19bot/app/handler"
)

func init() {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func main() {
	http.HandleFunc("/ping", handler.Ping)
	http.HandleFunc("/telegram/webhook", handler.HandleTelegramWebHook)
	http.ListenAndServe(":"+os.Getenv("APP_PORT"), nil)
}
