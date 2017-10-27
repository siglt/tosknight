package util

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"os"
)

func IsGitDir(path string) error {
	// Check the path is existed
	status, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("The status of the path %s is not expected: %v", path, err)
	}
	if status.IsDir() != true {
		return fmt.Errorf("The path %s is not a directory: %v", path)
	}
	// TODO: Check if the directory is a git repo.
	return nil
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

func GetFileName(URL string) string {
	hasher := sha1.New()
	hasher.Write([]byte(URL))
	return base64.URLEncoding.EncodeToString(hasher.Sum(nil))
}
