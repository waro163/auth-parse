package utils

import "net/http"

type IRequestClient interface {
	Do(*http.Request) (*http.Response, error)
}

var DefaultClient IRequestClient = http.DefaultClient
