package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Database        string `mapstructure:"DB_DATABASE"`
	Dsn             string `mapstructure:"DB_SOURCE"`
	DbUsername      string `mapstructure:"DB_USERNAME"`
	DbPassword      string `mapstructure:"DB_PASSWORD"`
	DbAuthsource    string `mapstructure:"DB_AUTHSOURCE"`
	DbAuthmechanism string `mapstructure:"DB_AUTHMECHANISM"`
	ServerPort      string `mapstructure:"SERVER_PORT"`
	ElasticURL      string `mapstructure:"ES_SOURCE"`
	RandommerKey    string `mapstructure:"RANDOMMER_KEY"`
}

func LoadConfig(path string) (*Config, error) {
	viper.AddConfigPath(path)
	viper.AddConfigPath(".")
	// viper.SetConfigName("app")
	// viper.SetConfigType("env")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}
	return &config, nil

}
