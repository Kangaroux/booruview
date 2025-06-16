package gelbooru

import (
	"net/http"
	"net/url"

	api "github.com/jessienyan/booruview"
)

func httpGet(theUrl string, params url.Values) (*http.Response, error) {
	_url, err := url.Parse(theUrl + "?" + params.Encode())
	if err != nil {
		return nil, err
	}

	req := &http.Request{URL: _url, Header: http.Header{}}
	req.Header.Set("Cookie", "fringeBenefits=yup") // Enable all content

	return api.DoRequest(req)
}
