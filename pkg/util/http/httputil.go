package httputil

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// DoGet does GET request
func DoGet(timeOutSecs time.Duration , rootURL string, path string, params string) (string, error) {

	customClient := &http.Client{Timeout: timeOutSecs * time.Second}
	resp, err := customClient.Get(rootURL + "/" + path + "?" + params)

	if err != nil {
		log.Fatal(err)

		return "", err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode == http.StatusNotFound {
		err = errors.New("StatusNotFound")
	}

	bodyString := string(body)

	return bodyString, err
}
