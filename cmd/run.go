package cmd

import (
	"path/filepath"

	"github.com/siglt/tosknight/config"
	"github.com/siglt/tosknight/crawler"
	"github.com/siglt/tosknight/source"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// SOURCEFILE is the name of source file config.
const SOURCEFILE = "sourceFile"

var sourceFile string

// runCmd represents the run command.
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "run the spider logic",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		run()
	},
}

func init() {
	RootCmd.AddCommand(runCmd)

	runCmd.Flags().StringVarP(&sourceFile, SOURCEFILE, "s", "", "Source directory to read from")
}

func run() error {
	// Read the source config file.
	if sourceFile == "" {
		log.Fatalln("There is no source file given")
	}
	abs, err := filepath.Abs(sourceFile)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(abs)
	if err = config.ParseConfigFile(sourceFile); err != nil {
		log.Fatalln(err)
	}

	sourceManager := source.NewManager()
	sourceManager.ReadSourcesFromConfig()
	contentCrawler := crawler.New(sourceManager, "/home/ist/go/src/github.com/siglt/tosknight-storage")
	contentCrawler.Run()
	log.Println("Run called.")
	return nil
}
