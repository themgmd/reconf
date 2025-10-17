package reconf

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

type testSecretMock struct{}

// GetValue get value from mock secret client
func (c *testSecretMock) GetValue(context.Context, string) (string, error) {
	return "value", nil
}

func TestClient_GetValue(t *testing.T) {
	ctx := context.Background()

	client := &ConfigClient{
		config:       make(map[string]Value),
		secret:       make(map[string]string),
		secretClient: &testSecretMock{},
	}

	t.Run("GetValue from config", func(t *testing.T) {
		client.config = map[string]Value{
			"test": {
				values: []any{"value"},
			},
		}

		value := client.GetValue(ctx, "test")
		require.Equal(t, value.String(), "value")
	})

	t.Run("GetValue from secret", func(t *testing.T) {
		client.secret = map[string]string{
			"vault_key": "value",
		}

		value := client.GetValue(ctx, "vault_key")
		require.Equal(t, value.String(), "value")
	})
}

func TestClient_NewClient(t *testing.T) {
	client, err := NewClient()
	require.NoError(t, err)
	require.NotNil(t, client)
}
