package file

import "os/exec"

const (
	html2textCommand = "html2text"
)

// HTML2Markdown converts HTML to Markdown-formatted text.
func HTML2Markdown(content []byte, HTMLFile string) ([]byte, error) {
	html2textCmd := exec.Command(html2textCommand, HTMLFile)
	output, err := html2textCmd.CombinedOutput()
	if err != nil {
		return nil, err
	}
	return output, nil
}
