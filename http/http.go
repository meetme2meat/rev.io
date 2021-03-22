package http

import (
	"context"
	"net/http"

	"github.com/meetme2meat/rev.io/request"
)

type HTTPClient interface {
	GET(context.Context, string) (*http.Response, error)
	POST(context.Context, string) (*http.Response, error)
	DELETE(context.Context, string) (*http.Response, error)
	PUT(context.Context, string) (*http.Response, error)
	PATCH(context.Context, string) (*http.Response, error)
}

type HTTP struct {
}

func New() *HTTP {
	return &HTTP{}
}

func (h *HTTP) GET(ctx context.Context, url string) (*http.Response, error) {
	return nil, nil
}

func (h *HTTP) POST(ctx context.Context, url string) (*http.Response, error) {
	req, _ := http.NewRequest("POST", url, nil)
	header := request.HeaderFrom(ctx)
	header.CopyTo(req)
	return http.DefaultClient.Do(req)
}

func (h *HTTP) PUT(ctx context.Context, url string) (*http.Response, error) {
	return nil, nil
}

func (h *HTTP) DELETE(ctx context.Context, url string) (*http.Response, error) {
	return nil, nil
}

func (h *HTTP) PATCH(ctx context.Context, url string) (*http.Response, error) {
	return nil, nil
}
