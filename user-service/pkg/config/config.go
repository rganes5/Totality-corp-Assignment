package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Port string `mapstructure:"PORT"`
}

func LoadConfig() (config Config, err error) {

	viper.AddConfigPath("./")
	viper.AddConfigPath("./app")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err = viper.ReadInConfig(); err != nil {
		log.Println("Error reading the env file from Api-Gateway via Viper", err)
		return
	}

	if err = viper.Unmarshal(&config); err != nil {
		log.Println("Error Unmarshalling the config from Api-Gateway via Viper", err)
		return
	}

	return
}
