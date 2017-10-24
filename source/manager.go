package source

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

// AddSource adds a source to manager.
func (m *Manager) AddSource(s Source) {
	m.Sources = append(m.Sources, s)
}
