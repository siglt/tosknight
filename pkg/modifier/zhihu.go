package modifier

import "github.com/siglt/tosknight/util"

type ZhihuModifier struct {
	Modifier
}

func (dm DefaultModifier) IsModified(newFile string, oldFile string) (bool, error) {
	return util.IsModified(newFile, oldFile)
}
