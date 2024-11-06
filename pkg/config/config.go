package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	App struct {
		Port      string
		JWTSecret string `mapstructure:"jwt_secret"`
	}
	Database struct {
		Host     string
		Port     int
		User     string
		Password string
		DBName   string `mapstructure:"dbname"`
	}
	Logging struct {
		ELKHost string `mapstructure:"elk_host"`
		APMHost string `mapstructure:"apm_host"`
	}
}

func LoadConfig() (*Config, error) {
	var config Config

	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
