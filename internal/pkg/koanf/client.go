package koanf

import (
	"fmt"
	"strings"

	"github.com/IASamoylov/tg_calories_observer/internal/pkg/logger"
	"github.com/knadh/koanf/parsers/json"
	envprovider "github.com/knadh/koanf/providers/env"
	fileprovider "github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

// WithProviders adds config provider
type WithProviders func(*koanf.Koanf)

// NewClient создае новый, пред-настроенный Koanf клиент для работы с конфигами
func NewClient(providers ...WithProviders) *koanf.Koanf {
	client := koanf.New(".")
	for _, apply := range providers {
		apply(client)
	}

	return client
}

// WithEnvProvider провайдер для работы с перемеными окружения
func WithEnvProvider(prefix string, parsers map[string]func(string) any) func(*koanf.Koanf) {
	return func(client *koanf.Koanf) {
		provider := envprovider.ProviderWithValue(prefix, "_", func(key string, value string) (string, any) {
			parser, ok := parsers[key]
			if !ok {
				return strings.ToLower(strings.TrimPrefix(key, fmt.Sprintf("%s_", prefix))), value
			}

			return strings.ToLower(strings.TrimPrefix(key, fmt.Sprintf("%s_", prefix))), parser(value)
		})

		if err := client.Load(provider, nil); err != nil {
			logger.Fatalf("an error occurred when loading the application configuration from the provider %T: %s", provider, err)
		}
	}
}

// WithFileProvider провайдер для работы с файлами конфигурации
func WithFileProvider(path string) func(*koanf.Koanf) {
	return func(client *koanf.Koanf) {
		provider := fileprovider.Provider(path)

		if err := client.Load(provider, json.Parser()); err != nil {
			logger.Fatalf("an error occurred when loading the application configuration from the provider %T: %s", provider, err)
		}
	}
}
