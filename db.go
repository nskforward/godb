package godb

type Storage struct {
	root        string
	memStorage  map[string]map[int64][]byte
	diskTableMx *TwoLevelMutex
	memTableMx  *TwoLevelMutex
}

func NewStorage(dir string) *Storage {
	if !FileExists(dir) {
		panic(ValueError("storage directory must exists", dir))
	}
	return &Storage{
		root:        dir,
		memStorage:  make(map[string]map[int64][]byte),
		diskTableMx: NewTwoLevelMutex(),
		memTableMx:  NewTwoLevelMutex(),
	}
}
