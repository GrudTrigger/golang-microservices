package env

import (
	"fmt"

	"github.com/caarlos0/env/v11"
)

type postgresEnvConfig struct {
	Host         string `env:"POSTGRES_HOST,required"`
	Port         string `env:"POSTGRES_PORT,required"`
	User         string `env:"POSTGRES_USER,required"`
	Password     string `env:"POSTGRES_PASSWORD,required"`
	Db           string `env:"POSTGRES_DB,required"`
	SslMode      string `env:"POSTGRES_SSL_MODE,required"`
	MigrationDir string `env:"MIGRATION_DIRECTORY,required"`
}

type postgresConfig struct {
	row postgresEnvConfig
}

func NewPostgresConfig() (*postgresConfig, error) {
	var row postgresEnvConfig
	if err := env.Parse(&row); err != nil {
		return nil, err
	}
	return &postgresConfig{row: row}, nil
}

func (p *postgresConfig) URI() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		p.row.User,
		p.row.Password,
		p.row.Host,
		p.row.Port,
		p.row.Db,
	)
}

func (p *postgresConfig) MigrationDir() string {
	return p.row.MigrationDir
}
