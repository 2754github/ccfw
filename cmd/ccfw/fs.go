package main

import (
	"os"
	"path/filepath"
)

var _fs fs

type fs struct{}

func (fs *fs) read(filePath string) ([]byte, error) {
	return os.ReadFile(filepath.Clean(filePath))
}

// If the file already exists, it does nothing.
func (fs *fs) write(filePath string, data []byte) error {
	_, err := os.Stat(filePath)
	if err == nil {
		return nil
	}

	err = os.MkdirAll(filepath.Dir(filePath), 0o755) //nolint:gosec,mnd
	if err != nil {
		return err
	}

	return os.WriteFile(filePath, data, 0o644) //nolint:gosec,mnd
}
