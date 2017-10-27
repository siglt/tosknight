package crawler

import (
	"path/filepath"

	"github.com/asciimoo/colly"
	"github.com/siglt/tosknight/pkg/file"
	"github.com/siglt/tosknight/pkg/git"
	"github.com/siglt/tosknight/source"
	"github.com/siglt/tosknight/util"
	log "github.com/sirupsen/logrus"
)

func parseResponse(response *colly.Response, source source.Source) {
	if response.StatusCode != 200 {
		log.Errorf("The status code of the response %v is %d", response.Request.AbsoluteURL, response.StatusCode)
	}

	URL := source.URL
	gitManager := git.NewManager(StoragePath)
	directory := filepath.Join(StoragePath, util.GetFileName(URL))
	fileManager := file.NewManager(directory, gitManager)

	if err := fileManager.SaveFile(response.Body, source); err != nil {
		log.Errorf("%s: %v", URL, err)
	}
}
