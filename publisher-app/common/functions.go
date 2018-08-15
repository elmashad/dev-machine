package common

import (
	"github.com/spf13/viper"
	"log"
)

var AppConfig *viper.Viper

func SetConfig() {
	v := viper.New()
	v.SetConfigName("config")
	v.AddConfigPath(".")
	v.SetConfigType("toml")
	v.AutomaticEnv()
	v.WatchConfig()
	err := v.ReadInConfig() // Find and read the config file
	if err != nil { // Handle errors reading the config file
		log.Println("Skip reading from config file, Will read from env, " + err.Error())
	}

	AppConfig = v
}

func GetAppConfig() *viper.Viper {
	return AppConfig
}
func GetAppPort() string {
	return GetAppConfig().GetString("APP_PORT")
}