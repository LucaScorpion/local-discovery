package server

import (
	"net/http"
	"strings"
)

func getRemoteIp(request *http.Request) string {
	remoteIp := request.RemoteAddr

	// Check if the address is an IPv6 address with port, e.g.: "[::1]:4000"
	if remoteIp[0] == '[' {
		endIpIndex := strings.IndexRune(remoteIp, ']')
		remoteIp = remoteIp[1:endIpIndex]
	}

	// Check if the address is an IPv4 address with port, e.g.: "127.0.0.1:4000"
	if parts := strings.SplitN(remoteIp, ":", 3); len(parts) == 2 {
		remoteIp = parts[0]
	}

	// Consolidate localhost addresses.
	if remoteIp == "::1" {
		remoteIp = "127.0.0.1"
	}

	return remoteIp
}
