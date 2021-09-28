package godb

import (
	"io/ioutil"
	"path/filepath"
)

func (db *Storage) diskWrite(bucket string, key int64, payload []byte) error {
	dir, err := db.CreateBucket(bucket)
	if err != nil {
		return err
	}
	mtx := db.diskTableMx.Get(bucket)
	mtx.Lock()
	defer mtx.Unlock()
	return ioutil.WriteFile(filepath.Join(dir, Int64ToStr(key)), payload, 0755)
}

func (db *Storage) DiskRead(bucket string, key int64) ([]byte, error) {
	mtx := db.diskTableMx.Get(bucket)
	mtx.Lock()
	defer mtx.Unlock()
	file := filepath.Join(db.root, bucket, Int64ToStr(key))
	if !FileExists(file) {
		return nil, nil
	}
	return ioutil.ReadFile(file)
}
