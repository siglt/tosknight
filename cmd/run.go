package cmd

import (
	"path/filepath"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/siglt/tosknight/config"
	"github.com/siglt/tosknight/crawler"
	"github.com/siglt/tosknight/source"
)

const (
	// SOURCEFILE is the name of source file config.
	SOURCEFILE = "sourceFile"
	STORAGEDIR = "storageDir"
)

var (
	sourceFile string
	storageDir string
)

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

	runCmd.Flags().StringVarP(&sourceFile, SOURCEFILE, "c", "", "Path to source.yml")
	runCmd.Flags().StringVarP(&storageDir, STORAGEDIR, "s", "", "Path to the storage directory")
}

func run() error {
	// Read the source config file.
	if sourceFile == "" {
		log.Fatalln("There is no source file given")
	}
	absSourceFiles, err := filepath.Abs(sourceFile)
	if err != nil {
		log.Fatalln(err)
	}
	if err = config.ParseConfigFile(absSourceFiles); err != nil {
		log.Fatalln(err)
	}

	if storageDir == "" {
		log.Fatalln("There is no storage directory given")
	}
	absStorageDir, err := filepath.Abs(storageDir)

	sourceManager := source.NewManager()
	sourceManager.ReadSourcesFromConfig()
	contentCrawler := crawler.New(sourceManager, absStorageDir)
	contentCrawler.Run()
	log.Println("Run called.")
	return nil
}
