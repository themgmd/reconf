package reconf

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestClient_GetValue(t *testing.T) {
	client := &Client{
		config: make(map[string]Value),
		secret: make(map[string]string),
	}

	t.Run("GetValue from config", func(t *testing.T) {
		client.config = map[string]Value{
			"test": {
				values: []any{"value"},
			},
		}

		value := client.GetValue("test")
		require.Equal(t, value.String(), "value")
	})

	t.Run("GetValue from secret", func(t *testing.T) {
		client.secret = map[string]string{
			"vault_key": "value",
		}

		value := client.GetValue("vault_key")
		require.Equal(t, value.String(), "value")
	})
}
