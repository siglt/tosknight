package util

import (
	"os/exec"

	log "github.com/sirupsen/logrus"
)

const (
	diffSubcommand = "git diff"
	defaultBranch  = "master"
	nameOnlyOption = "--name-only"
)

func IsModified(path string) (bool, error) {
	diffCmd := exec.Command("git diff", nameOnlyOption, defaultBranch)
	diffCmd.Dir = path
	bytes, err := diffCmd.CombinedOutput()
	if err != nil {
		return false, err
	}
	log.Println(string(bytes))
	return false, nil
}
