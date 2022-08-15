package server

import (
	"net/http"
	"strings"
)

func getRemoteIp(request *http.Request) string {
	remoteAddr := request.RemoteAddr

	// Check if the address is an IPv6 address with port, e.g.: "[::1]:4000"
	if remoteAddr[0] == '[' {
		endIpIndex := strings.IndexRune(remoteAddr, ']')
		return remoteAddr[1:endIpIndex]
	}

	// Check if the address is an IPv4 address with port, e.g.: "127.0.0.1:4000"
	if parts := strings.SplitN(remoteAddr, ":", 3); len(parts) == 2 {
		return parts[0]
	}

	return remoteAddr
}
