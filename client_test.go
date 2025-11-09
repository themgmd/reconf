package reconf

import (
	"context"
	"github.com/themgmd/reconf/internal/constants"
	"os"
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

	t.Run("GetValue from configs", func(t *testing.T) {
		client.config = map[string]Value{
			"test": {
				values: []any{"value"},
			},
		}

		value := client.GetValue(ctx, "test")
		require.Equal(t, "value", value.String())
	})

	t.Run("GetValue from secret", func(t *testing.T) {
		client.secret = map[string]string{
			"vault_key": "value",
		}

		value := client.GetValue(ctx, "vault_key")
		require.Equal(t, "value", value.String())
	})

	t.Run("secretClient is nil", func(t *testing.T) {
		client.secretClient = nil
		client.secret = map[string]string{
			"vault_key": "value",
		}

		err := os.Setenv("VALUE", "value")
		require.NoError(t, err)

		value := client.GetValue(ctx, "vault_key")
		require.Equal(t, "value", value.String())
	})
}

func TestClient_NewClient(t *testing.T) {
	client, err := NewClient()
	require.NoError(t, err)
	require.NotNil(t, client)

	val := client.GetValue(context.Background(), "key").String()
	require.Equal(t, "value", val)
}

func TestClient_NewClientWithEnv(t *testing.T) {
	err := os.Setenv(constants.AppEnvKey, constants.EnvDev)
	require.NoError(t, err)

	client, err := NewClient()
	require.NoError(t, err)
	require.NotNil(t, client)

	val := client.GetValue(context.Background(), "http_port").String()
	require.Equal(t, "1010", val)
}
