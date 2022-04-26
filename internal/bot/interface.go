package bot

import (
	"social-media-holding-test-task/structs"
)

type IpService interface {
	ProcessIp(chatID int64, nickname string, ip structs.IPInfo) error
	GetAllIps(userId int) (structs.UsersRequests, error)

	GetUser(chatId int64) (structs.User, error)
	GetIpInfo(ip string) (structs.IPInfo, error)
}

type AdminService interface {
	CheckAdminRight(chatId int64) (bool, error)
	GetAllUsersChatID() ([]int64, error)

	AddAdmin(nickname string) (bool, error)
	DeleteAdmin(nickname string) (bool, error)

	GetIdFromNickname(nickname string) (int, error)
}
