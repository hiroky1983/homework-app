package config

import (
	"context"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Env          string `envconfig:"ENVIRONMENT" required:"true"`
	AppName      string `envconfig:"APP_NAME" required:"true"`
	AppPort      string `envconfig:"APP_PORT" required:"true" default:"8080"`
	AppURL       string `envconfig:"APP_URL" required:"true"`
	PostgresUser  string `envconfig:"POSTGRES_USER" required:"true"`
	PostgresPW	 string `envconfig:"POSTGRES_PW" required:"true"`
	PostgresDB   string `envconfig:"POSTGRES_DB" required:"true"`
	PostgresPort string `envconfig:"POSTGRES_PORT" required:"true"`
	PostgresHost string `envconfig:"POSTGRES_HOST" required:"true"`
	Seclet 			 string `envconfig:"SECRET" required:"true"`
	APIDomain    string `envconfig:"API_DOMAIN" required:"true"`
}

func NewConfig(ctx context.Context) (*Config, error) {
	cfg := &Config{}
	if err := envconfig.Process("", cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
