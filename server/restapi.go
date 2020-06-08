package server

import (
	"eer-edgex/device/dbs"
	"eer-edgex/privacy"
	UserService "eer-edgex/service"
	"encoding/json"
	"fmt"
	"github.com/drone/routes"
	"log"
	"net/http"
)

func Receiver() (string, string) {
	deviceName := ""
	cmd := ""
	return deviceName, cmd
}

type Resp struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

type Body struct {
	UserID     string `json:"id"`
	DeviceName string `json:"name"`
}

type Data struct {
	HeartRate         int     `json:"heart_rate"`
	Temperature       float64 `json:"temperature"`
	StepNumber        int     `json:"step_number"`
	ClimbHeight       float64 `json:"climb_height"`
	Calorie           float64 `json:"calorie"`
	SystolicPressure  int     `json:"systolic_pressure"`
	DiastolicPressure int     `json:"diastolic_pressure"`
}

func RESTAPI() {

	mux := routes.New()
	mux.Get("/api/device/:device-name", GetRoute)
	mux.Del("/api/device/:device-name", DeleteRoute)
	//mux.Put("/api/device/:device-name", GetRoute)
	mux.Post("/api/device", PostRoute)

	http.Handle("/", mux)
	_ = http.ListenAndServe(":9000", nil)
	fmt.Println("rest api stop")
}

// GET http://localhost:9000/api/device/device-name
func GetRoute(writer http.ResponseWriter, request *http.Request) {
	params := request.URL.Query()
	name := params.Get(":device-name")
	//_, _ = fmt.Fprintf(writer, name)

	// 在区块链上与deviceName相对应的fakeName
	fakeName := privacy.AnonymizeDevice(name)
	deviceData := UserService.GetDataFromDevice(fakeName)

	decryptData := privacy.AESDecryptData(deviceData)

	var result Data
	if err := json.Unmarshal([]byte(decryptData), &result); err == nil {
		fmt.Println(result)
	}

	if err := json.NewEncoder(writer).Encode(result); err != nil {
		log.Fatal(err)
	}
}

// POST http://localhost:9000/api/device
func PostRoute(writer http.ResponseWriter, request *http.Request) {
	var body Body
	if err := json.NewDecoder(request.Body).Decode(&body); err != nil {
		_ = request.Body.Close()
		log.Fatal(err)
	}

	//userId := body.UserID
	deviceName := body.DeviceName

	// 通过区块链生成对应的fakeName
	fakeName := privacy.AnonymizeDevice(deviceName)

	// 向edgex注册设备
	UserService.RegisterDevice(fakeName)

	// 在本地数据库中存储该设备名称
	dbs.AddDevice(fakeName)

	// 返回信息
	var resp Resp
	resp.Code = "200"
	resp.Msg = "Registered Successfully"

	if err := json.NewEncoder(writer).Encode(resp); err != nil {
		log.Fatal(err)
	}
}

// Delete http://localhost:9000/api/device/device-name
func DeleteRoute(writer http.ResponseWriter, request *http.Request) {
	params := request.URL.Query()
	name := params.Get(":device-name")

	// 通过区块链获取对应的fakeName
	fakeName := privacy.AnonymizeDevice(name)

	// 在edgex上删除设备
	UserService.DeleteDevice(fakeName)

	// 在本地数据库删除该设备
	dbs.DeleteDevice(fakeName)

	var resp Resp

	resp.Code = "200"
	resp.Msg = "Deleted Successfully"

	if err := json.NewEncoder(writer).Encode(resp); err != nil {
		log.Fatal(err)
	}
}
