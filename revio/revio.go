package revio

import (
	"encoding/base64"
	"fmt"

	"github.com/meetme2meat/rev.io/http"
	"github.com/meetme2meat/rev.io/inventory"
)

type Revio struct {
	username string
	code     string
	password string
	http.HTTPClient
}

func New(username, client_code, password string) Revio {
	r := Revio{
		username: username,
		code:     client_code,
		password: password,
	}

	r.HTTPClient = &http.HTTP{}

	return r
}

func NewWithHTTPClient(username, client_code, password string, httpclient http.HTTPClient) Revio {
	return Revio{
		username,
		client_code,
		password,
		httpclient,
	}
}

func (r Revio) base64Encode() string {
	// base64 encode of username@client_code:password
	return base64.StdEncoding.EncodeToString(
		[]byte(fmt.Sprintf("%s@%s:%s", r.username, r.code, r.password)),
	)
}

func (r Revio) Authorization() string {
	return "Basic: " + r.base64Encode()
}

func (r Revio) Inventory() inventory.Inventory {
	return inventory.Inventory{HTTPClient: r.HTTPClient}
}

// func (r Revio) Audits() audits.Audit {
// 	return audits.Audit{HTTPClient: r.HTTPClient}
// }

// package main
// import "github.com/meetme2meat/rev.io/revio"
// revio.New().Inventory().Create()
//
// r.Inventory()
