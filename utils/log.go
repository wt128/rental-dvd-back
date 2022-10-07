package utils

import (
	"fmt"
	"net/http"
)

func Log(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		addr := r.RemoteAddr
		method := r.Method
		path := r.URL.Path
		fmt.Printf("Remote: %s [%s] %s", addr, method, path)
		h.ServeHTTP(w, r)
	})
}
