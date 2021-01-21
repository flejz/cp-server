package db

type DB interface {
	Connect() error
}

func Connect(dbs []DB) error {
	for _, db := range dbs {
		if err := db.Connect(); err != nil {
			return err
		}
	}

	return nil
}
