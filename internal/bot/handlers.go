package bot

import (
	"errors"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
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
		return b.handleAdminCommand(message)
	case "sendall":
		return b.handleAdminSendAll(message)
	case "admadd":
		return b.handleAdminAdd(message)
	case "admdelete":
		return b.handleAdminDelete(message)
	case "checkallrequests":
		return b.handleAdminAllRequests(message)
	default:
		return b.handleUnknownCommand(message)
	}
}

func (b *Bot) handleStartCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Привет, присылай ip, я расскажу где ты живешь"+
		" \n Доступные команды: /all - посмотреть все свои запросы, /admin - админ панель.")
	_, err := b.bot.Send(msg)
	if err != nil {
		b.sendErrorMessage(message.Chat.ID)
		return err
	}
	return nil
}

func (b *Bot) handleUnknownCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Я не знаю чего ты хочешь от меня.")
	_, err := b.bot.Send(msg)
	if err != nil {
		b.sendErrorMessage(message.Chat.ID)
	}
	return err
}

func (b *Bot) handleIp(message *tgbotapi.Message) error {
	ipInfo := rest.GetIpInfo(message.Text)
	msg := getMessage(message.Chat.ID, ipInfo)
	_, err := b.bot.Send(msg)
	if err != nil {
		b.sendErrorMessage(message.Chat.ID)
		return err
	}
	err = b.service.ProcessIp(message.Chat.ID, message.Chat.FirstName, ipInfo)
	if err != nil {
		return errors.New("b.service.ProcessIp: " + err.Error())
	}
	return nil

}

func getMessage(chatID int64, ip structs.IPInfo) tgbotapi.MessageConfig {
	rawString := fmt.Sprintf("your continent: %s, your country: %s, your region: %s, your city: %s, your zip: %s,"+
		" your latitude: %f, your longitude: %f",
		ip.Continent, ip.Country, ip.Region, ip.City, ip.Zip, ip.Latitude, ip.Longitude)
	return tgbotapi.NewMessage(chatID, rawString)
}

func (b *Bot) handleAllUsersHistory(message *tgbotapi.Message) error {
	user, err := b.service.GetUser(message.Chat.ID)
	if err != nil {
		b.sendErrorMessage(message.Chat.ID)
		return errors.New("b.service.GetUser: " + err.Error())
	}

	result, err := b.service.GetAllIps(user.Id)
	if err != nil {
		b.sendErrorMessage(message.Chat.ID)
		return errors.New("b.service.GetAllIps: " + err.Error())
	}

	rawMsg, err := b.getAllUsersHistoryMessage(result)
	if err != nil {
		b.sendErrorMessage(message.Chat.ID)
		return errors.New("b.getAllUsersHistoryMessage: " + err.Error())
	}
	msg := tgbotapi.NewMessage(message.Chat.ID, rawMsg)
	_, err = b.bot.Send(msg)
	if err != nil {
		b.sendErrorMessage(message.Chat.ID)
		return err
	}

	return nil
}

func (b *Bot) getAllUsersHistoryMessage(result structs.UsersRequests) (string, error) {
	responseStr := fmt.Sprint("Запросы:\n")
	for ipStr, dates := range result.Ips {
		responseStr += fmt.Sprintf("Ip : %s ", ipStr)
		ip, err := b.service.GetIpInfo(ipStr)
		if err != nil {
			return "", errors.New("b.service.GetIpInfo: " + err.Error())
		}
		responseStr += fmt.Sprintf("Результат: \n continent: %s, country: %s, region: %s, city: %s \n",
			ip.Continent, ip.Country, ip.Region, ip.City)
		responseStr += fmt.Sprintf("Время запроса:\n")

		for ix, date := range dates {
			responseStr += fmt.Sprintf("%d - %s\n", ix+1, date)
		}
	}
	return responseStr, nil
}

func (b *Bot) sendErrorMessage(chatID int64) {
	tgbotapi.NewMessage(chatID, "Ошибка сервиса. Попробуйте еще раз.")
}
