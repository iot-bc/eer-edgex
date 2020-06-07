package utils

import (
	"fmt"
	"net/http"
	"strings"
)

func PostData(url string, data string) error {
	contentType := "application/json"

	content := strings.NewReader(fmt.Sprintf("%s", data))
	resp, err := http.Post(url, contentType, content)
	fmt.Println(resp)
	return err
}
