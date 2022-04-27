package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
	"social-media-holding-test-task/internal/bot"
	"social-media-holding-test-task/internal/core/admin"
	"social-media-holding-test-task/internal/core/ip"
	"social-media-holding-test-task/internal/core/user_info"
	db2 "social-media-holding-test-task/internal/db"
	"social-media-holding-test-task/internal/handler/rest"
)

func main() {

	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	db, err := sqlx.Connect("postgres", getPostgresDBConnectString())
	if err != nil {
		log.Fatal(err)
	}

	botAPI, err := tgbotapi.NewBotAPI(viper.Get("tg_token").(string))
	if err != nil {
		log.Fatal(err)
	}

	ipRepo := db2.NewIpStorage(db)
	ipService := ip.NewIpService(ipRepo)

	admStorage := db2.NewAdminStorage(db)
	admService := admin.NewAdminService(admStorage)

	userInfoRepo := db2.NewUserInfoStorage(db)
	userInfoService := user_info.NewUserInfoService(userInfoRepo, ipRepo)

	handler := rest.NewHandler(userInfoService)
	server := handler.InitRoutes()

	telegramBot := bot.NewBot(botAPI, ipService, admService)

	go runServer(server)
	telegramBot.Start()
}

func runServer(server *gin.Engine) {
	err := server.Run()
	if err != nil {
		log.Fatal("server: " + err.Error())
	}
}

//func getPostgresDBConnectString() string {
//	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
//		os.Getenv("DB_HOST"), viper.Get("db.postgres.port"),
//		viper.Get("db.postgres.username"), viper.Get("db.postgres.password"),
//		viper.Get("db.postgres.dbname"), viper.Get("db.postgres.sslmode"))
//}

func getPostgresDBConnectString() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", "localhost", "5436",
		"postgres", "qwerty", "postgres", "disable")

}
