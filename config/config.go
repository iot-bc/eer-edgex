package config

import (
	"fmt"
	"net/http"
	"strings"
)

func Config() {
	addAddressables()
	postDeviceProfile()
	addDeviceService()
}

func addAddressables() {
	url := "http://120.26.172.10:48081/api/v1/addressable"
	deviceserviceAddressables := `{
    	"name":"eer-deivce control",
    	"protocol":"HTTP",
    	"address":"120.26.172.10",
    	"port":49977,
    	"path":"/eer-control",
    	"publisher":"none",
    	"user":"none",
    	"password":"none",
    	"topic":"none"
		}`

	deviceAddressables := `{
    	"name":"eer-device address",
    	"protocol":"HTTP",
    	"address":"120.26.172.10",
    	"port":49999,
    	"path":"/eer-address",
    	"publisher":"none",
    	"user":"none",
    	"password":"none",
    	"topic":"none"
		}`

	httpPost(url, deviceserviceAddressables)
	httpPost(url, deviceAddressables)
}

func postDeviceProfile() {

}

func addDeviceService() {
	url := "http://120.26.172.10:48081/api/v1/deviceservice"
	content := `{
    	"name":"eer device service",
    	"description":"Human health management device monitor profile",
    	"labels":[],
    	"adminState":"unlocked",
    	"operatingState":"enabled",
    	"addressable":
		{"name":"eer-deivce control"}
		}`
	httpPost(url, content)
}

func httpPost(url string, content string) {
	contentType := "application/json"
	data := strings.NewReader(fmt.Sprintf("%s", content))
	resp, _ := http.Post(url, contentType, data)
	fmt.Println(resp)
}
