package main

import (
	"fmt"
	"io"
	"net/http"
)

func getHTML(rawURL string) (string, error) {

	res, err := http.Get(rawURL)

	if err != nil {
		return "", fmt.Errorf("error")
	}

	defer res.Body.Close()

	if res.StatusCode > 399 {
		return "", fmt.Errorf("unsuccess request")
	}

	if res.Header.Get("Content-Type") != "text/html" {
		return "", fmt.Errorf("bad content-type header")
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return "", fmt.Errorf("error")
	}

	return string(body), nil
}
