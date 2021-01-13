package cache

type Cache interface {
	Init() error
	Get(key string) (string, error)
	Set(key string, value string) error
}
