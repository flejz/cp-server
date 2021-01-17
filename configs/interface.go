package configs

type Config interface {
	Load() error
}

func Load(configs []*Config) error {

	for _, c := range configs {
		if err := &c.Load(); err != nil {
			return err
		}
	}

	return nil
}
