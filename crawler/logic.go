package crawler

import (
	"fmt"
	"path/filepath"

	"github.com/asciimoo/colly"
	"github.com/gaocegege/tosknight/git"
	"github.com/siglt/tosknight/util"
	log "github.com/sirupsen/logrus"
)

func parseResponse(response *colly.Response) {
	if response.StatusCode != 200 {
		log.Errorf("The status code of the response %v is %d", response.Request.AbsoluteURL, response.StatusCode)
	}

	gitManager := git.NewManager(StoragePath)
	directory := filepath.Join(StoragePath, util.GetFileName(response.Request.URL.String()))
	bufFilePath := filepath.Join(directory, util.BufFileName)
	latestfilePath := filepath.Join(directory, util.LatestFileName)

	util.MkDir(directory)
	_, err := util.SaveFile(response.Body, bufFilePath)
	if err != nil {
		log.Errorf("Failed to save content to file %s: %v", bufFilePath, err)
	}

	// If there is no file now, just create the directory and the file.
	if util.IsFileExists(latestfilePath) == false {
		log.Debugln("Create latest file")
		_, err = util.SaveFile(response.Body, latestfilePath)
		if err != nil {
			log.Errorf("Failed to save content to file %s: %v", latestfilePath, err)
		}
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
		log.Debugf("The content is changed, write to %s", persistentFilePath)
		util.CopyFile(latestfilePath, persistentFilePath)
		util.CopyFile(bufFilePath, latestfilePath)
		gitManager.AddAndCommit(fmt.Sprintf("%s: Update %s", response.Request.URL, persistentFileName))
	} else {
		log.Infof("There is no content changed in %s", response.Request.URL.String())
	}
}
