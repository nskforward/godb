package godb

import "path/filepath"

type Storage struct {
	root           string
	payloadStorage map[string]map[string][]byte
	keysStorage    map[string][]string
	diskTableMx    *TwoLevelMutex
	memTableMx     *TwoLevelMutex
}

func NewStorage(dirname string) *Storage {
	return &Storage{
		root:           filepath.Join(ProcessDir(), dirname),
		payloadStorage: make(map[string]map[string][]byte),
		keysStorage:    make(map[string][]string),
		diskTableMx:    NewTwoLevelMutex(),
		memTableMx:     NewTwoLevelMutex(),
	}
}
