package bot

import (
	"errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strings"
)

func (b *Bot) handleAdminCommand(message *tgbotapi.Message) error {
	admin, err := b.checkAdmin(message)
	if err != nil {
		b.sendErrorMessage(message.Chat.ID)
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
			b.sendErrorMessage(message.Chat.ID)
			return err
		}
	}
	return nil
}

func (b *Bot) handleAdminSendAll(message *tgbotapi.Message) error {
	admin, err := b.checkAdmin(message)
	if err != nil {
		b.sendErrorMessage(message.Chat.ID)
		return errors.New("b.checkAdmin: " + err.Error())
	}
	if admin {
		admRawMessage, err := b.getAdminMessageArgument(message)
		if err != nil {
			if strings.Contains(err.Error(), "no argument") {
				return nil
			}
			b.sendErrorMessage(message.Chat.ID)
			return errors.New(" b.getAdminMessageArgument: " + err.Error())
		}
		chats, err := b.admService.GetAllUsersChatID()
		if err != nil {
			b.sendErrorMessage(message.Chat.ID)
			return errors.New("b.admService.GetAllUsersChatID: " + err.Error())
		}
		for _, chat := range chats {
			msg := tgbotapi.NewMessage(chat, admRawMessage)
			_, err := b.bot.Send(msg)
			if err != nil {
				b.sendErrorMessage(message.Chat.ID)
				return errors.New("b.admService.GetAllUsersChatID: " + err.Error())
			}
		}
	}
	return nil
}

func (b *Bot) handleAdminAdd(message *tgbotapi.Message) error {
	admin, err := b.checkAdmin(message)
	if err != nil {
		b.sendErrorMessage(message.Chat.ID)
		return errors.New("b.checkAdmin: " + err.Error())
	}
	if admin {
		nickname, err := b.getAdminOneArgument(message)
		if err != nil {
			if strings.Contains(err.Error(), "no argument") {
				return nil
			}
			b.sendErrorMessage(message.Chat.ID)
			return errors.New(" b.getAdminOneArgument: " + err.Error())
		}
		_, err = b.admService.AddAdmin(nickname)
		if err != nil {
			b.sendErrorMessage(message.Chat.ID)
			return errors.New("b.admService.AddAdmin: " + err.Error())
		}
		msg := tgbotapi.NewMessage(message.Chat.ID, "Новый админ добавлен.")
		_, err = b.bot.Send(msg)
		if err != nil {
			b.sendErrorMessage(message.Chat.ID)
			return err
		}
	}
	return nil
}

func (b *Bot) handleAdminDelete(message *tgbotapi.Message) error {
	admin, err := b.checkAdmin(message)
	if err != nil {
		b.sendErrorMessage(message.Chat.ID)
		return errors.New("b.checkAdmin: " + err.Error())
	}
	if admin {
		nickname, err := b.getAdminOneArgument(message)
		if err != nil {
			if strings.Contains(err.Error(), "no argument") {
				return nil
			}
			b.sendErrorMessage(message.Chat.ID)
			return errors.New(" b.getAdminOneArgument: " + err.Error())
		}
		_, err = b.admService.DeleteAdmin(nickname)
		if err != nil {
			b.sendErrorMessage(message.Chat.ID)
			return errors.New("b.admService.AddAdmin: " + err.Error())
		}
		msg := tgbotapi.NewMessage(message.Chat.ID, "Админ удален.")
		_, err = b.bot.Send(msg)
		if err != nil {
			b.sendErrorMessage(message.Chat.ID)
			return err
		}
	}
	return nil
}

func (b *Bot) handleAdminAllRequests(message *tgbotapi.Message) error {
	admin, err := b.checkAdmin(message)
	if err != nil {
		b.sendErrorMessage(message.Chat.ID)
		return errors.New("b.sendErrorMessage: " + err.Error())
	}
	if admin {
		nickname, err := b.getAdminOneArgument(message)
		if err != nil {
			if strings.Contains(err.Error(), "no argument") {
				return nil
			}
			b.sendErrorMessage(message.Chat.ID)
			return errors.New("b.getAllUsersHistoryMessage: " + err.Error())
		}
		id, err := b.admService.GetIdFromNickname(nickname)
		if err != nil {
			b.sendErrorMessage(message.Chat.ID)
			return errors.New("b.admService.GetIdFromNickname: " + err.Error())
		}
		data, err := b.service.GetAllIps(id)
		if err != nil {
			b.sendErrorMessage(message.Chat.ID)
			return errors.New("b.service.GetAllIps: " + err.Error())
		}
		rawMsg, err := b.getAllUsersHistoryMessage(data)
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
	}
	return nil
}

func (b *Bot) checkAdmin(message *tgbotapi.Message) (bool, error) {
	admin, err := b.admService.CheckAdminRight(message.Chat.ID)
	if err != nil {
		b.sendErrorMessage(message.Chat.ID)
		return false, errors.New("b.admService.CheckAdminRight: " + err.Error())
	}
	if !admin {
		msg := tgbotapi.NewMessage(message.Chat.ID, "У вас нет прав!")
		_, err = b.bot.Send(msg)
		if err != nil {
			b.sendErrorMessage(message.Chat.ID)
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
			b.sendErrorMessage(message.Chat.ID)
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
			b.sendErrorMessage(message.Chat.ID)
			return "", err
		}
		return "", errors.New("no argument")
	}
	return strings.Replace(message.Text, "/sendall", "", 1), nil

}
