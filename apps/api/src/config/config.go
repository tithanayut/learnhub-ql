package config

import (
	"github.com/spf13/viper"
	"log"
)

type AppConfig struct {
	Database DatabaseConfig `mapstructure:"database"`
}

type DatabaseConfig struct {
	Host         string `mapstructure:"host"`
	Port         string `mapstructure:"port"`
	Username     string `mapstructure:"username"`
	Password     string `mapstructure:"password"`
	DatabaseName string `mapstructure:"databaseName"`
}

func LoadAppConfig() AppConfig {
	viper.AddConfigPath("./src/config")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	appConfig := AppConfig{}
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("failed reading the config: %v", err)
	}
	if err := viper.Unmarshal(&appConfig); err != nil {
		log.Fatalf("failed unmarshalling the config: %v", err)
	}
	return appConfig
}
