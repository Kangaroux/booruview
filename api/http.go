package api

import (
	"log"
	"net/http"
	"time"
)

func DoRequest(req *http.Request) (*http.Response, error) {
	earlier := time.Now()
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Print(err)
	} else {
		log.Printf("GET  [%1.4fs]  %d  %s", time.Since(earlier).Seconds(), resp.StatusCode, req.URL)
	}

	return resp, err
}
