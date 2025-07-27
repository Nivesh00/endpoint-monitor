package my_modules

import (
	"errors"
	"fmt"
	"io"
	// "log/slog"
	"net/http"
	// "strconv"
)


func SendRequest(url *string) (string, error) {

	resp, err := http.Get(*url)
	if (err != nil) {
		return "", err
	} else if (resp.StatusCode != http.StatusOK) {
		err_msg := fmt.Sprintf("HTTP statuscode for url %s is not valid. Code was %d", *url, resp.StatusCode)
		return "", errors.New(err_msg)

	} 

	bodyBytes, err := io.ReadAll(resp.Body)
	if (err != nil) { return "", err }

	resp_body := string(bodyBytes)
	return resp_body, nil
}
