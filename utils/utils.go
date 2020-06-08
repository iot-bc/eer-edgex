package utils

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func PostData(url string, data string) error {
	contentType := "application/json"

	content := strings.NewReader(fmt.Sprintf("%s", data))
	resp, err := http.Post(url, contentType, content)
	fmt.Println(resp)
	return err
}

// 截取1位小数
func FloatRound1(value float64) float64 {
	res, _ := strconv.ParseFloat(fmt.Sprintf("%.1f", value), 64)
	return res
}
