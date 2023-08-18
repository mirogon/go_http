package http_

type HttpResponseWriter interface {
	AddHeader(name string, value string)
	DeleteHeader(name string)
	AddCookie(cookie Cookie)
	SetBody(body []byte)
	Send()
}

// Interface for a HttpResponseWriter that is implemented with the go http.ResponseWriter
