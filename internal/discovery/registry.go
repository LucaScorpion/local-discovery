package discovery

import (
	"errors"
	"time"
)

type Registry struct {
	agents map[string][]*Agent
}

func NewRegistry() *Registry {
	return &Registry{}
}

func (reg *Registry) GetAgents(publicIp string) []*Agent {
	return reg.agents[publicIp]
}

func (reg *Registry) registerAgent(publicIp string, agent Agent) error {
	// Validate the agent info.
	if len(agent.Name) == 0 {
		return errors.New("name should not be empty")
	}
	if len(agent.LocalAddress) == 0 {
		return errors.New("local address should not be empty")
	}

	agent.registered = time.Now()
	reg.agents[publicIp] = append(reg.agents[publicIp], &agent)
	return nil
}
