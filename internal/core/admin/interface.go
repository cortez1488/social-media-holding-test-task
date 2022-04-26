package admin

type AdminStorage interface {
	CheckAdminRight(chatId int64) (bool, error)
	GetAllUsersChatID() ([]int64, error)

	AddAdmin(nickname string) (bool, error)
	DeleteAdmin(nickname string) (bool, error)

	GetIdFromNickname(nickname string) (int, error)
}
