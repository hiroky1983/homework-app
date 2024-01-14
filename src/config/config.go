package config

import (
	"context"

	"github.com/kelseyhightower/envconfig"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type Config struct {
	Env          string `envconfig:"ENVIRONMENT" required:"true"`
	AppName      string `envconfig:"APP_NAME" required:"true"`
	AppPort      string `envconfig:"APP_PORT" required:"true" default:"8080"`
	AppURL       string `envconfig:"APP_URL" required:"true"`
	PostgresUser string `envconfig:"POSTGRES_USER" required:"true"`
	PostgresPW   string `envconfig:"POSTGRES_PW" required:"true"`
	PostgresDB   string `envconfig:"POSTGRES_DB" required:"true"`
	PostgresPort string `envconfig:"POSTGRES_PORT" required:"true"`
	PostgresHost string `envconfig:"POSTGRES_HOST" required:"true"`
	Seclet       string `envconfig:"SECRET" required:"true"`
	APIDomain    string `envconfig:"API_DOMAIN" required:"true"`
	GoogleAPIKey string `envconfig:"GOOGLE_API_KEY" required:"true"`
	GoogleOAuthClientID string `envconfig:"GOOGLE_OAUTH_CLIENT_ID" required:"true"`
	GoogleOAuthClientSecret string `envconfig:"GOOGLE_OAUTH_CLIENT_SECRET" required:"true"`
	GoogleOAuthRedirectURL string `envconfig:"GOOGLE_OAUTH_REDIRECT_URL" required:"true"`
}

func NewConfig(ctx context.Context) (*Config, error) {
	cfg := &Config{}
	if err := envconfig.Process("", cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

func (c *Config)NewGoogleOauthConfig() *oauth2.Config {
	conf := &oauth2.Config{
		RedirectURL:  c.GoogleOAuthRedirectURL,
		ClientID:     c.GoogleOAuthClientID,
		ClientSecret: c.GoogleOAuthClientSecret,
		Scopes:       []string{"openid", "email", "profile"},
		Endpoint:     google.Endpoint,
	}

	return conf
}