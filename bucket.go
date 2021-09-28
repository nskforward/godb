package godb

import (
	"os"
	"path/filepath"
)

func (db *Storage) CreateBucket(bucket string) (string, error) {
	if !IsNameCorrect(bucket) {
		return "", ValueError("bucket name has not allowed characters", bucket)
	}
	dir := filepath.Join(db.root, bucket)
	return dir, os.MkdirAll(dir, 0755)
}

func (db *Storage) GetBucketDir(bucket string) string {
	return filepath.Join(db.root, bucket)
}
