package driver

import (
	"crypto/tls"
	"io"
	"log"
	"net/http"
	"time"
)

func Get(url string, body io.Reader) *http.Response {
	timeout := time.Duration(10 * time.Second)
	client := &http.Client{
		Timeout: timeout,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	request, err := http.NewRequest("GET", url, body)
	if err != nil {
		log.Fatal(err)
	}

	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}

	return response
}

func Post(url string, body io.Reader) *http.Response {
	timeout := time.Duration(10 * time.Second)
	client := &http.Client{
		Timeout: timeout,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	request, err := http.NewRequest("POST", url, body)
	if err != nil {
		log.Fatal(err)
	}

	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	return response
}
