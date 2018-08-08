package config

type Config struct {
	Database DatabaseConfig
}

type DatabaseConfig struct {
	Address string
}
