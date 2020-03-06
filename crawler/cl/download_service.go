package cl

import (
	"net/http"
	"os"
	"io"
)

func DownloadImg(url string, path string) (int64, error) {
	res, err := http.Get(url)
	if err != nil {
		return -1, err
	}
	//file, err := os.Create(path)
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return -1, err
	}
	return io.Copy(file, res.Body)
}