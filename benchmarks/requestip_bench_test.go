package benchmarks

import (
	"net/http"
	rip "request-ip"
	"testing"
)

func startHTTPServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		rip.GetClientIP(r)
	})
	go func() {
		http.ListenAndServe(":9999", mux)
	}()
}
func BenchmarkGetClientIP(b *testing.B) {
	startHTTPServer()
	client := &http.Client{}
	for n := 0; n < 1000; n++ {
		client.Get("http://localhost:9999/")
	}
}
