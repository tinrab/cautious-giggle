package config

type Config struct {
	Database PostgresConfig
	Event    NatsConfig
}

type PostgresConfig struct {
	Address string
}

type NatsConfig struct {
	Address string
}
