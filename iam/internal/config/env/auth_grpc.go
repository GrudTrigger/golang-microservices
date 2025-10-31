package env

import (
	"net"

	"github.com/caarlos0/env/v11"
)

type authGRPCEnvConfig struct {
	Host string `env:"GRPC_HOST,required"`
	Port string `env:"GRPC_PORT,required"`
}

type authGRPCConfig struct {
	row authGRPCEnvConfig
}

func NewAuthGRPCConfig() (*authGRPCConfig, error) {
	var row authGRPCEnvConfig
	if err := env.Parse(&row); err != nil {
		return nil, err
	}
	return &authGRPCConfig{
		row: row,
	}, nil
}

func (cfg *authGRPCConfig) Address() string {
	return net.JoinHostPort(cfg.row.Host, cfg.row.Port)
}
