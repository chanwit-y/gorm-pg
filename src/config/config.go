package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Configuration struct {
	Database DatabaseConfiguration
}

type DatabaseConfiguration struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
}

func New() *Configuration {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading congig file %s \n", err)
	}

	var configuration Configuration
	viper.Unmarshal(&configuration)

	return &configuration
}
