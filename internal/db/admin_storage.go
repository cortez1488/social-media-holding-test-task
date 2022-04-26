package db

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type adminStorage struct {
	db *sqlx.DB
}

func NewAdminStorage(db *sqlx.DB) *adminStorage {
	return &adminStorage{db: db}
}

func (r *adminStorage) CheckAdminRight(chatId int64) (bool, error) {
	query := fmt.Sprintf("SELECT id FROM users WHERE chatID = $1 and isadmin = TRUE")
	var t int
	err := r.db.Get(&t, query, chatId)
	if err == sql.ErrNoRows {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *adminStorage) GetAllUsersChatID() ([]int64, error) {
	query := fmt.Sprintf("SELECT chatID FROM users")
	var result []int64
	err := r.db.Select(&result, query)
	if err != nil {
		return nil, err
	}
	return result, err
}

func (r *adminStorage) AddAdmin(nickname string) (bool, error) {
	query := fmt.Sprintf("UPDATE users SET isadmin = 't' WHERE nickname = $1")
	_, err := r.db.Exec(query, nickname)
	if err != nil {
		return false, err
	}
	return true, nil
}
func (r *adminStorage) DeleteAdmin(nickname string) (bool, error) {
	query := fmt.Sprintf("UPDATE users SET isadmin = 'f' WHERE nickname = $1")
	_, err := r.db.Exec(query, nickname)
	if err != nil {
		return false, err
	}
	return true, nil
}
