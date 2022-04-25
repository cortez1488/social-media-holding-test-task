package ip

import "social-media-holding-test-task/internal/handler/rest"

type IpStorage interface {
	CreateUser(chatId int64, nickname string) (int, error)
	GetUser(chatId int64) (bool, int, error)

	CreateIp(userId int, info rest.IPInfo) error
	GetIpsFromUser(chatId int64) ([]rest.IPInfo, error)
}
