package provider

import (
	"encoding/json"
	"fmt"
	"github.com/hashicorp/consul/api"
	"sync"
)

// PropertiesProvider - Provide functions to interact with consul Consul
type PropertiesProvider struct {
	kv *api.KV
}

// GetProperties - Populate `propsStruct` with properties stored at `key`
func (provider *PropertiesProvider) GetProperties(key string, queryOptions *api.QueryOptions, propsStruct interface{}) error {
	kvPair, _, err := provider.kv.Get(key, queryOptions)
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
	props := make(map[string]interface{}, 2)

	kvPair, _, err := provider.kv.Get(key, queryOptions)
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

var propertiesProvider *PropertiesProvider
var onceProps sync.Once

// GetPropertiesProvider - return the properties provider
func GetPropertiesProvider() *PropertiesProvider {
	onceProps.Do(func() {
		consulProvider := GetConsulProvider()
		kv := consulProvider.GetConsulKV()
		propertiesProvider = &PropertiesProvider{kv}
	})

	return propertiesProvider
}
