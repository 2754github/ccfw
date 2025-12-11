package file

import (
	"os"
	"path/filepath"
)

func Read(path string) ([]byte, error) {
	return os.ReadFile(filepath.Clean(path))
}

// Write does nothing if the file already exists.
func Write(path string, data []byte) error {
	_, err := os.Stat(path)
	if err == nil {
		return nil
	}

	err = os.MkdirAll(filepath.Dir(path), 0o755) //nolint:gosec,mnd
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0o644) //nolint:gosec,mnd
}
