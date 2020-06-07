// -*- Mode: Go; indent-tabs-mode: t -*-
//
// Copyright (C) 2018 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package driver

import (
	HttpReceiver "eer-edgex/server"
	UserService "eer-edgex/service"
	"eer-edgex/utils"
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const (
	DELETE   string = "delete"
	REGISTER string = "register"
	GETDATA  string = "get"
	URL      string = "localhost:9000"
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

type myDevice struct {
	data string
}

func (d *myDevice) value() (string, error) {

	deviceName, cmd := HttpReceiver.Receiver()
	if strings.EqualFold(cmd, DELETE) {
		UserService.DeleteDevice(deviceName)

		feedback := "Deleted Successfully"

		_ = utils.PostData(URL, feedback)

	} else if strings.EqualFold(cmd, REGISTER) {
		UserService.RegisterDevice(deviceName)

		feedback := "Registered Successfully"

		_ = utils.PostData(URL, feedback)

	} else if strings.EqualFold(cmd, GETDATA) {
		data := UserService.GetDataFromDevice(deviceName)

		_ = utils.PostData(URL, data)

	}

	data, err := json.Marshal(randomData())
	return string(data), err
}

func newDevice() *myDevice {
	return &myDevice{data: ""}
}

func randomData() Data {
	rand.Seed(time.Now().UnixNano())

	heartRate := rand.Intn(110-70) + 70 // 心率在70到110之间波动

	temperature := float64(rand.Intn(10))/10 + 36.5 // 温度36.5到37.5之间波动
	temperature = FloatRound2(temperature)

	stepNumber := 0
	stepNumber += rand.Intn(7) // 每隔10秒增加7以内的随机步数

	climbHeight := 0.0
	climbHeight += rand.Float64()
	climbHeight = FloatRound2(climbHeight)

	calorie := 0.0
	calorie += float64(rand.Intn(5)) + rand.Float64()
	calorie = FloatRound2(calorie)

	systolicPressure := rand.Intn(120-100) + 100 // 收缩压在100到120之间波动

	diastolicPressure := rand.Intn(80-60) + 60 // 舒张压在60到80之间波动

	newData := Data{heartRate, temperature, stepNumber, climbHeight,
		calorie, systolicPressure, diastolicPressure}

	return newData
}

// 截取1位小数
func FloatRound2(value float64) float64 {
	res, _ := strconv.ParseFloat(fmt.Sprintf("%.1f", value), 64)
	return res
}
