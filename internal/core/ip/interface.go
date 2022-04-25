package ip

import (
	"social-media-holding-test-task/structs"
)

type IpStorage interface {
	CreateUser(chatId int64, nickname string) (int, error)
	GetUser(chatId int64) (bool, int, error)

	CreateIp(userId int, info structs.IPInfo) error
	GetIpsFromUser(userId int) ([]structs.IPInfo, error)
}
