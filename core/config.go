package core

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/spf13/viper"
)

type Server struct {
	Host    string
	Port    string
	Debug   bool
	Prefork bool
}

type Database struct {
	Type        string
	Host        string
	Port        int
	User        string
	Password    string
	Name        string
	DBFile      string
	TablePrefix string
	SSLMode     string
	Migrate     bool
}

type Logger struct {
	LogLevel string
	LogPath  string
	LogName  string
	LogExt   string
}

type StargazerConfig struct {
	Server   Server
	Database Database
	Logger   Logger
}

func defaultStargazerConfig() StargazerConfig {
	return StargazerConfig{
		Server: Server{
			Host:    "localhost",
			Port:    "11451",
			Debug:   true,
			Prefork: false,
		},
		Database: Database{
			Type:        "sqlite3",
			Port:        0,
			Host:        "localhost",
			User:        "",
			Password:    "",
			Name:        "stargazer",
			DBFile:      "stargazer.db",
			TablePrefix: "sg_",
			SSLMode:     "disable",
			Migrate:     true,
		},
		Logger: Logger{
			LogLevel: "debug",
			LogPath:  "logs",
			LogName:  "stargazer",
			LogExt:   "log",
		},
	}
}

func NewStargazerConfig() StargazerConfig {
	v := viper.New()
	v.AddConfigPath(".")
	v.AddConfigPath("$HOME/.config/.stargazer")
	v.AddConfigPath("/etc/stargazer")
	v.SetConfigName("stargazer")
	v.SetConfigType("toml")

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			if err = v.SafeWriteConfigAs("stargazer.toml"); err != nil {
				log.Fatalf("Failed to write config: %s", err)
			}
		} else {
			log.Fatalf("Failed to read config: %s", err)
		}
	}

	config := defaultStargazerConfig()

	if err := v.Unmarshal(&config); err != nil {
		log.Fatalf("Failed to unmarshal config: %s", err)
	}

	return config
}
