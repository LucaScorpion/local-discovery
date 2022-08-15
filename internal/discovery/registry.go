package discovery

import (
	"errors"
	"time"
)

type registry struct {
	agents map[string][]*Agent
}

func newRegistry() *registry {
	return &registry{}
}

func (reg *registry) getAgents(publicIp string) []*Agent {
	return reg.agents[publicIp]
}

func (reg *registry) registerAgent(publicIp string, agent Agent) error {
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
