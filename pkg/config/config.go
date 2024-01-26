package config

import (
	"log"

	"github.com/spf13/viper"
)

// mapping the config vars with struct
type Config struct {
	DBUser           string `mapstructure:"DBUSER"`
	DBPass           string `mapstructure:"DBPASS"`
	DBIp             string `mapstructure:"DBIP"`
	DBName           string `mapstructure:"DBNAME"`
	Port             string `mapstructure:"PORT"`
	JwtSecret        string `mapstructure:"JWT_SECRET"`
	JwtExpireMinutes int    `mapstructure:"JWT_EXPIRE_MINUTES"`
}

// initializing config vars
func InitConfig() *Config {
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	//automatically reads the config vars
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading env file", err)
	}
	var config *Config
	//converts the read config vars into mapped struct type
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal("Error reading env file", err)
	}
	return config
}

// global var to access from any package
var LocalConfig *Config

// calling initconfig() to initialize the config vars
func SetConfig() {
	LocalConfig = InitConfig()
}
