package core

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/spf13/viper"
)

type Server struct {
	Address string `mapstructure:"address"`
	Port    int    `mapstructure:"port"`
	Prefork bool   `mapstructure:"prefork"`
	TLS     struct {
		Enabled  bool   `mapstructure:"enabled"`
		CertFile string `mapstructure:"cert_file"`
		KeyFile  string `mapstructure:"key_file"`
	} `mapstructure:"tls"`
}

type StargazerConfig struct {
	Server Server
}

func NewStargazerConfig() StargazerConfig {
	config := StargazerConfig{}

	v := viper.New()
	v.SetConfigType("toml")
	v.SetConfigFile("stargazer.toml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Failed to read config file: %s", err)
	}

	if err := v.Unmarshal(&config); err != nil {
		log.Fatalf("Failed to unmarshal config: %s", err)
	}

	return config
}
