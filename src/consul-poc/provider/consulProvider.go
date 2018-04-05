package provider

import (
	"github.com/hashicorp/consul/api"
)

type ConsulProvider struct {
	client *api.Client
}

// GetConsulAgent - Return Consul Agent
func (provider *ConsulProvider) GetConsulAgent() *api.Agent {
	return provider.client.Agent()
}

// GetConsulKV - Return Consul KV
func (provider *ConsulProvider) GetConsulKV() *api.KV {
	return provider.client.KV()
}

// NewConsulProvider - return a newly created structure
func NewConsulProvider() (*ConsulProvider, error) {
	// Default consul configuration address : "http://127.0.0.1:8500"
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		return nil, err
	}
	return &ConsulProvider{client}, nil
}
