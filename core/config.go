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
	Secret  string
	TLS     struct {
		Enabled  bool   `mapstructure:"enabled"`
		CertFile string `mapstructure:"cert_file"`
		KeyFile  string `mapstructure:"key_file"`
	}
}

type Database struct {
	Type        string
	Host        string
	Port        int
	User        string
	Password    string
	Name        string
	DBFile      string
	TablePrefix string `mapstructure:"table_prefix"`
	SSLMode     string
	Migrate     bool
}

type Logger struct {
	LogLevel string `mapstructure:"log_level"`
	LogPath  string `mapstructure:"log_path"`
	LogName  string `mapstructure:"log_name"`
	LogExt   string `mapstructure:"log_ext"`
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
			Secret:  "",
			TLS: struct {
				Enabled  bool   `mapstructure:"enabled"`
				CertFile string `mapstructure:"cert_file"`
				KeyFile  string `mapstructure:"key_file"`
			}{
				Enabled:  false,
				CertFile: "",
				KeyFile:  "",
			},
		},
		Database: Database{
			Type:        "sqlite3",
			Port:        0,
			DBFile:      "stargazer.db",
			TablePrefix: "sg_",
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
