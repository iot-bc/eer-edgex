package device

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Data struct {
	HeartRate         int     `json:"heart_rate"`
	Temperature       float64 `json:"temperature"`
	StepNumber        int     `json:"step_number"`
	ClimbHeight       float64 `json:"climb_height"`
	Calorie           float64 `json:"calorie"`
	SystolicPressure  int     `json:"systolic_pressure"`
	DiastolicPressure int     `json:"diastolic_pressure"`
}

func postData(deviceName string) error {
	randomData := generateRandomData(deviceName)
	jsonData, err := json.Marshal(randomData)
	url := "http://localhost:48080/api/v1/event" //请求地址
	contentType := "application/json"
	content := `{"device":` + deviceName + `, "readings":[{"name":"eer-data", "value":` + string(jsonData) + `}]}`

	data := strings.NewReader(fmt.Sprintf("%s", content))
	resp, _ := http.Post(url, contentType, data)
	fmt.Println(resp)
	return err
}

func generateRandomData(deviceName string) Data {
	// 设备的旧数据
	oldData := getDataFromDevice(deviceName)

	rand.Seed(time.Now().UnixNano())
	heartRate := rand.Intn(110-70) + 70             // 心率在70到110之间波动
	temperature := float64(rand.Intn(10))/10 + 36.5 // 温度36.5到37.5之间波动
	temperature = FloatRound1(temperature)

	stepNumber := oldData.StepNumber
	stepNumber += rand.Intn(7) // 每隔10秒增加7以内的随机步数

	climbHeight := oldData.ClimbHeight
	climbHeight += rand.Float64()
	climbHeight = FloatRound1(climbHeight)

	calorie := oldData.Calorie
	calorie += float64(rand.Intn(5)) + rand.Float64()
	calorie = FloatRound1(calorie)

	systolicPressure := rand.Intn(120-100) + 100 // 收缩压在100到120之间波动
	diastolicPressure := rand.Intn(80-60) + 60   // 舒张压在60到80之间波动
	newData := Data{heartRate, temperature, stepNumber, climbHeight,
		calorie, systolicPressure, diastolicPressure}

	return newData
}

func getDataFromDevice(deviceName string) Data {
	url := "http://localhost:48080/api/v1/event/device/" + deviceName + "/1"
	res, _ := http.Get(url)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	deviceData := string(body)

	var result Data
	if err := json.Unmarshal([]byte(deviceData), &result); err == nil {
		fmt.Println(result)
	}
	return result
}

// 截取1位小数
func FloatRound1(value float64) float64 {
	res, _ := strconv.ParseFloat(fmt.Sprintf("%.1f", value), 64)
	return res
}
