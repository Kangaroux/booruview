package api

import (
	"net/http"
	"regexp"
	"time"

	"github.com/rs/zerolog/log"
)

var (
	httpClient   = &http.Client{Timeout: 5 * time.Second}
	reMaskApiKey = regexp.MustCompile(`(api_key=\w{7})(\w+)`)
)

func DoRequest(req *http.Request) (*http.Response, error) {
	earlier := time.Now()
	resp, err := httpClient.Do(req)

	method := req.Method
	if method == "" {
		method = "GET"
	}

	cleanUrl := reMaskApiKey.ReplaceAllString(req.URL.String(), "$1******")

	if err != nil {
		log.Err(err).Str("method", method).Str("url", cleanUrl).Msg("http request failed")
	} else {
		log.Info().
			Dur("duration", time.Since(earlier)).
			Int("status", resp.StatusCode).
			Str("method", method).
			Str("url", cleanUrl).
			Send()
	}

	return resp, err
}
