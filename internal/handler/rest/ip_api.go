package rest

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"io"
	"net/http"
	"social-media-holding-test-task/structs"
)

func GetIpInfo(ip string) structs.IPInfo {
	uri := getFullUri(viper.GetString("api.ipapi"), ip, viper.GetString("api.access_key"))
	resp, err := http.Get(uri)
	if err != nil {
		logrus.Fatalln("Request to API doesn't work " + err.Error())
	}

	responseBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		logrus.Fatalln("Can't get bytes from response, read bytes: " + err.Error())
	}
	err = resp.Body.Close()
	if err != nil {
		logrus.Fatalln("Can't close body of response: " + err.Error())
	}

	var response structs.IPInfo
	err = json.Unmarshal(responseBytes, &response)
	if err != nil {
		logrus.Fatalln("Can't unmarshal bytes into struct: " + err.Error())
	}

	return response

}

func getFullUri(url, ip, accessKey string) string {
	return fmt.Sprintf("%s%s?access_key=%s", url, ip, accessKey)
}
