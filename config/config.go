package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"sync"
)

type Config struct {
	IsDebug       bool `env:"DEBUG" env-default:"false"`
	IsDevelopment bool `env:"DEVEL" env-default:"true"`
	Listen        struct {
		Type   string `env:"LISTEN_TYPE" env-default:"port"`
		BindIP string `env:"HOST" env-default:"0.0.0.0"`
		Port   string `env:"DEBUG" env-default:"8088"`
	}
	AppConfig struct {
		LogLevel  string
		AdminUser struct {
			Email    string `env:"ADMIN_EMAIL" env-required:"true"`
			Password string `env:"ADMIN_PWD" env-required:"true"`
		}
	}
}

var config *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		log.Print("gather config")

		config = &Config{}

		if err := cleanenv.ReadEnv(config); err != nil {
			log.Fatal(err)
		}
	})
	return config
}
