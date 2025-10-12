package mock

import "context"

// SecretClient mock for secret client
type SecretClient struct {
	config map[string]string
}

// SetValue set value in mock secret client
func (c *SecretClient) SetValue(key string, value string) {
	c.config[key] = value
}

// GetValue get value from mock secret client
func (c *SecretClient) GetValue(_ context.Context, key string) (string, error) {
	return c.config[key], nil
}
