package server

import (
	"encoding/json"
	"local-discovery/internal/discovery"
	"net/http"
)

func agents(reg *discovery.Registry) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")

		switch request.Method {
		case http.MethodGet:
			getAgents(reg, writer, request)
		case http.MethodPost:
			postAgent(reg, writer, request)
		default:
			writer.WriteHeader(405)
		}
	}
}

func getAgents(reg *discovery.Registry, writer http.ResponseWriter, request *http.Request) {
	list := reg.GetAgents(getRemoteIp(request))

	// Make sure the list is always a list, not nil.
	if list == nil {
		list = []*discovery.Agent{}
	}

	b, _ := json.Marshal(list)
	writer.Write(b)
}

func postAgent(reg *discovery.Registry, writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Created agent goes here"))
}
