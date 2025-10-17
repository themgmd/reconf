package reconf

import "context"

type configCtxKey struct{}

var configCtxKeyValue configCtxKey

// WithContext обернуть конфиг в контекст
func WithContext(ctx context.Context, cfg Client) context.Context {
	return context.WithValue(ctx, configCtxKeyValue, cfg)
}

// FromContext достать конфиг из контекста
// если в конфиге нет контекста создаем новый и возвращаем
func FromContext(ctx context.Context) (Client, error) {
	config := ctx.Value(configCtxKeyValue)
	if config == nil {
		return NewClient()
	}

	return config.(Client), nil
}
