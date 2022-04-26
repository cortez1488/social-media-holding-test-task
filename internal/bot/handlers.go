package bot

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"social-media-holding-test-task/internal/handler/rest"
	"social-media-holding-test-task/structs"
)

func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	switch message.Command() {
	case "start":
		return b.handleStartCommand(message)
	case "all":
		return b.handleAllUsersHistory(message)
	case "admin":
		b.handleAdminCommand(message)
	case "sendall":
		b.handleAdminSendAll(message)
	//case "admadd":
	//	b.handleAdminAdd(message)
	//case "admdelete":
	//	b.handleAdminDelete(message)
	//case "checkallrequests":
	//	b.handleAdminAllRequests(message)
	default:
		return b.handleUnknownCommand(message)
	}
	return nil
}

func (b *Bot) handleStartCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Привет, присылай ip, я расскажу где ты живешь"+
		" \n Доступные команды: /all - посмотреть все свои запросы, /admin - админ панель.")
	_, err := b.bot.Send(msg)
	if err != nil {
		return err
	}
	return nil
}

func (b *Bot) handleUnknownCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "я не знаю ч ты хочешь от меня")
	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) handleIp(message *tgbotapi.Message) error {
	ipInfo := rest.GetIpInfo(message.Text)
	msg := getMessage(message.Chat.ID, ipInfo)
	_, err := b.bot.Send(msg)
	if err != nil {
		return err
	}
	return b.service.ProcessIp(message.Chat.ID, message.Chat.FirstName, ipInfo)

}

func getMessage(chatID int64, ip structs.IPInfo) tgbotapi.MessageConfig {
	rawString := fmt.Sprintf("your continent: %s, your country: %s, your region: %s, your city: %s, your zip: %s, your latitude: %s, your langirude: %s",
		ip.Continent, ip.Country, ip.Region, ip.City, ip.Zip, ip.Latitude, ip.Longitude)
	return tgbotapi.NewMessage(chatID, rawString)
}

func (b *Bot) handleAllUsersHistory(message *tgbotapi.Message) error {
	id, err := b.service.GetUser(message.Chat.ID)
	if err != nil {
		return err
	}

	result, err := b.service.GetAllIps(id)
	if err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, b.getAllUsersHistoryMessage(result))
	_, err = b.bot.Send(msg)
	if err != nil {
		return err
	}

	return nil
}

func (b *Bot) getAllUsersHistoryMessage(result structs.UsersRequests) string {
	responseStr := fmt.Sprint("")
	for ipStr, dates := range result.Ips {
		responseStr += fmt.Sprintf("Ip : %s ", ipStr)
		ip, err := b.service.GetIpInfo(ipStr)
		if err != nil {
			log.Fatal(err)
		}
		responseStr += fmt.Sprintf("Result: \n continent: %s, country: %s, region: %s, city: %s \n",
			ip.Continent, ip.Country, ip.Region, ip.City)
		responseStr += fmt.Sprintf("Request times:\n")

		for ix, date := range dates {
			responseStr += fmt.Sprintf("%d - %s\n", ix+1, date)
		}
	}
	return responseStr
}
