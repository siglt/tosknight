package meta

import (
	"path/filepath"

	"github.com/siglt/tosknight/util"
	"gopkg.in/yaml.v2"
)

const (
	metaFileName = ".meta.yml"
)

type Meta struct {
	Name string `yaml:"name"`
}

func WriteMeta(directory, URL string) error {
	meta := Meta{
		Name: URL,
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
