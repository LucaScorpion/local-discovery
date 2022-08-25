package discovery

import (
	"errors"
	"sync"
	"time"
)

type Registry struct {
	agents     map[string][]*Agent
	agentsLock sync.Mutex
}

func NewRegistry() *Registry {
	return &Registry{
		agents: map[string][]*Agent{},
	}
}

func (reg *Registry) GetAgents(publicIp string) []*Agent {
	reg.agentsLock.Lock()
	defer reg.agentsLock.Unlock()

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

	reg.agentsLock.Lock()
	defer reg.agentsLock.Unlock()

	// Remove agents where the name and local address are the same as the new agent.
	reg.doRemoveAgent(publicIp, agent)

	agent.registered = time.Now()
	reg.agents[publicIp] = append(reg.agents[publicIp], &agent)
	return &agent, nil
}

func (reg *Registry) RemoveAgent(publicIp string, agent Agent) {
	reg.agentsLock.Lock()
	defer reg.agentsLock.Unlock()

	reg.doRemoveAgent(publicIp, agent)
}

func (reg *Registry) doRemoveAgent(publicIp string, agent Agent) {
	// Here we already have a lock.

	newAgents := make([]*Agent, 0)
	for _, check := range reg.agents[publicIp] {
		if !IsSameAgent(agent, *check) {
			newAgents = append(newAgents, check)
		}
	}
	reg.agents[publicIp] = newAgents
}
