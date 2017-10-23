package source

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func ParseSourceFile(sourceFile string) {
	viper.AddConfigPath(sourceFile)
	viper.AutomaticEnv()
	viper.ReadInConfig()
	log.Println(viper.Get("Web"))
}
