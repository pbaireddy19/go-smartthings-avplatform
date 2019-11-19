package smartthings

import "net/http"

type Config struct {
	HTTPClient  *http.Client
	AccessToken string
}

func NewConfig(accessToken string) *Config {
	return &Config{
		HTTPClient:  http.DefaultClient,
		AccessToken: accessToken,
	}
}
