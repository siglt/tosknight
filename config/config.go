package config

import "github.com/spf13/viper"

// ParseConfigFile parses the source file and insert the
// content into viper.
func ParseConfigFile(configFile string) error {
	viper.SetConfigFile(configFile)
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}
