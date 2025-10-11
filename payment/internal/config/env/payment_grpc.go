package env

import (
	"net"

	"github.com/caarlos0/env/v11"
)

type paymentEnvConfig struct {
	Host string `env:"GRPC_HOST,required"`
	Port string `env:"GRPC_PORT,required"`
}

type paymentConfig struct {
	row paymentEnvConfig
}

func NewPaymentConfig() (*paymentConfig, error) {
	var row paymentEnvConfig
	if err := env.Parse(&row); err != nil {
		return nil, err
	}
	return &paymentConfig{row: row}, nil
}

func (p *paymentConfig) Address() string {
	return net.JoinHostPort(p.row.Host, p.row.Port)
}
