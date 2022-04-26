package db

import (
	"database/sql"
	"errors"
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

func (r *ipStorage) GetUser(chatId int64) (bool, structs.User, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE chatID = $1", "users")
	var user structs.User
	err := r.db.Get(&user, query, chatId)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, structs.User{}, nil
		} else {
			return false, structs.User{}, err
		}
	}
	return true, user, nil
}

func (r *ipStorage) CreateIp(userId int, info structs.IPInfo) error {
	var err error

	query := fmt.Sprintf("SELECT id FROM %s WHERE ip = $1", "ip_info")
	var ipId int
	err = r.db.Get(&ipId, query, info.IP)

	if err != nil && err != sql.ErrNoRows {
		return err
	}
	if err == nil {
		tx, _ := r.db.Beginx()
		searchedId, err := userSearchedIpInsert(tx, ipId, userId)
		if err != nil {
			return err
		}

		err = searchDateInsert(tx, searchedId)
		if err != nil {
			return err
		}
		return tx.Commit()
	} else if err == sql.ErrNoRows {
		tx, _ := r.db.Beginx()

		ipId, err = createNewIp(tx, info)
		if err != nil {
			return err
		}

		searchedId, err := userSearchedIpInsert(tx, ipId, userId)
		if err != nil {
			return err
		}

		err = searchDateInsert(tx, searchedId)
		if err != nil {
			return err
		}
		return tx.Commit()
	}
	return errors.New("Unknown error: " + err.Error())
}

func createNewIp(tx *sqlx.Tx, info structs.IPInfo) (int, error) {
	query := fmt.Sprintf("INSERT INTO %s (ip, continent_name, country_name,"+
		"region_name, city, zip, latitude, longitude) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id", "ip_info")
	row := tx.QueryRow(query, info.IP, info.Continent, info.Country, info.Region, info.City, info.Zip,
		info.Latitude, info.Longitude)

	var ipId int
	err := row.Scan(&ipId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	return ipId, nil
}

func userSearchedIpInsert(tx *sqlx.Tx, ipId, userId int) (int, error) {
	query := fmt.Sprintf("INSERT INTO %s (ip_id, user_id) VALUES ($1, $2) RETURNING id",
		"user_searched_ip")
	row := tx.QueryRow(query, ipId, userId)
	var searchedId int
	err := row.Scan(&searchedId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	return searchedId, nil
}

func searchDateInsert(tx *sqlx.Tx, searchedId int) error {
	query := fmt.Sprintf("INSERT INTO %s (user_searched_ip_id, timedate) VALUES ($1, $2)", "search_date")
	_, err := tx.Exec(query, searchedId, time.Now())
	if err != nil {
		return err
	}
	return nil
}

type ipDateData struct {
	Ip   string    `db:"ip"`
	Date time.Time `db:"timedate"`
}

func (r *ipStorage) GetIpsFromUser(userId int) (structs.UsersRequests, error) {
	query := fmt.Sprintf("SELECT ip_info.ip, search_date.timedate FROM user_searched_ip " +
		"JOIN search_date ON search_date.user_searched_ip_id = user_searched_ip.id " +
		"JOIN ip_info ON user_searched_ip.ip_id = ip_info.id " +
		"JOIN users ON users.id = user_searched_ip.user_id " +
		"WHERE users.id = $1 ")

	var rawData []ipDateData
	err := r.db.Select(&rawData, query, userId)

	response := structs.UsersRequests{Ips: make(map[string][]time.Time)}

	for _, ip := range rawData {
		ips, exists := response.Ips[ip.Ip]
		if exists {
			response.Ips[ip.Ip] = append(ips, ip.Date)
		} else {
			response.Ips[ip.Ip] = []time.Time{ip.Date}
		}
	}

	if err != nil {
		return structs.UsersRequests{}, err
	}

	return response, nil
}

func (r *ipStorage) GetIpInfo(ip string) (structs.IPInfo, error) {
	query := fmt.Sprintf("SELECT * FROM ip_info where ip = $1")
	data := structs.IPInfo{}

	err := r.db.Get(&data, query, ip)
	if err != nil {
		return structs.IPInfo{}, err
	}

	return data, nil
}
