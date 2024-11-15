package config

import (
	"log"

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

	// Set default values (you can adjust these according to your preferences)
	viper.SetDefault("App.Port", "8080")
	viper.SetDefault("App.JWTSecret", "whatIsTheSecretAbout78")
	viper.SetDefault("Database.Host", "localhost")
	viper.SetDefault("Database.Port", 5432)
	viper.SetDefault("Database.User", "root")
	viper.SetDefault("Database.Password", "password")
	viper.SetDefault("Database.DBName", "boiler_db")
	viper.SetDefault("Logging.ELKHost", "localhost:9200")
	viper.SetDefault("Logging.APMHost", "localhost:8200")

	// Set config file path and name
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	// Allow environment variables to override config file values
	viper.AutomaticEnv()

	// Read in the configuration file
	if err := viper.ReadInConfig(); err != nil {
		// Handle file not found or read errors
		log.Printf("Error reading config file, %s", err)
		// If you want, you can either exit or fallback to defaults.
	}

	// Unmarshal the configuration into the struct
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	// Return the loaded config
	return &config, nil
}

// package config

// import (
// 	"github.com/spf13/viper"
// )

// type Config struct {
// 	App struct {
// 		Port      string
// 		JWTSecret string `mapstructure:"jwt_secret"`
// 	}
// 	Database struct {
// 		Host     string
// 		Port     int
// 		User     string
// 		Password string
// 		DBName   string `mapstructure:"dbname"`
// 	}
// 	Logging struct {
// 		ELKHost string `mapstructure:"elk_host"`
// 		APMHost string `mapstructure:"apm_host"`
// 	}
// }

// func LoadConfig() (*Config, error) {
// 	var config Config

// 	viper.AddConfigPath("config")
// 	viper.SetConfigName("config")
// 	viper.SetConfigType("yaml")

// 	viper.AutomaticEnv()

// 	if err := viper.ReadInConfig(); err != nil {
// 		return nil, err
// 	}

// 	if err := viper.Unmarshal(&config); err != nil {
// 		return nil, err
// 	}

// 	return &config, nil
// }
