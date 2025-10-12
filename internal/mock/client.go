package mock

import "github.com/themgmd/reconf"

// ConfigClient mock config client
type ConfigClient struct {
	config map[string]reconf.Value
}

// SetValue set value in mock config client
func (c *ConfigClient) SetValue(key string, value reconf.Value) {
	c.config[key] = value
}

// GetValue get value from mock config client
func (c *ConfigClient) GetValue(key string) reconf.Value {
	return c.config[key]
}
