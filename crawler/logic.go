package crawler

import (
	"path/filepath"

	"github.com/asciimoo/colly"
	"github.com/gaocegege/tosknight/util"
	log "github.com/sirupsen/logrus"
)

func parseResponse(response *colly.Response) {
	if response.StatusCode != 200 {
		log.Errorf("The status code of the response %v is %d", response.Request.AbsoluteURL, response.StatusCode)
	}

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
		return
	}

	isModified, err := util.IsModified(bufFilePath, latestfilePath)
	if err != nil {
		log.Errorf("Failed to check if the content is modified: %v", err)
	}
	if isModified == true {
		persistentFilePath := filepath.Join(directory, util.PersistentFileName())
		log.Debugf("The content is changed, write to %s", persistentFilePath)
		util.CopyFile(latestfilePath, persistentFilePath)
		util.CopyFile(bufFilePath, latestfilePath)
	} else {
		log.Infof("There is no content changed in %s", response.Request.URL.String())
	}
}
