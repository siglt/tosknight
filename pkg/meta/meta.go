package meta

import (
	"path/filepath"

	"github.com/siglt/tosknight/source"
	"github.com/siglt/tosknight/util"
	"gopkg.in/yaml.v2"
)

const (
	metaFileName = ".meta.yml"
)

// Meta if the type for meta file.
type Meta struct {
	Name string `yaml:"name"`
	URL  string `yaml:url`
}

// WriteMeta writes meta data to the .meta.yml in directory.
func WriteMeta(directory string, source source.Source) error {
	meta := Meta{
		Name: source.Name,
		URL:  source.URL,
	}

	content, err := yaml.Marshal(meta)
	if err != nil {
		return err
	}

	_, err = util.SaveFile(content, filepath.Join(directory, metaFileName))
	if err != nil {
		return err
	}
	return nil
}
