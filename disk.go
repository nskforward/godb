package godb

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

func (db *Storage) diskWrite(bucket string, key string, payload []byte) error {
	dir, err := db.CreateBucket(bucket)
	if err != nil {
		return err
	}
	if !IsNameCorrect(key) {
		return ValueError("bucket key contains not allowed characters", key)
	}
	mtx := db.diskTableMx.Get(bucket)
	mtx.Lock()
	defer mtx.Unlock()

	return ioutil.WriteFile(filepath.Join(dir, key), payload, 0755)
}

func (db *Storage) diskRead(bucket string, key string) ([]byte, error) {
	if !IsNameCorrect(bucket) {
		return nil, ValueError("bucket name contains not allowed characters", bucket)
	}
	if !IsNameCorrect(key) {
		return nil, ValueError("bucket key contains not allowed characters", key)
	}
	mtx := db.diskTableMx.Get(bucket)
	mtx.Lock()
	defer mtx.Unlock()
	file := filepath.Join(db.root, bucket, key)
	if !FileExists(file) {
		return nil, nil
	}
	return ioutil.ReadFile(file)
}

func (db *Storage) diskKeys(bucket string) ([]string, error) {
	if !IsNameCorrect(bucket) {
		return nil, ValueError("bucket name contains not allowed characters", bucket)
	}
	dir := db.GetBucketDir(bucket)
	if !FileExists(dir) {
		return []string{}, nil
	}
	mtx := db.diskTableMx.Get(bucket)
	mtx.Lock()
	defer mtx.Unlock()

	arr, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	list := make([]string, 0, 64)
	for _, f := range arr {
		if f.IsDir() {
			continue
		}
		if f.Name() == ".autoinc" {
			continue
		}
		list = append(list, f.Name())
	}
	return list, nil
}

func (db *Storage) diskRemove(bucket string, key string) error {
	if !IsNameCorrect(bucket) {
		return ValueError("bucket name contains not allowed characters", bucket)
	}
	if !IsNameCorrect(key) {
		return ValueError("bucket key contains not allowed characters", key)
	}
	mtx := db.diskTableMx.Get(bucket)
	mtx.Lock()
	defer mtx.Unlock()
	file := filepath.Join(db.root, bucket, key)
	if !FileExists(file) {
		return nil
	}
	return os.Remove(file)
}

func (db *Storage) diskRemoveAll(bucket string) error {
	if !IsNameCorrect(bucket) {
		return ValueError("bucket name contains not allowed characters", bucket)
	}
	mtx := db.diskTableMx.Get(bucket)
	mtx.Lock()
	defer mtx.Unlock()
	dir := filepath.Join(db.root, bucket)
	return os.RemoveAll(dir)
}

func (db *Storage) diskKeyExists(bucket string, key string) bool {
	file := filepath.Join(db.root, bucket, key)
	return FileExists(file)
}
