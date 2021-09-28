package godb

type Storage struct {
	root           string
	payloadStorage map[string]map[string][]byte
	keysStorage    map[string][]string
	diskTableMx    *TwoLevelMutex
	memTableMx     *TwoLevelMutex
}

func NewStorage(dir string) *Storage {
	if !FileExists(dir) {
		panic(ValueError("storage directory must exists", dir))
	}
	return &Storage{
		root:           dir,
		payloadStorage: make(map[string]map[string][]byte),
		keysStorage:    make(map[string][]string),
		diskTableMx:    NewTwoLevelMutex(),
		memTableMx:     NewTwoLevelMutex(),
	}
}
