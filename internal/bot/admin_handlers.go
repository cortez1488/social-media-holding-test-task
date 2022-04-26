package bot

import (
	"errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strings"
)

func (b *Bot) handleAdminCommand(message *tgbotapi.Message) {
	admin := b.checkAdmin(message)
	if admin {
		msg := tgbotapi.NewMessage(message.Chat.ID, "Вы можете: \n"+
			"-Отправить всем пользователям сообщение /sendall [сообщение] \n "+
			"-Добавиить админа /admadd [никнейм юзера] \n "+
			"-Удалить админа /admdelete [никнейм юзера] \n "+
			"-Вывести все айпи что проверял пользователь /checkallrequests [никнейм юзера]")
		_, err := b.bot.Send(msg)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (b *Bot) handleAdminSendAll(message *tgbotapi.Message) {
	admin := b.checkAdmin(message)
	if admin {
		admRawMessage := b.getAdminMessageArgument(message)
		chats, err := b.admService.GetAllUsersChatID()
		if err != nil {
			log.Fatal(errors.New("b.admService.GetAllUsersChatID: " + err.Error()))
		}
		for _, chat := range chats {
			msg := tgbotapi.NewMessage(chat, admRawMessage)
			_, err := b.bot.Send(msg)
			if err != nil {
				log.Fatal(errors.New("b.admService.GetAllUsersChatID: " + err.Error()))
			}
		}
	}

}

//func (b *Bot) handleAdminAdd(message *tgbotapi.Message) {
//	admin := b.checkAdmin(message)
//
//}
//
//func (b *Bot) handleAdminDelete(message *tgbotapi.Message) {
//	admin := b.checkAdmin(message)
//
//}
//
//func (b *Bot) handleAdminAllRequests(message *tgbotapi.Message) {
//	admin := b.checkAdmin(message)
//
//}

func (b *Bot) checkAdmin(message *tgbotapi.Message) bool {
	admin, err := b.admService.CheckAdminRight(message.Chat.ID)
	if err != nil {
		log.Fatal(errors.New("b.admService.CheckAdminRight: " + err.Error()))
	}
	if !admin {
		msg := tgbotapi.NewMessage(message.Chat.ID, "У вас нет прав!")
		_, err = b.bot.Send(msg)
		if err != nil {
			log.Fatal(err)
		}
		return false
	}
	return true
}

func (b *Bot) getAdminOneArgument(message *tgbotapi.Message) string {
	cmndAndArgs := strings.Split(message.Text, " ")
	if len(cmndAndArgs) < 2 {
		msg := tgbotapi.NewMessage(message.Chat.ID, "Вы не ввели аргумент")
		_, err := b.bot.Send(msg)
		if err != nil {
			log.Fatal(err)
		}
	}
	return cmndAndArgs[1]
}

func (b *Bot) getAdminMessageArgument(message *tgbotapi.Message) string {
	cmndAndArgs := strings.Split(message.Text, " ")
	if len(cmndAndArgs) < 2 {
		msg := tgbotapi.NewMessage(message.Chat.ID, "Вы не ввели аргумент")
		_, err := b.bot.Send(msg)
		if err != nil {
			log.Fatal(err)
		}
	}
	return strings.Replace(message.Text, "/sendall", "", 1)

}
