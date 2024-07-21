package utils

import "net/http"

type IReqeustClient interface {
	Do(*http.Request) (*http.Response, error)
}

var DefaultClient IReqeustClient = http.DefaultClient
