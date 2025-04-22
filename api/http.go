package api

import (
	"log"
	"net/http"
	"time"
)

func HttpGet(url string) (*http.Response, error) {
	earlier := time.Now()
	resp, err := http.Get(url)
	log.Printf("GET  [%1.4fs]  %d  %s", time.Since(earlier).Seconds(), resp.StatusCode, url)
	return resp, err
}
