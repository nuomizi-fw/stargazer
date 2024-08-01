package core

import (
	"github.com/fsnotify/fsnotify"
	"github.com/gofiber/fiber/v2/log"
	"github.com/spf13/viper"
)

type ServerConfig struct {
	Address string
	Port    int
}

type LoggerConfig struct{}

type Auth struct {
	Secret    string
	Iteration int
}

type Config struct {
	Server ServerConfig
	Logger LoggerConfig
	Auth   Auth
}

func NewConfig() *Config {
	config := Config{}

	v := viper.New()
	v.SetConfigType("toml")
	v.AddConfigPath("/etc/stargazer")
	v.AddConfigPath("$HOME/.stargazer")
	v.AddConfigPath(".")
	v.SetConfigFile("stargazer.toml")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			// Write default config file
			if err = v.SafeWriteConfig(); err != nil {
				log.Fatalf("Failed to write default config file: %s", err)
			}
		} else {
			// Config file was found but another error was produced
			log.Fatalf("Failed to read config file: %s", err)
		}
	}

	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Infof("Config file changed: %s", e.Name)
	})
	viper.WatchConfig()

	if err := v.Unmarshal(&config); err != nil {
		panic(err)
	}

	return &config
}
