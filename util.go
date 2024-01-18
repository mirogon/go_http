package http_

import "net/http"

func HandleCORS(request *http.Request, responseWriter HttpResponseWriter, origin string) bool {
	safeOrigin := OriginToSafeOrigin(origin)
	if request.Method == "OPTIONS" {
		HandlePreflightOptionsRequest(responseWriter, safeOrigin)
		return true
	}
	responseWriter.AddHeader("Access-Control-Allow-Origin", safeOrigin)
	responseWriter.AddHeader("Access-Control-Allow-Credentials", "true")
	responseWriter.AddHeader("Vary", "Referer")
	return false
}

func HandlePreflightOptionsRequest(httpResponseWriter HttpResponseWriter, safeOrigin string) {
	httpResponseWriter.AddHeader("Access-Control-Allow-Origin", safeOrigin)
	httpResponseWriter.AddHeader("Access-Control-Allow-Headers", "Content-Type, Authorization")
	httpResponseWriter.AddHeader("Access-Control-Allow-Credentials", "true")
	httpResponseWriter.AddHeader("Vary", "Referer")
	httpResponseWriter.DeleteHeader("Content-Type")
	httpResponseWriter.DeleteHeader("Content-Length")
	httpResponseWriter.Send()
}

func OriginToSafeOrigin(origin string) string {
	if origin == "https://polterai.com" {
		return "https://polterai.com"
	} else if origin == "http://localhost:3000" {
		return "http://localhost:3000"
	}
	return "http://localhost:5173"
}
