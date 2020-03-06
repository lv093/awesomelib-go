package cl

import (
	"net/http"
	"io"
)

func GetClHtmlData(url string) (io.Reader, error) {
	return doGetRequest(url, nil, nil)
}

func doGetRequest(url string, param []string, headers []string) (io.Reader, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	return res.Body, nil
}