package reconf

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestContext_WithContext(t *testing.T) {
	ctx := context.Background()
	cfg := &ConfigClient{}

	ctxWithCfg := WithContext(ctx, cfg)
	require.NotNil(t, ctxWithCfg.Value(configCtxKeyValue))
}

func TestContext_FromContext(t *testing.T) {
	cfg := &ConfigClient{}

	t.Run("config in context", func(t *testing.T) {
		ctx := context.Background()
		ctxWithCfg := WithContext(ctx, cfg)

		cfgFromCtx, err := FromContext(ctxWithCfg)
		require.NoError(t, err)
		require.NotNil(t, cfgFromCtx)
		require.Equal(t, cfg, cfgFromCtx)
	})

}
