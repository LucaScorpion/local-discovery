package server

import (
	"encoding/json"
	"fmt"
	"local-discovery/internal/discovery"
	"net/http"
)

func agents(reg *discovery.Registry) handlerWithErrorFunc {
	return func(writer http.ResponseWriter, request *http.Request) error {
		writer.Header().Set("Content-Type", "application/json")

		switch request.Method {
		case http.MethodGet:
			return getAgents(reg, writer, request)
		case http.MethodPost:
			return postAgent(reg, writer, request)
		default:
			return newHttpError(http.StatusMethodNotAllowed, "")
		}
	}
}

func getAgents(reg *discovery.Registry, writer http.ResponseWriter, request *http.Request) error {
	list := reg.GetAgents(getRemoteIp(request))

	// Make sure the list is always a list, not nil.
	if list == nil {
		list = []*discovery.Agent{}
	}

	json.NewEncoder(writer).Encode(list)
	return nil
}

func postAgent(reg *discovery.Registry, writer http.ResponseWriter, request *http.Request) error {
	agentDto := &discovery.Agent{}
	enc := json.NewEncoder(writer)

	if err := json.NewDecoder(request.Body).Decode(agentDto); err != nil {
		return newHttpError(http.StatusBadRequest, fmt.Sprintf("Invalid JSON: %v", err))
	}

	agent, err := reg.RegisterAgent(getRemoteIp(request), *agentDto)
	if err != nil {
		return newHttpError(http.StatusBadRequest, fmt.Sprintf("Invalid agent: %v", err))
	}

	writer.WriteHeader(http.StatusCreated)
	enc.Encode(agent)
	return nil
}
