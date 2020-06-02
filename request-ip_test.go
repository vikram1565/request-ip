package ip

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetClientIP(t *testing.T) {
	tests := []struct {
		headerName  string
		want        string
		headerKey   string
		headerValue string
	}{
		{"localhost", "127.0.0.1", "", ""},
		{"x-client-ip", "172.20.10.4", "x-client-ip", "172.20.10.4"},
		{"x-forwarded-for", "172.20.10.4", "x-forwarded-for", "172.20.10.4:8080,182.4.10.4, 180.2.10.10"},
		{"cf-connecting-ip", "172.20.10.4", "cf-connecting-ip", "172.20.10.4"},
		{"fastly-client-ip", "172.20.10.4", "fastly-client-ip", "172.20.10.4"},
		{"true-client-ip", "172.20.10.4", "true-client-ip", "172.20.10.4"},
		{"x-real-ip", "172.20.10.4", "x-real-ip", "172.20.10.4"},
		{"x-cluster-client-ip", "172.20.10.4", "x-cluster-client-ip", "172.20.10.4"},
		{"x-forwarded", "172.20.10.4", "x-forwarded", "172.20.10.4"},
		{"forwarded-for", "172.20.10.4", "forwarded-for", "172.20.10.4"},
		{"forwarded", "172.20.10.4", "forwarded", "172.20.10.4"},
	}
	for _, tt := range tests {
		t.Run(tt.headerName, func(t *testing.T) {
			var got string
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if tt.headerKey != "" {
					r.Header.Set(tt.headerKey, tt.headerValue)
				}
				got = GetClientIP(r)
				t.Log("Client IP :", got)
			}))
			defer ts.Close()

			http.Get(ts.URL)
			if got != tt.want {
				t.Fatalf("GetClientIP(%v) => got %v, want %v ", tt.headerName, got, tt.want)
			}
		})
	}
}

func BenchGetClientIP(b *testing.B) {
	r := *http.Request{}
	for n := 0; n < b.N; n++ {
		sendRequest(client, "http://127.0.0.1:8080/")
	}
}
