package config

import (
	"log"

	"github.com/spf13/viper"
)

type ServerConfig struct {
	ServerPort string `mapstructure:"server_port"`
}

func LoadConfig(path string, v any) error {
	viper.AddConfigPath(path)
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	if err := viper.Unmarshal(v); err != nil {
		return err
	}
	return nil
}
func NewServerConfig(path string) ServerConfig {
	var config ServerConfig
	if err := LoadConfig(path, &config); err != nil {
		log.Fatalf("error: %v", err)
	}
	return config
}

type RabbitMQConfig struct {
	Source   string `mapstructure:"rabbitmq_source"`
	Username string `mapstructure:"rabbitmq_username"`
	Password string `mapstructure:"rabbitmq_password"`
}

func NewRabbitMQConfig(path string) RabbitMQConfig {
	var config RabbitMQConfig
	if err := LoadConfig(path, &config); err != nil {
		log.Fatalf("error: %v", err)
	}
	return config
}
