package http_

import (
	"encoding/json"
	"fmt"
	"net/http"

	util "github.com/mirogon/go_util"
)

type HttpResponseWriterImpl struct {
	responseWriter http.ResponseWriter
	body           []byte
}

func CreateHttpResponseWriter(responseWriter http.ResponseWriter) HttpResponseWriterImpl {
	httpResponseWriter := HttpResponseWriterImpl{responseWriter: responseWriter}
	return httpResponseWriter
}

func (rw *HttpResponseWriterImpl) AddHeader(name string, value string) {
	rw.responseWriter.Header().Add(name, value)
}

func (rw *HttpResponseWriterImpl) DeleteHeader(name string) {
	rw.responseWriter.Header().Del(name)
}

func (rw *HttpResponseWriterImpl) AddCookie(cookie Cookie) {
	rw.responseWriter.Header().Add("Set-Cookie", cookie.GetValueString())
}

func (rw *HttpResponseWriterImpl) SetBody(body []byte) {
	rw.body = body
}

func (rw *HttpResponseWriterImpl) Send() {
	fmt.Fprintf(rw.responseWriter, "%s", rw.body)
}

func SendHttpResponse[V any](rw HttpResponseWriter, response V) {
	jsonData, _ := json.Marshal(response)
	rw.SetBody(jsonData)
	rw.Send()
}

func ParseRequestIntoRequestType[T any](req *http.Request) (T, error) {
	buffer, err := util.GetHttpRequestBody(req)
	if err != nil {
		var empty T
		return empty, err
	}

	var request T
	err = json.Unmarshal(buffer, &request)
	if err != nil {
		var empty T
		return empty, err
	}

	return request, nil
}
