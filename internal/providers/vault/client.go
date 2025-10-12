package vault

import (
	"context"
	"fmt"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"

	"github.com/themgmd/reconf/internal/constants"
)

// Secret клиент для работы с секретами от волта
type Secret interface {
	KvV2Read(ctx context.Context, path string, options ...vault.RequestOption) (*vault.Response[schema.KvV2ReadResponse], error)
}

// Client обертка над вольтом для получения секретов
type Client struct {
	secret Secret
}

// NewClient конструктор, на вход принимает клиент для запросов к секретам vault.Secrets
func NewClient(secret Secret) *Client {
	return &Client{secret: secret}
}

// GetValue получить клиент по ключу
func (c *Client) GetValue(ctx context.Context, key string) (string, error) {
	env := os.Getenv(constants.AppEnvKey)
	if env == "" {
		env = constants.EnvDev
	}

	vaultpath := fmt.Sprintf("secret/data/%s", env)

	// read the secret
	s, err := c.secret.KvV2Read(ctx, key, vault.WithMountPath(vaultpath))
	if err != nil {
		return "", err
	}

	return s.Data.Data[key].(string), nil
}
