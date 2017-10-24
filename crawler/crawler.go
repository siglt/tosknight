package crawler

import (
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
	log.Println(c.SourceManager.Sources)
	for _, source := range c.SourceManager.Sources {
		log.Println(c.waitGroup)
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
			log.Errorln("The status code of the response %v is %d", response.Request.AbsoluteURL, response.StatusCode)
		}
	})

	collector.Visit(s.URL)
}
