package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"os"
	"social-media-holding-test-task/internal/bot"
	"social-media-holding-test-task/internal/core/admin"
	"social-media-holding-test-task/internal/core/ip"
	"social-media-holding-test-task/internal/core/user_info"
	db2 "social-media-holding-test-task/internal/db"
	"social-media-holding-test-task/internal/handler/rest"
	"strings"
	"time"
)

func main() {

	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("config read error" + err.Error())
	}

	db := initPostgresDB()
	migrateDB(db)

	botAPI, err := tgbotapi.NewBotAPI(viper.GetString("tg_token"))
	if err != nil {
		log.Fatal("tg bot initialization " + err.Error())
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

func initPostgresDB() *sqlx.DB {
	log.Println("sql connect string:", getPostgresDBConnectString())
	var db *sqlx.DB
	var err error
	var errConnectionRefusedCounter int

	for {
		db, err = sqlx.Connect(viper.GetString("db.postgres.drivername"), getPostgresDBConnectString())

		if err != nil {
			if strings.Contains(err.Error(), "connect: connection refused") {
				errConnectionRefusedCounter++
				time.Sleep(time.Millisecond * 1000)
				log.Println(errConnectionRefusedCounter+1, "attempt to connect to database")
			}

			if errConnectionRefusedCounter >= 5 {
				log.Fatal("PostgresDB initialization " + err.Error())
			}
		} else {
			break
		}
	}
	return db
}

func getPostgresDBConnectString() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("DB_HOST"), viper.Get("db.postgres.port"),
		viper.Get("db.postgres.username"), viper.Get("db.postgres.password"),
		viper.Get("db.postgres.dbname"), viper.Get("db.postgres.sslmode"))
}

//func getPostgresDBConnectString() string {
//	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", "localhost", "5436",
//		"postgres", "qwerty", "postgres", "disable")
//
//}

func migrateDB(db *sqlx.DB) {
	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	files, err := ioutil.ReadDir("./schema")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		fmt.Println(f.Name())
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file:./schema/",
		viper.GetString("dbname"), driver)

	if err != nil {
		log.Fatalln("Error with database migration creating:", err)
	}

	err = m.Up()
	if err != nil {
		if !strings.Contains(err.Error(), "no change") {
			log.Fatalln("Error with database migration:", err)
		}
	}
}
