package cmd

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "tosknight",
	Short: "Take care of you in internet",
	Long:  ``,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) {},
}

func init() {
	RootCmd.PersistentFlags().BoolP("debug", "d", false, "set log level to debug")
	viper.BindPFlag("debug", RootCmd.PersistentFlags().Lookup("debug"))

	cobra.OnInitialize(initLogger)
}

func initLogger() {
	viper.GetBool("debug")
	if debug := viper.GetBool("debug"); debug == true {
		log.SetLevel(log.DebugLevel)
		log.Debugln("Set log level to DEBUG")
	}
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
