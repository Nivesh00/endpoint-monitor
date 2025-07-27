package main

import (
	"fmt"
	// "errors"
	"encoding/json"
	"io"
	"os"
	// "time"
	"log/slog"
	"github.com/Nivesh00/endpoint-monitor/my_modules"
	"github.com/Nivesh00/endpoint-monitor/my_templates"
)

func main() {

	// Open urls.json file
	jsonUrlsFile, err := os.Open("my_templates/urls.json")
	if err != nil {
		slog.Error("Cannot open urls.json file", slog.Any("err", err))
		os.Exit(1)
	}
	slog.Info("Successfully opened urls.json")
	defer jsonUrlsFile.Close()

	/*
		Parse file into bytes, then parse into struct
	*/

	byteValue, err := io.ReadAll(jsonUrlsFile)
	var urls my_templates.Urls
	if err != nil {
		slog.Error("Cannot read urls.json file", slog.Any("err", err))
		os.Exit(1)
	}
	json.Unmarshal(byteValue, &urls)
	slog.Info("Successfully read urls.json")

	/*
		Read all data fron struct
	*/

	for i := 0; i < len(urls.Urls); i++ {
		
		// url_data := urls.Urls[i].ToStr()
		// Reference current URL
		var curr_url *my_templates.Url = &urls.Urls[i]

		// Reference attributes
		var curr_endpt 	  *string = &curr_url.Endpoint
		// var curr_fmt 	  *string = &curr_url.ResponseFormat
		// var curr_in 	*[]string = &curr_url.Contains
		// var curr_not_in *[]string = &curr_url.NotContains

		fmt.Println(*curr_endpt)
		resp := my_modules.SendRequest(curr_endpt)
		fmt.Println(resp)
	}
}