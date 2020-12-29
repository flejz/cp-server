package cache

type Cache interface {
	connect() error
	get(key string) (error, string)
	set(key string, value string) error
}
