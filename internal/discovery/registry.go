package discovery

import (
	"errors"
	"time"
)

type Registry struct {
	// TODO: add mutex lock
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

	// Remove agents where the name and local address are the same as the new agent.
	reg.RemoveAgent(publicIp, agent)

	agent.registered = time.Now()
	reg.agents[publicIp] = append(reg.agents[publicIp], &agent)
	return &agent, nil
}

func (reg *Registry) RemoveAgent(publicIp string, agent Agent) {
	newAgents := make([]*Agent, 0)
	for _, check := range reg.agents[publicIp] {
		if !IsSameAgent(agent, *check) {
			newAgents = append(newAgents, check)
		}
	}
	reg.agents[publicIp] = newAgents
}
