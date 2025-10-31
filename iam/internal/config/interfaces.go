package config

type LoggerConfig interface {
	Level() string
	AsJson() bool
}

type IAMGRPCConfig interface {
	Address() string
}

type PostgresConfig interface {
	URI() string
	MigrationDir() string
}

type RedisConfig interface {
}
