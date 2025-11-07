package testnode

import "net/http"

func Hello() string {
	return "HI"
}

func Request() (string, error) {
	req, err := http.NewRequest("GET", "https://google.com", nil)
	if err != nil {
		return "", err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	return resp.Status, nil
}
