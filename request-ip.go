package ip

import (
	"net"
	"net/http"
	"strings"
)

// Standard headers list
var requestHeaders = []string{"X-Client-Ip", "X-Forwarded-For", "Cf-Connecting-Ip", "Fastly-Client-Ip", "True-Client-Ip", "X-Real-Ip", "X-Cluster-Client-Ip", "X-Forwarded", "Forwarded-For", "Forwarded"}

// GetClientIP - returns IP address string; The IP address if known, defaulting to empty string if unknown.
func GetClientIP(r *http.Request) string {

	for _, header := range requestHeaders {
		switch header {
		case "X-Forwarded-For": // Load-balancers (AWS ELB) or proxies.
			if host, correctIP := getClientIPFromXForwardedFor(r.Header.Get(header)); correctIP {
				return host
			}
		default:
			if host := r.Header.Get(header); isCorrectIP(host) {
				return host
			}
		}
	}

	//  remote address checks.
	host, _, splitHostPortError := net.SplitHostPort(r.RemoteAddr)
	if splitHostPortError == nil && isCorrectIP(host) {
		return host
	}
	return ""
}

// getClientIPFromXForwardedFor  - returns first known ip address else return empty string
func getClientIPFromXForwardedFor(headers string) (string, bool) {
	if headers == "" {
		return "", false
	}
	// x-forwarded-for may return multiple IP addresses in the format: "client IP, proxy 1 IP, proxy 2 IP"
	// Therefore, the right-most IP address is the IP address of the most recent proxy
	// and the left-most IP address is the IP address of the originating client.
	forwardedIps := strings.Split(headers, ",")
	for _, ip := range forwardedIps {
		// header can contain spaces too, strip those out.
		ip = strings.TrimSpace(ip)
		// make sure we only use this if it's ipv4 (ip:port)
		if splitted := strings.Split(ip, ":"); len(splitted) == 2 {
			ip = splitted[0]
		}
		if isCorrectIP(ip) {
			return ip, true
		}
	}
	return "", false
}

// isCorrectIP - return true if ip string is valid textual representation of an IP address, else returns false
func isCorrectIP(ip string) bool {
	return net.ParseIP(ip) != nil
}
