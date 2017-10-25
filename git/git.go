package git

import (
	"os/exec"

	log "github.com/sirupsen/logrus"
)

const (
	gitCommand       = "git"
	commitSubCommand = "commit"
	messageOption    = "-m"

	addSubCommand = "add"
	allOption     = "--all"
)

type Manager struct {
	path string
	Repo string
}

func NewManager(path string) *Manager {
	return &Manager{
		path: path,
		Repo: path,
	}
}

func (m Manager) AddAll() error {
	addCmd := exec.Command(gitCommand, addSubCommand, allOption)
	addCmd.Dir = m.Repo
	output, err := addCmd.CombinedOutput()
	if err != nil {
		return err
	}
	log.Debugln("`git add --all` output:", string(output))
	return nil
}

func (m Manager) CreateCommit(message string) error {
	commitCmd := exec.Command(gitCommand, commitSubCommand, messageOption, message)
	commitCmd.Dir = m.Repo
	output, err := commitCmd.CombinedOutput()
	if err != nil {
		return err
	}
	log.Debugf("`git commit -s %s` output: \n%s", message, string(output))
	return nil
}

func (m Manager) AddAndCommit(message string) error {
	if err := m.AddAll(); err != nil {
		return err
	}
	if err := m.CreateCommit(message); err != nil {
		return err
	}
	return nil
}
