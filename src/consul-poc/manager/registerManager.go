package manager

import (
	"consul-poc/provider"
	"fmt"
	"github.com/hashicorp/consul/api"
	"sync"
	"time"
)

type ConsulServiceManager struct {
	agent *api.Agent
}

func (manager *ConsulServiceManager) RegisterService(name string, port int) error {
	id := fmt.Sprintf("%s-%v", name, time.Now().Unix())

	registered, err := manager.IsServiceNameRegistered(name)
	if err != nil {
		return err
	}

	if !registered {
		fmt.Println("Registering to consul ...")
		register := api.AgentServiceRegistration{
			ID:   id,
			Name: name,
			Port: port,
			Tags: []string{"go"},
			Check: &api.AgentServiceCheck{
				HTTP:     "http://localhost:4000/health",
				Interval: "10s",
			},
		}
		manager.agent.ServiceRegister(&register)
	} else {
		fmt.Println("Service already registered.")
	}
	return nil
}

// DeregisterService - Deregister a service from Consul
func (manager *ConsulServiceManager) DeregisterService(serviceID string) error {
	fmt.Println("Deregistering from consul ...")
	return manager.agent.ServiceDeregister(serviceID)
}

// IsServiceRegistered - Tell if a service id is registered for this agent
func (manager *ConsulServiceManager) IsServiceRegistered(id string) (bool, error) {
	services, err := manager.agent.Services()
	if err != nil {
		return false, err
	}

	_, ok := services[id]
	return ok, nil
}

// IsServiceNameRegistered - Tell if a service name is registered for this agent
func (manager *ConsulServiceManager) IsServiceNameRegistered(name string) (bool, error) {
	services, err := manager.agent.Services()
	if err != nil {
		return false, err
	}

	for _, agentService := range services {
		if name == agentService.Service {
			return true, nil
		}
	}
	return false, nil
}

var consulServiceManager *ConsulServiceManager
var onceServiceManager sync.Once

func GetConsulServiceManager() *ConsulServiceManager {
	onceServiceManager.Do(func() {
		consulProvider := provider.GetConsulProvider()
		agent := consulProvider.GetConsulAgent()
		consulServiceManager = &ConsulServiceManager{agent}
	})
	return consulServiceManager
}
