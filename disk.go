package godb

import (
	"io/ioutil"
	"path/filepath"
)

func (db *Storage) diskWrite(bucket string, key string, payload []byte) error {
	dir, err := db.CreateBucket(bucket)
	if err != nil {
		return err
	}
	mtx := db.diskTableMx.Get(bucket)
	mtx.Lock()
	defer mtx.Unlock()
	return ioutil.WriteFile(filepath.Join(dir, key), payload, 0755)
}

func (db *Storage) DiskRead(bucket string, key string) ([]byte, error) {
	mtx := db.diskTableMx.Get(bucket)
	mtx.Lock()
	defer mtx.Unlock()
	file := filepath.Join(db.root, bucket, key)
	if !FileExists(file) {
		return nil, nil
	}
	return ioutil.ReadFile(file)
}
