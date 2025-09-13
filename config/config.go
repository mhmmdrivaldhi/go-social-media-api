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
	
	AppConfig = &Config{
		ApplicationName: viper.GetString("APP_NAME"),
		APIPort:         viper.GetString("API_PORT"),	
		Port:            viper.GetString("DB_PORT"),
		Host:            viper.GetString("DB_HOST"),
		Username:        viper.GetString("DB_USERNAME"),
		Password:        viper.GetString("DB_PASSWORD"),
		DBName:          viper.GetString("DB_NAME"),
		DBUrl:           viper.GetString("DB_URL"),
	}
}