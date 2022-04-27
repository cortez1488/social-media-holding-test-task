package rest

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"io"
	"log"
	"net/http"
	"social-media-holding-test-task/structs"
)

var APIIURI = viper.Get("ipapi").(string)
var ACCESS_KEY = viper.Get("access_key").(string)

func GetIpInfo(ip string) structs.IPInfo {
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

	var response structs.IPInfo
	err = json.Unmarshal(responseBytes, &response)
	if err != nil {
		log.Fatalln("Can't unmarshal bytes into struct: " + err.Error())
	}

	return response

}

func getFullUri(url, ip, accessKey string) string {
	return fmt.Sprintf("%s%s?access_key=%s", url, ip, accessKey)
}
