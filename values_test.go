package reconf

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

func TestValue_UnmarshalYAML(t *testing.T) {
	yamlContent := `
key: value
key2: [value1, value2]
key3:
  - val1
  - val2
`

	t.Run("success", func(t *testing.T) {
		data := make(map[string]Value)

		err := yaml.Unmarshal([]byte(yamlContent), data)
		require.NoError(t, err)

		keyValues := []any{"value"}
		key2Values := []any{"value1", "value2"}
		key3Values := []any{"val1", "val2"}

		require.Equal(t, data["key"].values, keyValues)
		require.Equal(t, data["key2"].values, key2Values)
		require.Equal(t, data["key3"].values, key3Values)
	})
}

func TestValue_IsNil(t *testing.T) {
	tests := []struct {
		name     string
		data     Value
		expected bool
	}{
		{
			name:     "nil",
			data:     Value{},
			expected: true,
		},
		{
			name: "empty",
			data: Value{
				values: []any{},
			},
			expected: true,
		},
		{
			name: "nil first value",
			data: Value{
				values: []any{nil},
			},
			expected: true,
		},
		{
			name: "not nil",
			data: Value{
				values: []any{"value"},
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			isNil := tt.data.IsNil()
			require.Equal(t, tt.expected, isNil)
		})
	}
}
