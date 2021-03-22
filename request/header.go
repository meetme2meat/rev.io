package request

import (
	"context"
	"net/http"
)

const (
	ACCEPT       = "application/json"
	CONTENT_TYPE = "application/json"
)

type Header struct {
	Authorization string
	Accept        string
	ContentType   string
}

func HeaderFrom(ctx context.Context) *Header {
	rh := ctx.Value("HEADERS").(*Header)
	if rh == nil {
		return &Header{}
	}

	if rh.Accept == "" {
		rh.Accept = ACCEPT
	}

	if rh.ContentType == "" {
		rh.ContentType = CONTENT_TYPE
	}

	return rh
}

func (rh *Header) CopyTo(r *http.Request) {
	r.Header.Add("Authorization", rh.Authorization)
	r.Header.Add("Accept", rh.Accept)
	r.Header.Add("Content-Type", rh.ContentType)
}
