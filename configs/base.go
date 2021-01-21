package configs

type Config interface {
	Load() error
}

func Load(configs []Config) error {
	for _, config := range configs {
		if err := config.Load(); err != nil {
			return err
		}
	}

	return nil
}
