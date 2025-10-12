package reconf

import "context"

// Secret хранилище секретов
type Secret interface {
	GetValue(ctx context.Context, key string) (string, error)
}
