package mock

import "github.com/themgmd/reconf"

// ConfigClient mock configs client
type ConfigClient struct {
	config map[string]reconf.Value
}

// SetValue set value in mock configs client
func (c *ConfigClient) SetValue(key string, value reconf.Value) {
	c.config[key] = value
}

// GetValue get value from mock configs client
func (c *ConfigClient) GetValue(key string) reconf.Value {
	return c.config[key]
}
