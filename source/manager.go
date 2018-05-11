package source

import (
	"github.com/spf13/viper"

	"github.com/siglt/tosknight/config"
)

// Manager is the type for source manager.
type Manager struct {
	Sources []Source
}

// NewManager returns a new Manager.
func NewManager() *Manager {
	return &Manager{
		Sources: []Source{},
	}
}

// ReadSourcesFromConfig reads sources from the source file.
func (m *Manager) ReadSourcesFromConfig() {
	for _, category := range viper.Get(config.WEBS).([]interface{}) {
		categoryMap := category.(map[interface{}]interface{})
		categoryName := categoryMap[config.NAME].(string)
		for _, source := range categoryMap[config.ITEMS].([]interface{}) {
			sourceMap := source.(map[interface{}]interface{})
			m.AddSource(Source{
				URL:      sourceMap[config.ITEMURL].(string),
				Name:     sourceMap[config.ITEMNAME].(string),
				Category: categoryName,
			})
		}
	}
}

// AddSource adds a source to manager.
func (m *Manager) AddSource(s Source) {
	m.Sources = append(m.Sources, s)
}
