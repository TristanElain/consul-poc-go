package provider

import (
	"encoding/json"
	"fmt"
	"github.com/hashicorp/consul/api"
)

// PropertiesProvider - Provide functions to interact with consul Consul
type PropertiesProvider struct {
	client *api.Client
}

// GetProperties - Populate `propsStruct` with properties stored at `key`
func (provider *PropertiesProvider) GetProperties(key string, queryOptions *api.QueryOptions, propsStruct interface{}) error {
	kv := provider.client.KV()

	kvPair, _, err := kv.Get(key, queryOptions)
	if err != nil {
		return err
	}
	if kvPair == nil {
		return fmt.Errorf("No value for key '%s'\n", key)
	}
	err = json.Unmarshal(kvPair.Value, propsStruct)
	if err != nil {
		return err
	}
	return nil
}

// GetPropertiesMap - Return properties as a map[string]interface
func (provider *PropertiesProvider) GetPropertiesMap(key string, queryOptions *api.QueryOptions) (map[string]interface{}, error) {
	kv := provider.client.KV()
	props := make(map[string]interface{}, 2)

	kvPair, _, err := kv.Get(key, queryOptions)
	if err != nil {
		return props, err
	}
	if kvPair == nil {
		return props, fmt.Errorf("No value for key '%s'\n", key)
	}
	err = json.Unmarshal(kvPair.Value, &props)
	if err != nil {
		return props, err
	}
	return props, nil
}

// NewPropertiesProvider - return a new created structure
func NewPropertiesProvider() (*PropertiesProvider, error) {
	// Default consul configuration address : "http://127.0.0.1:8500"
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		return nil, err
	}
	return &PropertiesProvider{client}, nil
}
