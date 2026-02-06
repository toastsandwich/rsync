package utils

import (
	"os"
	"path/filepath"
)

const path = ".rsync/config.json"

func ConfigPath() (fullpath string, dir string, err error) {
	var home string
	home, err = os.UserHomeDir()
	if err != nil {
		return "", "", err
	}
	fullpath = filepath.Join(home, path)
	dir = filepath.Dir(fullpath)
	return
}

func OpenConfig() (*os.File, error) {
	fp, _, err := ConfigPath()
	if err != nil {
		return nil, err
	}

	return os.OpenFile(fp, os.O_RDWR, 0777)
}
