package rest

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type IPInfo struct {
	id        int
	IP        string  `json:"ip"`
	Continent string  `json:"continent_name"`
	Country   string  `json:"country_name"`
	Region    string  `json:"region_name"`
	City      string  `json:"city"`
	Zip       string  `json:"zip"`
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}

const APIIURI = "http://api.ipstack.com/"
const ACCESS_KEY = "23608c246cd680f479368fab5a2ceea2"

func GetIpInfo(ip string) IPInfo {
	uri := getFullUri(APIIURI, ip, ACCESS_KEY)
	resp, err := http.Get(uri)
	if err != nil {
		log.Fatalln("Request to API doesn't work " + err.Error())
	}

	responseBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("Can't get bytes from response, read bytes: " + err.Error())
	}
	err = resp.Body.Close()
	if err != nil {
		log.Fatalln("Can't close body of response: " + err.Error())
	}

	var response IPInfo
	err = json.Unmarshal(responseBytes, &response)
	if err != nil {
		log.Fatalln("Can't unmarshal bytes into struct: " + err.Error())
	}

	return response

}

func getFullUri(url, ip, accessKey string) string {
	return fmt.Sprintf("%s%s?access_key=%s", url, ip, accessKey)
}
