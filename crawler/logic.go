package crawler

import (
	"fmt"
	"path/filepath"

	"github.com/asciimoo/colly"
	"github.com/gaocegege/tosknight/pkg/git"
	"github.com/siglt/tosknight/pkg/meta"
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
	bufFilePath := filepath.Join(directory, util.BufFileName)
	latestfilePath := filepath.Join(directory, util.LatestFileName)

	util.MkDir(directory)
	_, err := util.SaveFile(response.Body, bufFilePath)
	if err != nil {
		log.Errorf("Failed to save content to file %s: %v", bufFilePath, err)
	}

	// If there is no file now, just create the directory and the file.
	if util.IsFileExists(latestfilePath) == false {
		log.Infof("Create latest file for %s", URL)
		_, err = util.SaveFile(response.Body, latestfilePath)
		if err != nil {
			log.Errorf("Failed to save content to file %s: %v", latestfilePath, err)
		}

		// Write the URL to the meta file to generate UI.
		if err := meta.WriteMeta(directory, source); err != nil {
			log.Errorf("Failed to save meta data for %s: %v", directory, err)
		}

		// Commit the changes to storage repo.
		gitManager.AddAndCommit(fmt.Sprintf("%s: First Commit", response.Request.URL))
		return
	}

	isModified, err := util.IsModified(bufFilePath, latestfilePath)
	if err != nil {
		log.Errorf("Failed to check if the content is modified: %v", err)
	}
	if isModified == true {
		persistentFileName := util.PersistentFileName()
		persistentFilePath := filepath.Join(directory, persistentFileName)
		log.Infof("The content is changed, write to %s", persistentFilePath)

		// Save the latest content to the persistent file
		// and save the new content to the latest file.
		util.CopyFile(latestfilePath, persistentFilePath)
		util.CopyFile(bufFilePath, latestfilePath)

		// Commit the persistent file and latest file.
		gitManager.AddAndCommit(fmt.Sprintf("%s: Update %s", response.Request.URL, persistentFileName))
	} else {
		log.Debugf("There is no content changed in %s", URL)
	}
}
