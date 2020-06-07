package service

import (
	"eer-edgex/utils"
	"fmt"
	"io/ioutil"
	"net/http"
)

func RegisterDevice(deviceName string) {
	url := "http://localhost:48081/api/v1/device"
	content := `{
		"name": ` + deviceName + `,` +
		`"description":"Monitor and Collect students' physical data and movement",
    	"adminState":"unlocked",
    	"operatingState":"enabled",
    	"protocols":{"device protocol":{"device address":"device 1"}},
    	"labels": [],
    	"location":"",
    	"service":{"name":"GraduationDesignSystem control device service"},
    	"profile":{"name":"device monitor profile"}
	}`

	_ = utils.PostData(url, content)
}

func GetDataFromDevice(deviceName string) string {
	url := `http://localhost:48080/api/v1/event/device/` + deviceName + `/10`
	res, _ := http.Get(url)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	deviceData := string(body)
	//fmt.Println(deviceData)
	return deviceData
}

func DeleteDevice(deviceName string) {
	url := "http://localhost:48081/api/v1/device/name/" + deviceName

	req, _ := http.NewRequest("DELETE", url, nil)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}
