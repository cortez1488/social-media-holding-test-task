package ip

import (
	"errors"
	"fmt"
	"social-media-holding-test-task/internal/handler/rest"
)

type ipService struct {
	storage IpStorage
}

func NewIpService(storage IpStorage) *ipService {
	return &ipService{storage: storage}
}

func (s *ipService) ProcessIp(chatID int64, nickname string, ip rest.IPInfo) error {
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

func (s *ipService) GetAllIps(chatID int64) ([]string, error) {
	result, err := s.storage.GetIpsFromUser(chatID)
	fmt.Println(result)
	if err != nil {
		return nil, errors.New("s.storage.GetIpsFromUser: " + err.Error())
	}
	return []string{}, nil
}
