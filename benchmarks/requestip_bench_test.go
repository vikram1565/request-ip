package benchmarks

import (
	"log"
	"net/http"
	rip "request-ip"
	"testing"
)

func BenchmarkGetClientIP(b *testing.B) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Your client ip is ", rip.GetClientIP(r))
	})
	go func() {
		http.ListenAndServe(":9000", nil)
	}()
	client := &http.Client{}
	for n := 0; n < 5; n++ {
		client.Get("http://localhost:9000/")
	}
}
