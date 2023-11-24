package lib

import (
	"os"
	"path"
)

func ReadFile(name string) ([]byte, error) {

	dir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	file, err := os.ReadFile(path.Join(dir, name))
	if err != nil {
		return nil, err
	}
	return file, nil
}

func WriteFile(name string, data []byte) error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	f, err := os.Create(path.Join(dir, name))
	if err != nil {
		return err
	}

	if _, err := f.Write(data); err != nil {
		return err
	}
	return nil
}
