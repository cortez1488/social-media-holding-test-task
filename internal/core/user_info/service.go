package user_info

import (
	"social-media-holding-test-task/internal/core/ip"
	"social-media-holding-test-task/structs"
)

type userInfoService struct {
	userInfoStorage userInfoStorage
	ipStorage       ip.IpStorage
}

func NewUserInfoService(userInfoStorage userInfoStorage, ipStorage ip.IpStorage) *userInfoService {
	return &userInfoService{userInfoStorage: userInfoStorage, ipStorage: ipStorage}
}

func (s *userInfoService) DeleteIp(ip string) error {
	return s.userInfoStorage.DeleteIp(ip)
}

func (s *userInfoService) GetHistoryByTgID(chatID int64) (structs.UsersRequests, error) {
	_, user, err := s.ipStorage.GetUser(chatID)
	if err != nil {
		return structs.UsersRequests{}, err
	}
	result, err := s.ipStorage.GetIpsFromUser(user.Id)
	if err != nil {
		return structs.UsersRequests{}, err
	}
	return result, nil
}
func (s *userInfoService) GetUsers() ([]structs.User, error) {
	return s.userInfoStorage.GetUsers()
}
func (s *userInfoService) GetUser(id int) (structs.User, error) {
	return s.userInfoStorage.GetUser(id)
}
