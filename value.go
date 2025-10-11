package reconf

import (
	"fmt"
	"strconv"

	"gopkg.in/yaml.v3"
)

// Valuer config value
type Valuer interface {
	yaml.Unmarshaler

	IsNil() bool
	String() string
	Int() int
	Int32() int32
	Int64() int64
	Bool() bool
	Float64() float64
	Float32() float32
}

// Value .
type Value struct {
	values []any
}

// UnmarshalYAML Метод для yaml.Unmarshaler
func (v *Value) UnmarshalYAML(value *yaml.Node) error {
	var raw any
	// Распарсить значение узла в any
	if err := value.Decode(&raw); err != nil {
		return err
	}

	// В зависимости от типа raw заполняем values
	switch val := raw.(type) {
	case []any:
		v.values = val
	default:
		// Если одиночное значение, упаковываем в слайс
		v.values = []any{val}
	}
	return nil
}

// Helper: взять первое значение, если есть
func (v *Value) first() any {
	if len(v.values) == 0 {
		return nil
	}
	return v.values[0]
}

func (v *Value) int() int {
	val := v.first()
	if val == nil {
		return 0
	}
	switch i := val.(type) {
	case int:
		return i
	case int32:
		return int(i)
	case int64:
		return int(i)
	case float64:
		return int(i)
	case string:
		n, err := strconv.Atoi(i)
		if err == nil {
			return n
		}
	}
	return 0
}

// IsNil .
func (v *Value) IsNil() bool {
	return v.values == nil || len(v.values) == 0 || v.values[0] == nil
}

// String возвращает значение как string
func (v *Value) String() string {
	val := v.first()
	if val == nil {
		return ""
	}
	switch s := val.(type) {
	case string:
		return s
	case fmt.Stringer:
		return s.String()
	case int:
		return strconv.Itoa(s)
	case int64:
		return strconv.FormatInt(s, 10)
	case float64:
		return strconv.FormatFloat(s, 'f', -1, 64)
	case bool:
		return strconv.FormatBool(s)
	default:
		return fmt.Sprintf("%v", val)
	}
}

// Int возвращает значение как int
func (v *Value) Int() int {
	return v.int()
}

// Int32 возвращает значение как int32
func (v *Value) Int32() int32 {
	return int32(v.int())
}

// Int64 возвращает значение как int64
func (v *Value) Int64() int64 {
	return int64(v.int())
}

// Bool возвращает значение как bool
func (v *Value) Bool() bool {
	val := v.first()
	if val == nil {
		return false
	}
	switch b := val.(type) {
	case bool:
		return b
	case string:
		if b == "true" || b == "1" {
			return true
		}
	case int, uint, int32, int64, float64:
		return b != 0
	}

	return false
}

// Float64 возвращает значение как float64
func (v *Value) Float64() float64 {
	val := v.first()
	if val == nil {
		return 0
	}
	switch f := val.(type) {
	case float64:
		return f
	case float32:
		return float64(f)
	case int:
		return float64(f)
	case int64:
		return float64(f)
	case string:
		num, err := strconv.ParseFloat(f, 64)
		if err == nil {
			return num
		}
	}
	return 0
}

// Float32 возвращает значение как float32
func (v *Value) Float32() float32 {
	val := v.first()
	if val == nil {
		return 0
	}
	switch f := val.(type) {
	case string:
		num, err := strconv.ParseFloat(f, 32)
		if err == nil {
			return float32(num)
		}
	default:
		num := v.Float64()
		return float32(num)
	}

	return 0
}
