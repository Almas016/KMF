package cmd

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Port string
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		log.Println("No config file found, using default configuration.")
	}

	viper.SetDefault("Port", "8080")

	config := &Config{
		Port: viper.GetString("Port"),
	}

	return config, nil
}
