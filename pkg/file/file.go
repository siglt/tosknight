package file

import (
	"fmt"
	"path/filepath"

	"github.com/siglt/tosknight/pkg/git"
	"github.com/siglt/tosknight/pkg/meta"
	"github.com/siglt/tosknight/source"
	"github.com/siglt/tosknight/util"
	log "github.com/sirupsen/logrus"
)

const (
	BufHTMLFileName        = ".buf.html"
	BufMarkdownFileName    = ".buf.md"
	LatestHTMLFileName     = ".latest.html"
	LatestMarkdownFileName = "latest.md"
)

// Manager is the type for file manager.
type Manager struct {
	directory  string
	gitManager *git.Manager
}

// NewManager returns a new manager.
func NewManager(directory string, gm *git.Manager) *Manager {
	MkDir(directory)
	return &Manager{
		directory:  directory,
		gitManager: gm,
	}
}

// SaveFile saves all file related to one source.
func (m Manager) SaveFile(content []byte, source source.Source) error {
	// Save file to buf file.
	if err := m.SaveBuf(content); err != nil {
		return err
	}

	// Save file to latest file.
	if err := m.SaveLatest(content, source); err != nil {
		return err
	}

	if err := m.SaveBackup(content); err != nil {
		return err
	}
	return nil
}

// SaveBuf saves buf in HTML and Markdown.
func (m Manager) SaveBuf(content []byte) error {
	bufHTML := filepath.Join(m.directory, BufHTMLFileName)
	_, err := SaveFile(content, bufHTML)
	if err != nil {
		log.Errorf("Failed to save content to file %s: %v", bufHTML, err)
	}

	bufMarkdown := filepath.Join(m.directory, BufMarkdownFileName)
	markdownContent, err := HTML2Markdown(content, bufHTML)
	if err != nil {
		return err
	}

	_, err = SaveFile(markdownContent, bufMarkdown)
	if err != nil {
		log.Errorf("Failed to save content to file %s: %v", bufMarkdown, err)
	}
	return nil
}

// SaveLatest Saves latest file.
func (m Manager) SaveLatest(content []byte, source source.Source) error {
	latestHTML := filepath.Join(m.directory, LatestHTMLFileName)
	// If there is no file now, just create the directory and the file.
	if IsFileExists(latestHTML) == false {
		log.Infof("Create latest file for %s", m.directory)
		_, err := SaveFile(content, latestHTML)
		if err != nil {
			log.Errorf("Failed to save content to file %s: %v", latestHTML, err)
		}

		latestMarkdown := filepath.Join(m.directory, LatestMarkdownFileName)
		markdownContent, err := HTML2Markdown(content, latestHTML)
		if err != nil {
			log.Errorf("Failed to convert markdown to %s: %v", latestMarkdown, err)
		}

		_, err = SaveFile(markdownContent, latestMarkdown)
		if err != nil {
			log.Errorf("Failed to save content to file %s: %v", latestMarkdown, err)
		}

		// Write the m.directory to the meta file to generate UI.
		if err := meta.WriteMeta(m.directory, source); err != nil {
			log.Errorf("Failed to save meta data for %s: %v", m.directory, err)
		}

		// Commit the changes to storage repo.
		gitManager.AddAndCommit(fmt.Sprintf("%s: First Commit", response.Request.m.directory))
		return nil
	}
	return nil
}

// SaveBackup saves the snapshot of the tos.
func (m Manager) SaveBackup(content []byte) error {
	latestHTML := filepath.Join(m.directory, LatestHTMLFileName)
	bufHTML := filepath.Join(m.directory, BufHTMLFileName)
	bufMarkdown := filepath.Join(m.directory, BufMarkdownFileName)
	latestMarkdown := filepath.Join(m.directory, LatestMarkdownFileName)

	isModified, err := util.IsModified(bufMarkdown, latestMarkdown)
	if err != nil {
		log.Errorf("Failed to check if the content is modified: %v", err)
	}
	if isModified == true {
		persistentHTMLFileName := PersistentHTML()
		persistentHTML := filepath.Join(m.directory, persistentHTMLFileName)
		persistentMarkdownFileName := PersistentMarkdown()
		persistentMarkdown := filepath.Join(m.directory, persistentMarkdownFileName)
		log.Infof("The content is changed, write to %s and %s", persistentHTML, persistentMarkdown)

		// Save the latest content to the persistent file
		// and save the new content to the latest file.
		CopyFile(latestHTML, persistentHTML)
		CopyFile(bufHTML, latestHTML)
		CopyFile(latestMarkdown, persistentMarkdown)
		CopyFile(bufMarkdown, latestMarkdown)

		// Commit the persistent file and latest file.
		gitManager.AddAndCommit(fmt.Sprintf("%s: Update %s", response.Request.m.directory, persistentFileName))
	} else {
		log.Debugf("There is no content changed in %s", m.directory)
	}
	return nil
}
