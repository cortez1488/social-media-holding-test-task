package rest

import "social-media-holding-test-task/structs"

type UserInfoService interface {
	GetUsers() ([]structs.User, error)
	GetUser(id int) (structs.User, error)
	GetHistoryByTgID(chatID int64) (structs.UsersRequests, error)
	DeleteIp(ip string) error
}
