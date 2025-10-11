package config

type LoggerConfig interface {
	Level() string
	AsJson() bool
}

type PostgresConfig interface {
	URI() string
	MigrationDir() string
}

type OrderHttpConfig interface {
	InventoryClientAddress() string
	PaymentClientAddress() string
	Address() string
	TimeOut() string
}
