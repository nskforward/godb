package godb

import (
	"io/ioutil"
	"path/filepath"
	"strconv"
)

type AutoIncrement struct {
	value int64
	path  string
}

func (db *Storage) loadAutoIncrement(bucket string) (AutoIncrement, error) {
	file := filepath.Join(db.storageDir, bucket, "autoinc")
	if !FileExists(file) {
		return AutoIncrement{0, file}, nil
	}
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return AutoIncrement{}, err
	}
	i, err := strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		return AutoIncrement{}, err
	}
	return AutoIncrement{i, file}, nil
}

func (autoinc *AutoIncrement) inc() int64 {
	autoinc.value++
	return autoinc.value
}

func (autoinc *AutoIncrement) save() error {
	s := strconv.FormatInt(autoinc.value, 10)
	return ioutil.WriteFile(autoinc.path, []byte(s), 0755)
}
