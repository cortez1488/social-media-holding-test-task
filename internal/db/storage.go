package db

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"social-media-holding-test-task/structs"
	"time"
)

type ipStorage struct {
	db *sqlx.DB
}

func NewIpStorage(db *sqlx.DB) *ipStorage {
	return &ipStorage{db: db}
}

func (r *ipStorage) CreateUser(chatId int64, nickname string) (int, error) {
	query := fmt.Sprintf("INSERT INTO %s (nickname, chatID) VALUES ($1, $2) RETURNING id", "users")
	row := r.db.QueryRowx(query, nickname, chatId)

	var id int
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *ipStorage) GetUser(chatId int64) (bool, int, error) {
	query := fmt.Sprintf("SELECT id FROM %s WHERE chatID = $1", "users")
	var id int
	err := r.db.Get(&id, query, chatId)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, 0, nil
		} else {
			return false, 0, err
		}
	}
	if id == 0 {

		return false, 0, nil

	}
	return true, id, nil
}

func (r *ipStorage) CreateIp(userId int, info structs.IPInfo) error {

	query := fmt.Sprintf("SELECT id FROM %s WHERE ip = $1 RETURNING id", "ip_info")
	row := r.db.QueryRow(query, info.IP)
	var temp int
	err := row.Scan(&temp)
	if err != nil {
		if err == sql.ErrNoRows {

			tx, _ := r.db.Beginx()
			query = fmt.Sprintf("INSERT INTO %s (ip, continent_name, country_name,"+
				"region_name, city, zip, latitude, longitude) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id", "ip_info")
			row = tx.QueryRow(query, info.IP, info.Continent, info.Country, info.Region, info.City, info.Zip,
				info.Latitude, info.Longitude)

			var ipId int
			err := row.Scan(&ipId)
			if err != nil {
				tx.Rollback()
				return err
			}

			query = fmt.Sprintf("INSERT INTO %s (ip_id, user_id) VALUES ($1, $2) RETURNING id",
				"user_searched_ip")
			row = tx.QueryRow(query, ipId, userId)
			var searchedId int
			err = row.Scan(&searchedId)
			if err != nil {
				tx.Rollback()
				return err
			}

			query = fmt.Sprintf("INSERT INTO %s (user_searched_ip_id, timedate) VALUES ($1, $2)", "search_date")
			row = tx.QueryRow(query, searchedId, time.Now())

			return tx.Commit()
		}
	} else {
		return err
	}
	return nil

}

func (r *ipStorage) GetIpsFromUser(userId int) ([]structs.IPInfo, error) {
	fmt.Println("storage logic")
	return nil, nil
}
