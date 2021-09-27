package godb

import (
	"os"
	"path/filepath"
	"strings"
)

func FileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func IsNameCorrect(name string) bool {
	if len(name) == 0 {
		return false
	}
	return !strings.ContainsAny(name, "/\\:*?\"<>|")
}

func ProcessDir() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}
	return dir
}
