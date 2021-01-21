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
