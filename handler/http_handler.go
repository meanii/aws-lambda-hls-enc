package handler

import (
	"net/http"
)

// HandlerMethodOptions is the handler for OPTIONS method
func HandlerMethodOptions(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.WriteHeader(http.StatusOK)
	}
}

// HandlerMethodHead is the handler for HEAD method
func HandlerMethodHead(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodHead {
		w.WriteHeader(http.StatusOK)
	}
}
