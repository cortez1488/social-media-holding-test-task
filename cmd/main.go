package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"social-media-holding-test-task/internal/bot"
	"social-media-holding-test-task/internal/core/admin"
	"social-media-holding-test-task/internal/core/ip"
	db2 "social-media-holding-test-task/internal/db"
)

func main() {
	db, err := sqlx.Connect("postgres", getPostgresDBConnectString())
	if err != nil {
		log.Fatal(err)
	}

	botAPI, err := tgbotapi.NewBotAPI("5187131287:AAH7x1R1GzEIpOK_RCgz9xieOqjzIRVmhug")
	if err != nil {
		log.Fatal(err)
	}

	ipRepo := db2.NewIpStorage(db)
	ipService := ip.NewIpService(ipRepo)

	admStorage := db2.NewAdminStorage(db)
	admService := admin.NewAdminService(admStorage)

	telegramBot := bot.NewBot(botAPI, ipService, admService)
	telegramBot.Start()
}

func getPostgresDBConnectString() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", "localhost", "5436",
		"postgres", "qwerty", "postgres", "disable")

}
