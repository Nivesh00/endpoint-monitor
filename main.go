package main

import (
	// "errors"
	"encoding/json"
	"fmt"
	"io"
	"os"
	// "time"
	"log/slog"
	"github.com/Nivesh00/endpoint-monitor/my_modules"
	"github.com/Nivesh00/endpoint-monitor/my_templates"
)

func main() {

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	// Open urls.json file
	jsonUrlsFile, err := os.Open("my_templates/urls.json")
	if err != nil {
		logger.Error("Cannot open urls.json file", "err", err)
		os.Exit(1)
	}
	logger.Info("Successfully opened urls.json")
	defer jsonUrlsFile.Close()

	/*
		Parse file into bytes, then parse into struct
	*/

	byteValue, err := io.ReadAll(jsonUrlsFile)
	var urls my_templates.Urls
	if err != nil {
		logger.Error("Cannot read urls.json file", "err", err)
		os.Exit(1)
	}
	json.Unmarshal(byteValue, &urls)
	logger.Info("Successfully read urls.json")

	/*
		Read all data fron struct
	*/
	for _, url_item := range urls.Urls {
		
		// item.ToStr()
		resp, err := my_modules.SendRequest(&url_item.Endpoint)
		if err != nil {
			logger.Warn("Error for endpoint " + url_item.Endpoint, "error", err)
			continue
		}

		/*
			Check whether responses (do not) contain expressions
		*/
		status, err := my_modules.ValidateResponse(&resp, &url_item.Contains, &url_item.NotContains)
		if err != nil {
		
			msg := fmt.Sprintf("Url %s returns an invalid response", url_item.Endpoint)
			logger.Warn(msg, "status", err)
		} else {

			msg := fmt.Sprintf("Url %s returns an valid response", url_item.Endpoint)
			logger.Info(msg, "status", "ok")
		}

		fmt.Println(status)

	}
}