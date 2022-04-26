package ip

import (
	"social-media-holding-test-task/structs"
)

type IpStorage interface {
	CreateUser(chatId int64, nickname string) (int, error)
	GetUser(chatId int64) (bool, structs.User, error)

	GetIpInfo(ip string) (structs.IPInfo, error)
	CreateIp(userId int, info structs.IPInfo) error
	GetIpsFromUser(userId int) (structs.UsersRequests, error)
}
