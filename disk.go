package godb

import (
	"io/ioutil"
	"path/filepath"
	"strings"
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

func (db *Storage) diskRead(bucket string, key string) ([]byte, error) {
	mtx := db.diskTableMx.Get(bucket)
	mtx.Lock()
	defer mtx.Unlock()
	file := filepath.Join(db.root, bucket, key)
	if !FileExists(file) {
		return nil, nil
	}
	return ioutil.ReadFile(file)
}

func (db *Storage) Keys(bucket string) ([]string, error) {
	mtx := db.diskTableMx.Get(bucket)
	mtx.Lock()
	defer mtx.Unlock()

	arr, err := ioutil.ReadDir(db.GetBucketDir(bucket))
	if err != nil {
		return nil, err
	}
	list := make([]string, 0, 64)
	for _, f := range arr {
		if f.IsDir() {
			continue
		}
		names := strings.Split(f.Name(), ".")
		if len(names) != 2 {
			continue
		}
		if names[1] != "json" {
			continue
		}
		list = append(list, names[0])
	}
	return list, nil
}
