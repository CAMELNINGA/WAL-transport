package config

import (
	"fmt"
	"time"

	"github.com/CAMELNINGA/cdc-postgres/pkg/postgres"
	"github.com/spf13/viper"
)

type Config struct {
	Database  postgres.DatabaseCfg `mapstructure:"db"`
	Listener  Listener             `mapstructure:"listener"`
	LoggerCfg LoggerCfg
}

type Listener struct {
	RefreshConnection time.Duration
}

// LoggerCfg path of the logger config.
type LoggerCfg struct {
	Caller bool
	Level  string
	Format string
}

func NewConfig() *Config {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	err := viper.ReadInConfig()

	if err != nil {
		fmt.Printf("%v", err)
	}

	conf := &Config{}
	err = viper.Unmarshal(conf)
	if err != nil {
		fmt.Printf("unable to decode into config struct, %v", err)
	}

	return conf
}
