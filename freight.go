package freight

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
)

type Rate struct {
	Provider  string
	ApiKey    string
	ApiSecret string
	Config    map[string]string
	Instance  interface{}
	Logger    zerolog.Logger
}

type FrightRate struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Handle      string  `json:"handle"`
	Price       float64 `json:"price"`
	Currency    string  `json:"currency"`
}

func NewRate(provider string, apiKey string, apiSecret string) *Rate {
	return &Rate{
		Provider:  provider,
		ApiKey:    apiKey,
		ApiSecret: apiSecret,
	}
}

func (r *Rate) SetLogger(logger *zerolog.Logger) {

	if logger == nil {
		r.Logger = zerolog.New(os.Stdout)
		return
	}

	r.Logger = *logger

}

func (r *Rate) SetConfig(config map[string]string) error {

	if len(config) == 0 {
		r.Logger.Info().Msg("Config cannot be empty")
		return fmt.Errorf("config cannot be empty") // or handle the error as needed
	}

	for key, value := range config {
		if len(key) > 98 || len(value) > 98 {
			r.Logger.Info().Msg("Config key or value exceeds 98 characters")
			return fmt.Errorf("config key or value exceeds 98 characters")
		}
	}

	r.Config = config

	return nil
}

type Provider interface {
	SetLogger(logger *zerolog.Logger)
	SetConfig(config map[string]string) error
	Calculate(details DeliveryDetails) ([]FrightRate, error)
}
