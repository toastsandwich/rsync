package utils

import (
	"os"
	"path/filepath"
)

const path = ".rsync/config.json"

func HomeDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	return home
}

func ConfigPath() (fullpath string, dir string, err error) {
	fullpath = filepath.Join(HomeDir(), path)
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
