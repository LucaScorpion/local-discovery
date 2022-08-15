package discovery

import (
	"errors"
	"time"
)

type Registry struct {
	agents map[string][]*Agent
}

func NewRegistry() *Registry {
	return &Registry{
		agents: map[string][]*Agent{},
	}
}

func (reg *Registry) GetAgents(publicIp string) []*Agent {
	return reg.agents[publicIp]
}

func (reg *Registry) RegisterAgent(publicIp string, agent Agent) (*Agent, error) {
	// Validate the agent info.
	if len(agent.Name) == 0 {
		return nil, errors.New("name should not be empty")
	}
	if len(agent.LocalAddress) == 0 {
		return nil, errors.New("local address should not be empty")
	}

	if agent.Info == nil {
		agent.Info = map[string]any{}
	}

	agent.registered = time.Now()
	reg.agents[publicIp] = append(reg.agents[publicIp], &agent)
	return &agent, nil
}
