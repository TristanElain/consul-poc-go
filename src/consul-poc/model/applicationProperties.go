package model

// ApplicationProperties - Represents the configuration stored in Consul KV Storage
type ApplicationProperties struct {
	Version    float64 `json:"version"`
	ArtifactID string  `json:"artifactId"`
}

// NewEmptyApplicationProperties - Create an empty ApplicationProperties structure
func NewEmptyApplicationProperties() ApplicationProperties {
	return ApplicationProperties{}
}
