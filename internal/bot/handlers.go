package bot

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"social-media-holding-test-task/internal/handler/rest"
)

func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	switch message.Command() {
	case "start":
		return b.handleStartCommand(message)
	default:
		return b.handleUnknownCommand(message)
	}

}

func (b *Bot) handleStartCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "")
	_, err := b.bot.Send(msg)
	if err != nil {
		return err
	}
	return nil
}

func (b *Bot) handleUnknownCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "")
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

func getMessage(chatID int64, ip rest.IPInfo) tgbotapi.MessageConfig {
	rawString := fmt.Sprintf("your continent: %s, your country: %s, your region: %s, your city: %s, your zip: %s",
		ip.Continent, ip.Country, ip.City, ip.Zip)
	return tgbotapi.NewMessage(chatID, rawString)
}
