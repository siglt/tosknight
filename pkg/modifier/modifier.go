package modifier

// Modifier is the interface for modifier.
type Modifier interface {
	IsModified(newFile string, oldFile string) (bool, error)
}
