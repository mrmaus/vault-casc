package internal

import (
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
)

//todo:
func FileToString(path string) (string, error) {
	content, err := ioutil.ReadFile(path)
	return string(content), err
}

//todo:
func StringToFile(content string, path string) {
	dir := filepath.Dir(path)
	if !folderExists(dir) {
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
	err := ioutil.WriteFile(path, []byte(content), 0644)
	if err != nil {
		panic(err)
	}
}

// Returns file name without extension
func TrimFileExtension(file string) string {
	return strings.TrimSuffix(file, path.Ext(file))
}

//todo:
func AppendFileExtension(file string, extension string) string {
	return file + "." + extension
}

func folderExists(dir string) bool {
	if _, err := os.Stat(dir); !os.IsNotExist(err) {
		return true
	}
	return false
}
