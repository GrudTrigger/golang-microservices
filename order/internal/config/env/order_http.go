package env

import (
	"net"

	"github.com/caarlos0/env/v11"
)

type orderHttpEnvConfig struct {
	Host          string `env:"HTTP_HOST,required"`
	Port          string `env:"HTTP_PORT,required"`
	HttpTimeout   string `env:"HTTP_READ_TIMEOUT,required"`
	InventoryHost string `env:"INVENTORY_GRPC_HOST,required"`
	InventoryPort string `env:"INVENTORY_GRPC_PORT,required"`
	PaymentHost   string `env:"PAYMENT_GRPC_HOST,required"`
	PaymentPort   string `env:"PAYMENT_GRPC_PORT,required"`
}

type orderHttpConfig struct {
	row orderHttpEnvConfig
}

func NewOrderHttpConfig() (*orderHttpConfig, error) {
	var row orderHttpEnvConfig
	if err := env.Parse(&row); err != nil {
		return nil, err
	}
	return &orderHttpConfig{row: row}, nil
}

func (o *orderHttpConfig) InventoryClientAddress() string {
	return net.JoinHostPort(o.row.InventoryHost, o.row.InventoryPort)
}

func (o *orderHttpConfig) PaymentClientAddress() string {
	return net.JoinHostPort(o.row.PaymentHost, o.row.PaymentPort)
}

func (o *orderHttpConfig) Address() string {
	return net.JoinHostPort(o.row.Host, o.row.Port)
}

func (o *orderHttpConfig) TimeOut() string {
	return o.row.HttpTimeout
}
