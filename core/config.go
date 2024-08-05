package core

import (
	"github.com/fsnotify/fsnotify"
	"github.com/gofiber/fiber/v2/log"
	"github.com/spf13/viper"
)

type Server struct {
	Address string
	Port    int
}

type Logger struct{}

type StargazerConfig struct {
	Server Server
	Logger Logger
}

func NewStargazerConfig() *StargazerConfig {
	var config StargazerConfig

	v := viper.New()
	v.SetConfigType("toml")
	v.AddConfigPath("/etc/stargazer")
	v.AddConfigPath("$HOME/.stargazer")
	v.AddConfigPath(".")
	v.SetConfigFile("stargazer.toml")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			if err = v.SafeWriteConfig(); err != nil {
				log.Fatalf("Failed to write default config file: %s", err)
			}
		} else {
			log.Fatalf("Failed to read config file: %s", err)
		}
	}

	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Infof("Config file changed: %s", e.Name)
	})
	viper.WatchConfig()

	if err := v.Unmarshal(&config); err != nil {
		log.Fatalf("Failed to unmarshal config: %s", err)
	}

	return &config
}
