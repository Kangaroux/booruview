package api

import (
	"log"
	"net/http"
	"time"
)

var httpClient = &http.Client{Timeout: 5 * time.Second}

func DoRequest(req *http.Request) (*http.Response, error) {
	earlier := time.Now()
	resp, err := httpClient.Do(req)

	if err != nil {
		log.Print(err)
	} else {
		log.Printf("%s  [%1.4fs]  %d  %s", req.Method, time.Since(earlier).Seconds(), resp.StatusCode, req.URL)
	}

	return resp, err
}
