package server

import (
	"local-discovery/internal/discovery"
	"log"
	"net/http"
	"strings"
)

func StartServer() {
	reg := discovery.NewRegistry()
	mux := http.NewServeMux()

	mux.HandleFunc("/api/agents", agents(reg))

	if err := http.ListenAndServe(":4000", mux); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

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

type jsonErr struct {
	Error string `json:"error"`
}
