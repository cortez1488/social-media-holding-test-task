package bot

import (
	"errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strings"
)

func (b *Bot) handleAdminCommand(message *tgbotapi.Message) error {
	admin, err := b.checkAdmin(message)
	if err != nil {
		return errors.New("b.checkAdmin: " + err.Error())
	}
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
	return nil
}

func (b *Bot) handleAdminSendAll(message *tgbotapi.Message) error {
	admin, err := b.checkAdmin(message)
	if err != nil {
		return errors.New("b.checkAdmin: " + err.Error())
	}
	if admin {
		admRawMessage, err := b.getAdminMessageArgument(message)
		if err != nil {
			if strings.Contains(err.Error(), "no argument") {
				return nil
			}
			return err
		}
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
	return nil
}

func (b *Bot) handleAdminAdd(message *tgbotapi.Message) error {
	admin, err := b.checkAdmin(message)
	if err != nil {
		return errors.New("b.checkAdmin: " + err.Error())
	}
	if admin {
		nickname, err := b.getAdminOneArgument(message)
		if err != nil {
			if strings.Contains(err.Error(), "no argument") {
				return nil
			}
			return err
		}
		_, err = b.admService.AddAdmin(nickname)
		if err != nil {
			return errors.New("b.admService.AddAdmin: " + err.Error())
		}
		msg := tgbotapi.NewMessage(message.Chat.ID, "Новый админ добавлен.")
		_, err = b.bot.Send(msg)
		if err != nil {
			return err
		}
	}
	return nil
}

func (b *Bot) handleAdminDelete(message *tgbotapi.Message) error {
	admin, err := b.checkAdmin(message)
	if err != nil {
		return errors.New("b.checkAdmin: " + err.Error())
	}
	if admin {
		nickname, err := b.getAdminOneArgument(message)
		if err != nil {
			if strings.Contains(err.Error(), "no argument") {
				return nil
			}
			return err
		}
		_, err = b.admService.DeleteAdmin(nickname)
		if err != nil {
			return errors.New("b.admService.AddAdmin: " + err.Error())
		}
		msg := tgbotapi.NewMessage(message.Chat.ID, "Админ удален.")
		_, err = b.bot.Send(msg)
		if err != nil {
			return err
		}
	}
	return nil
}

func (b *Bot) handleAdminAllRequests(message *tgbotapi.Message) error {
	admin, err := b.checkAdmin(message)
	if err != nil {
		return err
	}
	if admin {
		nickname, err := b.getAdminOneArgument(message)
		if err != nil {
			if strings.Contains(err.Error(), "no argument") {
				return nil
			}
			return err
		}
		id, err := b.admService.GetIdFromNickname(nickname)
		if err != nil {
			return errors.New("b.admService.GetIdFromNickname: " + err.Error())
		}
		data, err := b.service.GetAllIps(id)
		if err != nil {
			return errors.New("b.service.GetAllIps: " + err.Error())
		}
		msg := tgbotapi.NewMessage(message.Chat.ID, b.getAllUsersHistoryMessage(data))
		_, err = b.bot.Send(msg)
		if err != nil {
			return err
		}
	}
	return nil
}

func (b *Bot) checkAdmin(message *tgbotapi.Message) (bool, error) {
	admin, err := b.admService.CheckAdminRight(message.Chat.ID)
	if err != nil {
		return false, errors.New("b.admService.CheckAdminRight: " + err.Error())
	}
	if !admin {
		msg := tgbotapi.NewMessage(message.Chat.ID, "У вас нет прав!")
		_, err = b.bot.Send(msg)
		if err != nil {
			return false, err
		}
		return false, nil
	}
	return true, nil
}

func (b *Bot) getAdminOneArgument(message *tgbotapi.Message) (string, error) {
	cmndAndArgs := strings.Split(message.Text, " ")
	if len(cmndAndArgs) < 2 {
		msg := tgbotapi.NewMessage(message.Chat.ID, "Вы не ввели аргумент")
		_, err := b.bot.Send(msg)
		if err != nil {
			return "", err
		}
		return "", errors.New("no argument")
	}
	return cmndAndArgs[1], nil
}

func (b *Bot) getAdminMessageArgument(message *tgbotapi.Message) (string, error) {
	cmndAndArgs := strings.Split(message.Text, " ")
	if len(cmndAndArgs) < 2 {
		msg := tgbotapi.NewMessage(message.Chat.ID, "Вы не ввели аргумент")
		_, err := b.bot.Send(msg)
		if err != nil {
			return "", err
		}
		return "", errors.New("no argument")
	}
	return strings.Replace(message.Text, "/sendall", "", 1), nil

}
