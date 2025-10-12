package reconf

// Secret хранилище секретов
type Secret interface {
	GetValue(key string) string
}
