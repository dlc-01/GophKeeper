// Package config содержит конфигурацию клиента.
package config

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/caarlos0/env/v10"
	"github.com/dlc-01/GophKeeper/internal/general/logger"
	"os"
)

type Config struct {
	App struct {
		Name    string `env:"APP_NAME" envDefault:"clientTUI" json:"Name"`
		Version string `env:"APP_VERSION" envDefault:"0.1" json:"Version"`
	}
	Config     string `env:"CONF" envDefault:""`
	GRPCClient struct {
		Address string `env:"GRPC_CLIENT_ADDRESS" envDefault:"localhost" json:"GRPCClient" json:"GRPCAddress"`
		Port    string `env:"GRPC_SERVER_NETWORK" envDefault:"3200" json:"GRPCNetwork" json:"GRPCPort"`
	}
	Logger logger.ConfLogger
}

func New() (*Config, error) {
	var err error
	cfg := &Config{}
	flag.StringVar(&cfg.Config, "c", "", "path to config in json")
	flag.Parse()

	cfg, err = initConfigFormENV()
	if err != nil {
		return nil, fmt.Errorf(" error while initing while from env: %w", err)
	}

	if cfg.Config != "" {
		cfg, err = initConfigFormJSON(cfg.Config)
		if err != nil {
			return nil, fmt.Errorf(" error while initing while from json: %w", err)
		}
	}

	return cfg, nil
}

func initConfigFormENV() (*Config, error) {
	var cfg Config
	err := env.Parse(&cfg)
	if err != nil {
		return nil, fmt.Errorf("error while parsing env: %w", err)
	}

	return &cfg, nil
}

func initConfigFormJSON(path string) (*Config, error) {
	var cfg Config

	raw, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error while oppening json conf: %w", err)
	}

	err = json.Unmarshal([]byte(raw), &cfg)
	if err != nil {
		return nil, fmt.Errorf("error while unmarshaling json conf: %w", err)
	}

	return &cfg, nil
}
