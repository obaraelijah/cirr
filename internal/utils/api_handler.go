package utils

import (
	"io"
	"net/http"
	"os"
)

func GetReq(endpoint string) string {
	logger := GetCirrLogger()

	logger.Printf("Calling API endpoint: %s", endpoint)
	response, err := http.Get(endpoint)

	if err != nil {
		logger.Println(err.Error())
		os.Exit(1)
	}

	responseData, err := io.ReadAll(response.Body)

	if err != nil {
		logger.Fatal(err)
	}

	return string(responseData)
}
