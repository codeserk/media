package controller

import (
	"net"
	"net/http"
	"strings"
)

func ReadUserIP(r *http.Request) string {
	ipAddress := r.Header.Get("X-Real-Ip")
	if ipAddress == "" {
		ipAddress = r.Header.Get("X-Forwarded-For")
	}
	if ipAddress == "" {
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			return strings.Split(r.RemoteAddr, ":")[0]
		}
		ipAddress = ip
	}
	if ipAddress == "" || ipAddress == "::1" {
		return "127.0.0.1"
	}

	return ipAddress
}
