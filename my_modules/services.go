package my_modules

import (
	"io"
	"log/slog"
	"net/http"
	"strconv"
)


func SendRequest(url *string) string {

	resp, err := http.Get(*url)
	if (err != nil) {
		slog.Warn("Response error for url", *url, slog.Any("err", err))
		return ""

	} else if (resp.StatusCode != http.StatusOK) {
		error_msg := "Status code for" + *url + "was" + strconv.Itoa(resp.StatusCode)
		slog.Warn("HTTP statuscode error", slog.Any("err", error_msg))
		return ""

	} 

	bodyBytes, err := io.ReadAll(resp.Body)
	if (err != nil) {
		slog.Warn("Cannot read response body for URL", *url, slog.Any("err", err))
		return ""

	}
	return string(bodyBytes)
}
