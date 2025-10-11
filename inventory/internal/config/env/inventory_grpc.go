package env

import (
	"net"

	"github.com/caarlos0/env/v11"
)

type inventoryGRPCEnvConfig struct {
	Host string `env:"GRPC_HOST,required"`
	Port string `env:"GRPC_PORT,required"`
}

type inventoryGRPCConfig struct {
	row inventoryGRPCEnvConfig
}

func NewInventoryGRPCConfig() (*inventoryGRPCConfig, error) {
	var row inventoryGRPCEnvConfig
	if err := env.Parse(&row); err != nil {
		return nil, err
	}
	return &inventoryGRPCConfig{row: row}, nil
}

func (i *inventoryGRPCConfig) Address() string {
	return net.JoinHostPort(i.row.Host, i.row.Port)
}
