package crawler

import (
	"sync"

	"github.com/asciimoo/colly"
	"github.com/siglt/tosknight/source"
	"github.com/siglt/tosknight/util"
	log "github.com/sirupsen/logrus"
)

// StoragePath is the path to storage directory.
var StoragePath = ""

// Crawler is the type for spider.
type Crawler struct {
	SourceManager *source.Manager
	waitGroup     *sync.WaitGroup
}

// New returns a new Crawler.
func New(s *source.Manager, path string) *Crawler {
	err := util.IsGitDir(path)
	if err != nil {
		log.Fatalln(err)
	}

	StoragePath = path

	var wg sync.WaitGroup
	wg.Add(len(s.Sources))
	return &Crawler{
		SourceManager: s,
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
	collector.OnResponse(parseResponse)
	collector.Visit(s.URL)
}
