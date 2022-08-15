package server

import (
	"encoding/json"
	"fmt"
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
			writer.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

func getAgents(reg *discovery.Registry, writer http.ResponseWriter, request *http.Request) {
	list := reg.GetAgents(getRemoteIp(request))

	// Make sure the list is always a list, not nil.
	if list == nil {
		list = []*discovery.Agent{}
	}

	json.NewEncoder(writer).Encode(list)
}

func postAgent(reg *discovery.Registry, writer http.ResponseWriter, request *http.Request) {
	agentDto := &discovery.Agent{}
	enc := json.NewEncoder(writer)

	if err := json.NewDecoder(request.Body).Decode(agentDto); err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		enc.Encode(jsonErr{
			Error: fmt.Sprintf("Invalid JSON: %v", err),
		})
		return
	}

	agent, err := reg.RegisterAgent(getRemoteIp(request), *agentDto)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		enc.Encode(jsonErr{
			Error: fmt.Sprintf("Invalid agent: %v", err),
		})
		return
	}

	writer.WriteHeader(http.StatusCreated)
	enc.Encode(agent)
}
