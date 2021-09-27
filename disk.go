package godb

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func (db *Storage) DiskCreate(bucket string, payload interface{}, labels []Label) (*Record, error) {
	if !IsNameCorrect(bucket) {
		return nil, stringValueError("incorrect bucket name", bucket)
	}
	dir := filepath.Join(db.storageDir, bucket)
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	mtx := db.DiskMutex(bucket)
	mtx.Lock()
	defer mtx.Unlock()

	err = os.MkdirAll(dir, 0755)
	if err != nil {
		return nil, err
	}
	autoinc, err := db.loadAutoIncrement(bucket)
	if err != nil {
		return nil, err
	}
	rec := Record{
		ID:      autoinc.inc(),
		Payload: payloadBytes,
	}
	err = autoinc.save()
	if err != nil {
		return nil, err
	}
	file := filepath.Join(dir, fmt.Sprintf("%d.%s", rec.ID, db.ext))
	if FileExists(file) {
		return nil, stringValueError("key bucket already exists", file)
	}
	for _, label := range labels {
		err := db.createLabel(bucket, rec.ID, label)
		if err != nil {
			return nil, err
		}
	}
	err = ioutil.WriteFile(file, rec.bytes(), 0755)
	return &rec, err
}
