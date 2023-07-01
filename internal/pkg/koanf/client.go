package koanf

import (
	"fmt"
	"log"
	"strings"

	"github.com/knadh/koanf/parsers/json"
	envprovider "github.com/knadh/koanf/providers/env"
	fileprovider "github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

// WithProviders adds config provider
type WithProviders func(*koanf.Koanf)

// NewClient creates a new, pre-configured Koanf client for working with configs
func NewClient(providers ...WithProviders) *koanf.Koanf {
	client := koanf.New(".")
	for _, apply := range providers {
		apply(client)
	}

	return client
}

// WithEnvProvider includes configs from envs
func WithEnvProvider(prefix string) func(*koanf.Koanf) {
	return func(client *koanf.Koanf) {
		provider := envprovider.Provider(prefix, "_", func(s string) string {
			return strings.ToLower(strings.TrimPrefix(s, fmt.Sprintf("%s_", prefix)))
		})

		if err := client.Load(provider, nil); err != nil {
			log.Fatalf("an error occurred when loading the application configuration from the provider %T: %s", provider, err)
		}
	}
}

// WithFileProvider include configs from files
func WithFileProvider(path string) func(*koanf.Koanf) {
	return func(client *koanf.Koanf) {
		provider := fileprovider.Provider(path)

		if err := client.Load(provider, json.Parser()); err != nil {
			log.Fatalf("an error occurred when loading the application configuration from the provider %T: %s", provider, err)
		}
	}
}
