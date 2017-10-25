package source

import (
	"github.com/siglt/tosknight/config"
	"github.com/spf13/viper"
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
	for _, URL := range viper.GetStringSlice(config.WEBS) {
		m.AddSource(Source{
			URL: URL,
		})
	}
}

// AddSource adds a source to manager.
func (m *Manager) AddSource(s Source) {
	m.Sources = append(m.Sources, s)
}
