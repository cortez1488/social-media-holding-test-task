package user_info

import "social-media-holding-test-task/structs"

type userInfoStorage interface {
	GetUsers() ([]structs.User, error)
	GetUser(id int) (structs.User, error)

	DeleteIp(ip string) error
}
