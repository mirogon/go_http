package http_

import "net/http"

func HandleCORS(request *http.Request, responseWriter HttpResponseWriter) bool {
	if request.Method == "OPTIONS" {
		HandleOptionsRequest(responseWriter)
		return true
	}
	responseWriter.AddHeader("Access-Control-Allow-Origin", "http://localhost:5173")
	responseWriter.AddHeader("Access-Control-Allow-Credentials", "true")
	return false
}

func HandleOptionsRequest(httpResponseWriter HttpResponseWriter) {
	httpResponseWriter.AddHeader("Access-Control-Allow-Origin", "http://localhost:5173")
	httpResponseWriter.AddHeader("Access-Control-Allow-Headers", "Content-Type, Authorization")
	httpResponseWriter.AddHeader("Access-Control-Allow-Credentials", "true")
	httpResponseWriter.DeleteHeader("Content-Type")
	httpResponseWriter.DeleteHeader("Content-Length")
	httpResponseWriter.Send()
}
