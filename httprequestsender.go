package http_

import "net/http"

type HttpRequestSender interface {
	SendRequest(*http.Request) (*http.Response, error)
}
