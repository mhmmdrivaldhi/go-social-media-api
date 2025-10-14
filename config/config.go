package config

import "github.com/spf13/viper"

type AppConfig struct {
	ApplicationName string
	APIPort         string
	Port            string
	Host            string
	Username        string
	Password        string
	DBName          string
	DBUrl           string
}

type JWTConfig struct {
	JWTSignatureKey string
	JWTSigningMethod string
	AccessTokenLifeTime int
}

type Config struct {
	AppConfig
	JWTConfig
}

func (cfg *Config) LoadConfig() {
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	
	cfg.AppConfig = AppConfig{
		ApplicationName: viper.GetString("APP_NAME"),
		APIPort:         viper.GetString("API_PORT"),
		Port:            viper.GetString("PORT"),
		Host: 		  viper.GetString("DB_HOST"),
		Username: 	  viper.GetString("DB_USERNAME"),
		Password: 	  viper.GetString("DB_PASSWORD"),
		DBName: 	  viper.GetString("DB_NAME"),
		DBUrl: 		  viper.GetString("DB_URL"),
	}

	cfg.JWTConfig = JWTConfig{
		JWTSignatureKey: viper.GetString("JWT_SIGNATURE_KEY"),
		JWTSigningMethod: viper.GetString("JWT_SIGNING_METHOD"),
		AccessTokenLifeTime: 24,
	}
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	cfg.LoadConfig()
	return cfg, nil
}