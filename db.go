package godb

import (
	"os"
	"path/filepath"
	"sync"
)

type Storage struct {
	storageDir        string
	_diskMutex        sync.Mutex
	diskBucketMutexes map[string]*sync.Mutex
	_memMutex         sync.Mutex
	memBucketMutexes  map[string]*sync.Mutex
	memory            map[string]map[string][]byte
	ext               string
}

func NewStorage(storageDir, extention string) *Storage {
	storageDir = filepath.Join(ProcessDir(), storageDir)
	err := os.MkdirAll(storageDir, 0755)
	if err != nil {
		panic(err)
	}
	return &Storage{
		storageDir:        storageDir,
		diskBucketMutexes: make(map[string]*sync.Mutex),
		memory:            make(map[string]map[string][]byte),
		memBucketMutexes:  make(map[string]*sync.Mutex),
		ext:               extention,
	}
}

func (db *Storage) DiskMutex(bucket string) *sync.Mutex {
	db._diskMutex.Lock()
	defer db._diskMutex.Unlock()
	mtx, ok := db.diskBucketMutexes[bucket]
	if !ok {
		mtx = &sync.Mutex{}
		db.diskBucketMutexes[bucket] = mtx
		return mtx
	}
	return mtx
}

func (db *Storage) MemMutex(bucket string) *sync.Mutex {
	db._memMutex.Lock()
	defer db._memMutex.Unlock()
	mtx, ok := db.memBucketMutexes[bucket]
	if !ok {
		mtx = &sync.Mutex{}
		db.memBucketMutexes[bucket] = mtx
		return mtx
	}
	return mtx
}
