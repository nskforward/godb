package godb

import (
	"io/ioutil"
	"path/filepath"
	"strconv"
)

func (db *Storage) Autoincrement(bucket string) (int64, error) {
	mtx := db.diskTableMx.Get(bucket)
	mtx.Lock()
	defer mtx.Unlock()
	dir, err := db.CreateBucket(bucket)
	if err != nil {
		return 0, err
	}
	file := filepath.Join(dir, ".autoinc")
	if !FileExists(file) {
		err := ioutil.WriteFile(file, []byte("1"), 0755)
		return 1, err
	}
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return 0, err
	}
	i, err := strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		return 0, err
	}
	i++
	return i, ioutil.WriteFile(file, []byte(strconv.FormatInt(i, 10)), 0755)
}
