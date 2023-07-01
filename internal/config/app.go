package config

import (
	"fmt"
	"github.com/IASamoylov/tg_calories_observer/internal/config/debug"
	"github.com/IASamoylov/tg_calories_observer/internal/pkg/koanf"
	"log"
	"strings"
)

const (
	configPath string = "./internal/config/json"
)

// App application configuration
type App struct {
	Postgres Postgres `koanf:"postgres"`
	Telegram Telegram `koanf:"telegram"`
}

// NewConfig creates a new application configuration
func NewConfig() *App {
	app := &App{}

	client := koanf.NewClient(
		koanf.WithFileProvider(fmt.Sprintf("%s/config.json", configPath)),
		koanf.WithFileProvider(fmt.Sprintf("%s/%s.config.json", configPath, debug.Version)),
		koanf.WithEnvProvider("APP"),
	)

	if err := client.Unmarshal("", app); err != nil {
		log.Fatal(err)
	}

	return app
}

// Postgres settings
type Postgres struct {
	Host              string `koanf:"host"`
	User              string `koanf:"user"`
	Pass              string `koanf:"pass"`
	SslMode           string `koanf:"ssl_mode"`
	ConnectionTimeout string `koanf:"connect_timeout"`
}

// Conn returns a postgres connection string
func (cfg Postgres) Conn() string {
	base := fmt.Sprintf("postgresql://%s:%s@%s/%s", cfg.User, cfg.Pass, cfg.Host, debug.AppName)

	args := []string{
		fmt.Sprintf("sslmode=%s", cfg.SslMode),
		fmt.Sprintf("connect_timeout=%s", cfg.ConnectionTimeout),
	}

	return fmt.Sprintf("%s?%s", base, strings.Join(args, "&"))
}

// Telegram settings
type Telegram struct {
	Token string `koanf:"token"`
}
