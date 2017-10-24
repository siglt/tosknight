package crawler

import (
	"path/filepath"
	"sync"

	"github.com/asciimoo/colly"
	"github.com/gaocegege/tosknight/source"
	"github.com/gaocegege/tosknight/util"
	log "github.com/sirupsen/logrus"
)

// Crawler is the type for spider.
type Crawler struct {
	StoragePath   string
	SourceManager *source.Manager
	waitGroup     *sync.WaitGroup
}

// New returns a new Crawler.
func New(s *source.Manager, path string) *Crawler {
	err := util.IsGitDir(path)
	if err != nil {
		log.Fatalln(err)
	}

	var wg sync.WaitGroup
	wg.Add(len(s.Sources))
	return &Crawler{
		SourceManager: s,
		StoragePath:   path,
		waitGroup:     &wg,
	}
}

// Run runs the spider workers in parallel.
func (c Crawler) Run() {
	for _, source := range c.SourceManager.Sources {
		go c.parse(source)
	}
	c.waitGroup.Wait()
	log.Println("The crawler has done its work :)")
}

func (c Crawler) parse(s source.Source) {
	defer c.waitGroup.Add(-1)

	collector := colly.NewCollector()

	collector.OnResponse(func(response *colly.Response) {
		if response.StatusCode != 200 {
			log.Errorf("The status code of the response %v is %d", response.Request.AbsoluteURL, response.StatusCode)
		}

		directory := filepath.Join(c.StoragePath, util.GetFileName(s.URL))
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

		// If the file is modified, save it into a new file called `<date>.html`
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
			log.Infof("There is no content changed in %s", s.URL)
		}
	})

	collector.Visit(s.URL)
}
