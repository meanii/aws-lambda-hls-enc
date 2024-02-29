package helper

import (
	"net/http"
)

// Fetch fetches the content of the given URL
func Fetch(url string) (*http.Response, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
