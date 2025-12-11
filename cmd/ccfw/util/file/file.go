package file

import (
	"os"
	"path/filepath"
)

// List returns an empty slice if the directory does not exist.
func List(root string) ([]string, error) {
	paths := make([]string, 0)

	_, err := os.Stat(root)
	if err != nil {
		return paths, nil //nolint:nilerr
	}

	err = filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !d.IsDir() {
			paths = append(paths, path)
		}

		return nil
	})

	return paths, err
}

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

// Remove does nothing if the file does not exist.
func Remove(path string) error {
	_, err := os.Stat(path)
	if err != nil {
		return nil //nolint:nilerr
	}

	return os.Remove(path)
}
