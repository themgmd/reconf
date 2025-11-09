package reconf

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"

	"github.com/themgmd/reconf/internal/constants"
)

// Client .
type Client interface {
	GetValue(ctx context.Context, name string) Valuer
}

// ConfigClient .
type ConfigClient struct {
	config map[string]Value
	secret map[string]string

	secretClient Secret
}

// GetValue - receive configs variable
func (c *ConfigClient) GetValue(ctx context.Context, name string) Valuer {
	// at first look variable in configs map
	value := c.getConfigValue(name)
	if value != nil {
		return value
	}

	// if variable not found in configs map
	// look at secret map
	value = c.getSecretValue(ctx, name)
	if value != nil {
		return value
	}

	// if not found return nil value
	return &Value{}
}

func (c *ConfigClient) getConfigValue(name string) Valuer {
	configValue, ok := c.config[name]
	if !ok {
		return nil
	}

	return &configValue
}

func (c *ConfigClient) getSecretValue(ctx context.Context, name string) Valuer {
	secretKey, ok := c.secret[name]
	if !ok {
		return nil
	}

	// if secret client not initialized search variable in environment
	if c.secretClient == nil {
		value := os.Getenv(strings.ToUpper(secretKey))
		return &Value{
			values: []any{value},
		}
	}

	secretValue, err := c.secretClient.GetValue(ctx, secretKey)
	if err != nil {
		slog.Error("error getting secret value from secret client",
			"secret_key", secretKey,
			"error", err.Error(),
		)
	}

	return &Value{
		values: []any{secretValue},
	}
}

// SetSecretClient устанавливает клиент секретов
func (c *ConfigClient) SetSecretClient(secret Secret) {
	c.secretClient = secret
}

type configClient struct {
	Config map[string]Value  `json:"configs"`
	Secret map[string]string `json:"secret"`
}

// NewClient new configs client
func NewClient() (Client, error) {
	cfg := &ConfigClient{}

	workDirectory, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	// достаем переменные окружения
	env := os.Getenv(constants.AppEnvKey)
	configDir := os.Getenv(constants.LocalConfigKey)
	if configDir == "" {
		configDir = filepath.Join(workDirectory, "build", "configs")
	}

	// вычитываем дефолтную конфигурацию
	filename := fmt.Sprintf("%s/values.yaml", configDir)

	values, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("failed to read yaml: %v", err)
	}

	var content configClient
	if err = yaml.Unmarshal(values, &content); err != nil {
		log.Fatalf("failed to unmarshal yaml: %v", err)
	}

	// устанавливает значения дефолтного конфига
	cfg.config = content.Config
	cfg.secret = content.Secret

	// если не задано окружение возвращаем конфиг
	// с дефолтным набором переменных из главного файла
	if env == "" {
		return cfg, nil
	}

	// вычитываем конфигурацию для текущего окружения
	filename = fmt.Sprintf("%s/values_%s.yaml", configDir, strings.ToLower(env))

	envValues, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("failed to read yaml: %v", err)
	}

	if err = yaml.Unmarshal(envValues, &content); err != nil {
		log.Fatalf("failed to unmarshal yaml: %v", err)
	}

	// перезаписываем измененные данные в конфиге
	for k, v := range content.Config {
		cfg.config[k] = v
	}

	for k, v := range content.Secret {
		cfg.secret[k] = v
	}

	return cfg, nil
}
