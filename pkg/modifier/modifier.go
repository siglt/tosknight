package modifier

type Modifier interface {
	func IsModified(newFile string, oldFile string) (bool, error)
}
