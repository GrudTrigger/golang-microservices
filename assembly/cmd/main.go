package main

import (
	"fmt"

	"github.com/rocker-crm/assembly/internal/config"
)

const configPath = "../deploy/compose/assembly/.env"

func main() {
	err := config.Load(configPath)
	if err != nil {
		panic(fmt.Errorf("failed to load config: %w", err))
	}
}
