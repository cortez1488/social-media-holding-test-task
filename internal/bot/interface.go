package bot

import "social-media-holding-test-task/internal/handler/rest"

type IpService interface {
	ProcessIp(chatID int64, nickname string, ip rest.IPInfo) error
	GetAllIps(chatID int64) ([]string, error)
}
