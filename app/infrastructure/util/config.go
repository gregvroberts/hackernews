package util

import (
	"errors"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	AppPort    string `mapstructure:"APP_PORT"`
	DBDriver   string `mapstructure:"DB_DRIVER"`
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     uint16 `mapstructure:"DB_PORT"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBName     string `mapstructure:"DB_NAME"`
	DBSSLMode  string `mapstructure:"DB_SSL_MODE"`
}

func LoadConfig(configPath string, version string) (*Config, error) {
	config := Config{}
	viper.AddConfigPath(configPath)

	switch version {
	case "dev":
		viper.SetConfigName("dev")
	case "prod":
		viper.SetConfigName("prod")
	case "test":
		viper.SetConfigName("test")
	default:
		err := errors.New("Config path must be \"dev\", \"prod\", or \"test\"")
		log.Println(err)
		return nil, err
	}

	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&config)
	return &config, nil
}
