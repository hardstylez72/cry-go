package config

import (
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

type (
	Database struct {
		Conn string
	}

	ENV string

	Config struct {
		Env           ENV
		GRPCAddr      string
		GWAddr        string
		StaticPort    string
		ProxyPort     string
		Database      Database
		App           App
		Auth          Auth
		InstanceId    string
		MigrationsDir string
		TelegramToken string
		SnapshotHost  string
		Lazanya       string

		PrometheusPort     string
		JaegerUrl          string
		JaegerServiceName  string
		PayServiceGRPCAddr string

		HalperHost string

		AdminEmail string

		Standalone bool

		TestEVMWalletId      string
		TestStarkNetWalletId string

		AnkrToken string
	}

	App struct {
		Domain string
		Schema string
		Port   string
	}

	Auth struct {
		GoogleClientId     string
		GoogleClientSecret string
		CookieName         string

		RedirectOnSuccess string
		RedirectOnFailure string
	}
)

const (
	Dev   ENV = "dev"
	Local ENV = "local"
	Prod  ENV = "prod"
)

func envFromString(s string) ENV {
	switch s {
	case "dev":
		return Dev
	case "local":
		return Local
	case "prod":
		return Prod
	}
	panic("invalid env value: " + s)
}

var CFG *Config

func Load() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
	c := &Config{
		Env:        envFromString(mustenv("ENV")),
		GRPCAddr:   mayenv("GRPC_ADDR", ":90"),
		GWAddr:     mayenv("GW_ADDR", ":8083"),
		StaticPort: mayenv("STATIC_PORT", ":3333"),
		ProxyPort:  mayenv("PROXY_PORT", ":3334"),
		Database: Database{
			Conn: mustenv("POSTGRES_CONN"),
		},
		App: App{
			Domain: mayenv("DOMAIN", ""),
			Schema: mayenv("SCHEMA", ""),
			Port:   mayenv("PORT", ""),
		},
		Auth: Auth{
			GoogleClientId:     mayenv("GOOGLE_CLIENT_ID", ""),
			GoogleClientSecret: mayenv("GOOGLE_CLIENT_SECRET", ""),

			CookieName:        mayenv("COOKIE_NAME", ""),
			RedirectOnSuccess: mayenv("REDIRECT_ON_SUCCESS", ""),
			RedirectOnFailure: mayenv("REDIRECT_ON_FAILURE", ""),
		},
		InstanceId:           uuid.New().String(),
		MigrationsDir:        mayenv("MIGRATIONS_DIR", "/internal/server/migrations"),
		TelegramToken:        mayenv("TELEGRAM_TOKEN", ""),
		SnapshotHost:         mustenv("SNAPSHOT_ORG_HOST"),
		Lazanya:              mustenv("LAZANYA"),
		PrometheusPort:       mayenv("PROMETHEUS_PORT", ""),
		JaegerUrl:            mayenv("JAEGER_URL", ""),
		JaegerServiceName:    mayenv("JAEGER_SERVICE_NAME", "cry-backend"),
		PayServiceGRPCAddr:   mustenv("PAY_SERVICE_GRPC_ADDR"),
		HalperHost:           mustenv("HALPER_HOST"),
		AdminEmail:           mustenv("ADMIN_EMAIL"),
		Standalone:           mayenv("STANDALONE", "false") == "true",
		TestEVMWalletId:      mustenv("TEST_EVM_WALLET_ID"),
		TestStarkNetWalletId: mustenv("TEST_STARKNET_WALLET_ID"),
		AnkrToken:            mustenv("ANKR_TOKEN"),
	}

	CFG = c

	return c, nil
}
func mayenv(key string, defaultValue string) string {
	e := os.Getenv(key)
	if e != "" {
		return e
	}
	return defaultValue
}
func mustenv(key string) string {
	v := os.Getenv(key)
	if v == "" {
		panic(key + " is not set")
	}
	return v
}
