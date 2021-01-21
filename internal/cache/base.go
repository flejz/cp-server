package cache

type CacheInterface interface {
	Init() error
	Get(usr, key string) (string, error)
	Set(usr, key, value string) error
}

func Init(caches []CacheInterface) error {
	for _, cache := range caches {
		if err := cache.Init(); err != nil {
			return err
		}
	}

	return nil
}

type BaseCache struct {
	DefaultKey string
}

func (cache *BaseCache) key(key string) string {
	if key == "" {
		return cache.DefaultKey
	} else {
		return key
	}
}
func (cache *BaseCache) Get(usr, key string) (string, error) {
	return "", nil
}

func (cache *BaseCache) Set(usr, key, value string) error {
	return nil
}
