package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"regexp"
)

type Bot struct {
	bot     *tgbotapi.BotAPI
	service IpService
}

func NewBot(bot *tgbotapi.BotAPI, service IpService) *Bot {
	return &Bot{bot: bot, service: service}
}

func (b *Bot) Start() {
	updates := b.initUpdates()
	err := b.handleFunc(updates)
	if err != nil {
		log.Fatal(err)
	}
}

func (b *Bot) handleFunc(updates tgbotapi.UpdatesChannel) error {
	for update := range updates {
		if update.Message != nil {
			if update.Message.IsCommand() {
				err := b.handleCommand(update.Message)
				if err != nil {
					return err
				}
			}
			if isIpAddress(update.Message.Text) {
				err := b.handleIp(update.Message)
				if err != nil {
					return err
				}
			}

		}
	}
	return nil
}

func isIpAddress(rawString string) bool {
	matched, _ := regexp.MatchString(`^((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$`, rawString)
	return matched
}

func (b *Bot) initUpdates() tgbotapi.UpdatesChannel {
	updatesCfg := tgbotapi.NewUpdate(0)
	updatesCfg.Timeout = 60
	return b.bot.GetUpdatesChan(updatesCfg)
}
