package godb

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Label struct {
	Name  string
	Value string
}

type LabelIDs []int64

func (labelIDs *LabelIDs) encode() []byte {
	data, err := json.Marshal(labelIDs)
	if err != nil {
		panic(err)
	}
	return data
}

func (db *Storage) createLabel(bucket string, id int64, label Label) error {
	dir := filepath.Join(db.storageDir, bucket, "index", label.Name)
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return err
	}
	file := filepath.Join(dir, fmt.Sprintf("%s_%s", label.Name, label.Value))
	if !FileExists(file) {
		arr := LabelIDs{id}
		return ioutil.WriteFile(file, arr.encode(), 0755)
	}
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	var arr LabelIDs
	err = json.Unmarshal(data, &arr)
	if err != nil {
		return err
	}
	arr = append(arr, id)
	return ioutil.WriteFile(file, arr.encode(), 0755)
}
