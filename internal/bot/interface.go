package bot

import (
	"social-media-holding-test-task/structs"
)

type IpService interface {
	ProcessIp(chatID int64, nickname string, ip structs.IPInfo) error
	GetAllIps(userId int) (structs.UsersRequests, error)

	GetUser(chatId int64) (int, error)
	GetIpInfo(ip string) (structs.IPInfo, error)
}
