package util

import (
	"os/exec"
	"syscall"

	log "github.com/sirupsen/logrus"
)

const (
	diffCommand = "diff"
)

// IsModified checks if the file is modified.
func IsModified(newFile string, oldFile string) (bool, error) {
	diffCmd := exec.Command(diffCommand, newFile, oldFile)
	if err := diffCmd.Start(); err != nil {
		log.Fatalf("cmd.Start: %v")
	}

	if err := diffCmd.Wait(); err != nil {
		if exiterr, ok := err.(*exec.ExitError); ok {
			// The program has exited with an exit code != 0

			// This works on both Unix and Windows. Although package
			// syscall is generally platform dependent, WaitStatus is
			// defined for both Unix and Windows and in both cases has
			// an ExitStatus() method with the same signature.
			if status, ok := exiterr.Sys().(syscall.WaitStatus); ok && status.ExitStatus() == 1 {
				return true, nil
			}
		} else {
			log.Fatalf("cmd.Wait: %v", err)
		}
	}
	return false, nil
}
