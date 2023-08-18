package http_

import "net/http"

type HttpRequestSenderImpl struct{}

func (sender HttpRequestSenderImpl) SendRequest(req *http.Request) (*http.Response, error) {
	return http.DefaultClient.Do(req)
}
