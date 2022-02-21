package main

import (
	"net/http"
	"net/url"
)

func GetRequest() (*http.Request, error) {
	req, err := http.NewRequest("GET", "https://pro-api.coinmarketcap.com/v1/cryptocurrency/listings/latest", nil)
	if err != nil {
		return nil, err
	}

	q := url.Values{}
	q.Add("start", "1")
	q.Add("limit", "1")
	q.Add("convert", "USD")

	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", GetApiKey())
	req.URL.RawQuery = q.Encode()
	return req, nil

}
