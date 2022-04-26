package admin

type AdminService struct {
	storage AdminStorage
}

func NewAdminService(storage AdminStorage) *AdminService {
	return &AdminService{storage: storage}
}

func (s *AdminService) CheckAdminRight(chatId int64) (bool, error) {
	return s.storage.CheckAdminRight(chatId)
}

func (s *AdminService) GetAllUsersChatID() ([]int64, error) {
	return s.storage.GetAllUsersChatID()
}
func (s *AdminService) AddAdmin(nickname string) (bool, error) {
	return s.storage.AddAdmin(nickname)
}
func (s *AdminService) DeleteAdmin(nickname string) (bool, error) {
	return s.storage.DeleteAdmin(nickname)
}

func (s *AdminService) GetIdFromNickname(nickname string) (int, error) {
	return s.storage.GetIdFromNickname(nickname)
}
