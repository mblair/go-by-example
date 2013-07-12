package main

import "net/http"
import "fmt"

type MattHandler http.HandlerFunc

func hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintln(w, "Hello wrapped world")
}

func wrapMiddleware(f MattHandler) MattHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Wrapped", "true")
		f(w, r)
	}
}

func main() {
	handler := wrapMiddleware(hello)
	http.HandleFunc("/", handler)
	http.ListenAndServe(":4318", nil)
}
