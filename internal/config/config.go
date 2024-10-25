package config

import (
	"github.com/cristalhq/aconfig"
)

var cfg Config

type Postgres struct {
	Host        string `env:"HOST"`
	Port        int    `env:"PORT"`
	User        string `env:"USER"`
	Password    string `env:"PASSWORD"`
	Database    string `env:"DATABASE"`
	SSLMode     string `env:"SSL_MODE"`
	SSLCertPath string `env:"SSL_CERT_PATH"`
	NeedMigrate bool   `env:"NEED_MIGRATE"`
}

type Milvus struct {
	Address string `env:"ADDRESS"`
}

type Config struct {
	Debug bool `env:"DEBUG"`

	Address string `env:"ADDRESS"`

	Postgres Postgres `env:"POSTGRES"`
	Milvus   Milvus   `env:"MILVUS"`
}

func Load() error {
	err := aconfig.LoaderFor(&cfg, aconfig.Config{
		EnvPrefix: "CLOWO",
	}).Load()

	return err
}

func Get() Config {
	return cfg
}
