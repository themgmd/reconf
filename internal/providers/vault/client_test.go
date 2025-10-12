package vault

import (
	"context"
	"errors"
	"testing"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
	"github.com/stretchr/testify/require"
)

type testingSecretsMock struct{}

func (m *testingSecretsMock) KvV2Read(_ context.Context, key string, _ ...vault.RequestOption) (*vault.Response[schema.KvV2ReadResponse], error) {
	return &vault.Response[schema.KvV2ReadResponse]{
		Data: schema.KvV2ReadResponse{
			Data: map[string]interface{}{
				key: "value",
			},
		},
	}, nil
}

type testingSecretsErrMock struct{}

func (m *testingSecretsErrMock) KvV2Read(_ context.Context, key string, _ ...vault.RequestOption) (*vault.Response[schema.KvV2ReadResponse], error) {
	return nil, errors.New("error")
}

func TestClient_GetValue(t *testing.T) {
	ctx := context.Background()
	key := "key"

	t.Run("success", func(t *testing.T) {
		client := NewClient(&testingSecretsMock{})

		secret, err := client.GetValue(ctx, key)
		require.NoError(t, err)

		require.Equal(t, "value", secret)
	})

	t.Run("error", func(t *testing.T) {
		client := NewClient(&testingSecretsErrMock{})

		secret, err := client.GetValue(ctx, "key")
		require.Error(t, err)
		require.Empty(t, secret)
	})
}
