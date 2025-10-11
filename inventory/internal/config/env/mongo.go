package env

import (
	"fmt"

	"github.com/caarlos0/env/v11"
)

type mongoEnvConfig struct {
	Host     string `env:"MONGO_HOST,required"`
	Port     string `env:"MONGO_PORT,required"`
	Database string `env:"MONGO_DATABASE,required"`
	User     string `env:"MONGO_INITDB_ROOT_USERNAME,required"`
	Password string `env:"MONGO_INITDB_ROOT_PASSWORD,required"`
	AuthDB   string `env:"MONGO_AUTH_DB,required"`
}

type mongoConfig struct {
	row mongoEnvConfig
}

func NewMongoConfig() (*mongoConfig, error) {
	var row mongoEnvConfig
	if err := env.Parse(&row); err != nil {
		return nil, err
	}
	return &mongoConfig{row: row}, nil
}

func (m *mongoConfig) URI() string {
	return fmt.Sprintf("mongodb://%s:%s@%s:%s/%s?authSource=%s",
		m.row.User,
		m.row.Password,
		m.row.Host,
		m.row.Port,
		m.row.Database,
		m.row.AuthDB,
	)
}

func (m *mongoConfig) DatabaseName() string {
	return m.row.Database
}
