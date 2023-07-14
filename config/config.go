package config

import "github.com/spf13/viper"

type config struct {
	PORT        string
	DB_HOST     string
	DB_USER     string
	DB_PASSWORD string
	DB_DATABASE string
	DB_PORT     string
}

var ENV Config

func LoadConfig() {
	viper.AddConfigPath()
}
