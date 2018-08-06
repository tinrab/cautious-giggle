package config

type Config struct {
	StoreRepository StoreRepositoryConfig
}

type StoreRepositoryConfig struct {
	Address string
}
