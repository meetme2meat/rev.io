package http

import (
	"net/http"
)

type HTTPClient interface {
	GET(string) (*http.Response, error)
	POST(string) (*http.Response, error)
	DELETE(string) (*http.Response, error)
	PUT(string) (*http.Response, error)
	PATCH(string) (*http.Response, error)
}

type HTTP struct {
	Authorization string
	Accept        string
}

func New(auth string) *HTTP {
	return &HTTP{Authorization: auth}
}

func (h *HTTP) GET(url string) (*http.Response, error) {
	return nil, nil
}

func (h *HTTP) POST(url string) (*http.Response, error) {
	req, _ := http.NewRequest("POST", url, nil)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", h.Authorization)
	req.Header.Add("Content-Type", "text/json")
	return http.DefaultClient.Do(req)
}

func (h *HTTP) PUT(url string) (*http.Response, error) {
	return nil, nil
}

func (h *HTTP) DELETE(url string) (*http.Response, error) {
	return nil, nil
}

func (h *HTTP) PATCH(url string) (*http.Response, error) {
	return nil, nil
}
