package reconf

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"

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

// GetValue - receive config variable
func (c *ConfigClient) GetValue(ctx context.Context, name string) Valuer {
	// at first look variable in config map
	value := c.getConfigValue(name)
	if value != nil {
		return value
	}

	// if variable not found in config map
	// if secret client not initialized return nil value and log error
	if len(c.secret) > 0 && c.secretClient == nil {
		slog.Error("secret client is nil")
		return &Value{}
	}

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

// NewClient new config client
func NewClient() Client {
	cfg := &ConfigClient{}

	// достаем переменные окружения
	env := os.Getenv(constants.AppEnvKey)
	configDir := os.Getenv(constants.LocalConfigKey)
	if configDir == "" {
		configDir = "./build/configs"
	}

	// вычитываем дефолтную конфигурацию
	filename := fmt.Sprintf("%s/values.yaml", configDir)

	values, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("failed to read yaml: %v", err)
	}

	var content map[string]interface{}
	if err = yaml.Unmarshal(values, &content); err != nil {
		log.Fatalf("failed to unmarshal yaml: %v", err)
	}

	// устанавливает значения дефолтного конфига
	cfg.config = content["config"].(map[string]Value)
	cfg.secret = content["secret"].(map[string]string)

	// если не задано окружение возвращаем конфиг
	// с дефолтным набором переменных из главного файла
	if env == "" {
		return cfg
	}

	// вычитываем конфигурацию для текущего окружения
	filename = fmt.Sprintf("%s/values_%s.yaml", configDir, env)

	envValues, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("failed to read yaml: %v", err)
	}

	if err = yaml.Unmarshal(envValues, &content); err != nil {
		log.Fatalf("failed to unmarshal yaml: %v", err)
	}

	// перезаписываем измененные данные в конфиге
	for k, v := range content["config"].(map[string]Value) {
		cfg.config[k] = v
	}

	for k, v := range content["secret"].(map[string]string) {
		cfg.secret[k] = v
	}

	return cfg
}
