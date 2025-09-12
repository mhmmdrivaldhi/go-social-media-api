package config

import "github.com/spf13/viper"

type Config struct {
	ApplicationName string
	APIPort         string
	Port            string
	Host            string
	Username        string
	Password        string
	DBName          string
	DBUrl           string
}

var AppConfig *Config

func LoadConfig() {
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&AppConfig)
	if err != nil {
		panic(err)
	}

}