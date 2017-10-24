package config

import (
	"log"

	"github.com/spf13/viper"
)

// Buggy
func ParseConfigFile(configFile string) {
	viper.AddConfigPath(configFile)
	viper.AutomaticEnv()
	viper.ReadInConfig()
	log.Println(viper.Get("webs"))
}
