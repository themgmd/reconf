package mock

import "github.com/themgmd/reconf"

// Client mock config client
type Client struct {
	config map[string]reconf.Value
}

// SetValue set value in mock config client
func (c *Client) SetValue(key string, value reconf.Value) {
	c.config[key] = value
}

// GetValue get value from mock config client
func (c *Client) GetValue(key string) reconf.Value {
	return c.config[key]
}
