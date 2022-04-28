package db

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"social-media-holding-test-task/structs"
)

type userInfoStorage struct {
	db *sqlx.DB
}

func NewUserInfoStorage(db *sqlx.DB) *userInfoStorage {
	return &userInfoStorage{db: db}
}

func (r *userInfoStorage) GetUsers() ([]structs.User, error) {
	query := fmt.Sprintf("SELECT * FROM users")
	var data []structs.User
	err := r.db.Select(&data, query)
	if err != nil {
		return nil, err
	}

	return data, nil
}
func (r *userInfoStorage) GetUser(id int) (structs.User, error) {
	query := fmt.Sprintf("SELECT * FROM users WHERE id = $1")
	var data structs.User
	err := r.db.Get(&data, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return structs.User{}, nil
		}
		return structs.User{}, err
	}

	return data, nil
}

func (r *userInfoStorage) DeleteIp(ip string) error {
	query := fmt.Sprintf("DELETE FROM ip_info WHERE ip = $1")
	_, err := r.db.Exec(query, ip)
	if err != nil {
		return err
	}
	return nil
}
