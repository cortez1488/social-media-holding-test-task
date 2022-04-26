package structs

import "time"

type User struct {
	Id       int    `db:"id"`
	Nickname string `db:"nickname"`
	ChatID   int64  `db:"chatid"`
	IsAdmin  bool   `db:"isadmin"`
}

type IPInfo struct {
	Id        int     `db:"id"`
	IP        string  `json:"ip" db:"ip"`
	Continent string  `json:"continent_name" db:"continent_name"`
	Country   string  `json:"country_name" db:"country_name"`
	Region    string  `json:"region_name" db:"region_name"`
	City      string  `json:"city" db:"city"`
	Zip       string  `json:"zip" db:"zip"`
	Latitude  float32 `json:"latitude" db:"latitude"`
	Longitude float32 `json:"longitude" db:"longitude"`
}

type UsersRequests struct {
	Ips map[string][]time.Time
}
