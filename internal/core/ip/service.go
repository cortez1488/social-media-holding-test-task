package ip

import (
	"errors"
	"social-media-holding-test-task/structs"
)

type ipService struct {
	storage IpStorage
}

func NewIpService(storage IpStorage) *ipService {
	return &ipService{storage: storage}
}

func (s *ipService) ProcessIp(chatID int64, nickname string, ip structs.IPInfo) error {
	exists, id, err := s.storage.GetUser(chatID)
	if err != nil {
		return errors.New("s.storage.GetUser: " + err.Error())
	}
	if exists {
		err := s.storage.CreateIp(id, ip)
		if err != nil {
			return errors.New("s.storage.CreateIp: " + err.Error())
		}

	} else {
		id, err = s.storage.CreateUser(chatID, nickname)
		if err != nil {
			return errors.New("s.storage.CreateUser: " + err.Error())
		}

		err := s.storage.CreateIp(id, ip)
		if err != nil {
			return errors.New("s.storage.CreateIp: " + err.Error())
		}
	}
	return nil
}

func (s *ipService) GetAllIps(userId int) (structs.UsersRequests, error) {
	result, err := s.storage.GetIpsFromUser(userId)

	if err != nil {
		return structs.UsersRequests{}, errors.New("s.storage.GetIpsFromUser: " + err.Error())
	}
	return result, nil
}

func (s *ipService) GetUser(chatId int64) (int, error) {
	_, id, err := s.storage.GetUser(chatId)
	if err != nil {
		return 0, errors.New("s.storage.GetUser: " + err.Error())
	}

	return id, nil
}

func (s *ipService) GetIpInfo(ip string) (structs.IPInfo, error) {
	return s.storage.GetIpInfo(ip)
}
