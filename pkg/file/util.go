package file

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"io"
	"os"
	"time"
)

var now = time.Now().Format(("2006-01-02-15:04:05"))

func PersistentHTML() string {
	return fmt.Sprintf(".%s.html", now)
}

func PersistentMarkdown() string {
	return fmt.Sprintf("%s.md", now)
}

func GetFileName(URL string) string {
	hasher := sha1.New()
	hasher.Write([]byte(URL))
	return base64.URLEncoding.EncodeToString(hasher.Sum(nil))
}

// IsFileExists returns true if the file exists.
func IsFileExists(fileName string) bool {
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		return false
	}
	return true
}

// OpenFile opens file from 'name', and create one if not exist.
func OpenFile(fileName string, flag int, perm os.FileMode) (*os.File, error) {
	var file *os.File
	var err error

	file, err = os.OpenFile(fileName, flag, perm)
	if err != nil && os.IsNotExist(err) {
		file, err = os.Create(fileName)
		if err != nil {
			return nil, err
		}
	}

	return file, err
}

// SaveFile saves the content to the file.
func SaveFile(content []byte, fileName string) (int, error) {
	file, err := OpenFile(fileName, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return 0, err
	}

	num, err := file.Write(content)
	if err != nil {
		return 0, err
	}

	return num, nil
}

func CopyFile(src string, dst string) {
	in, err := os.Open(src)
	if err != nil {
		return
	}
	defer in.Close()
	out, err := os.Create(dst)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		cerr := out.Close()
		if err == nil {
			err = cerr
		}
	}()
	if _, err = io.Copy(out, in); err != nil {
		return
	}
	err = out.Sync()
	return
}

func MkDir(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, os.ModePerm)
	}
}
