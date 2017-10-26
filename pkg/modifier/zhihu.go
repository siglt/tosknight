package modifier

type ZhihuModifier struct {
	Modifier
}

func (zm ZhihuModifier) IsModified(newFile string, oldFile string) (bool, error) {
	return false, nil
}
