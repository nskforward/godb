package tests

import (
	"bytes"
	"strconv"
	"strings"
	"testing"

	"github.com/nskforward/godb"
)

func TestKeys(t *testing.T) {
	storageRoot := "/Users/a17847869/go/src/github.com/nskforward/godb/tests/tmp"
	db := godb.NewStorage(storageRoot)
	err := db.RemoveAll("samples")
	if err != nil {
		t.Fatalf("fail: %s", err)
	}
	err = db.Write("samples", "1", []byte("1"))
	if err != nil {
		t.Fatalf("fail: %s", err)
	}
	err = db.Write("samples", "2", []byte("2"))
	if err != nil {
		t.Fatalf("fail: %s", err)
	}

	keys, err := db.Keys("samples")
	if err != nil {
		t.Fatalf("fail: %s", err)
	}
	if strings.Join(keys, "") != "12" {
		t.Fatalf("fail: %s, actual:'%v'", err, keys)
	}

	err = db.Write("samples", "3", []byte("3"))
	if err != nil {
		t.Fatalf("fail: %s", err)
	}

	keys, err = db.Keys("samples")
	if err != nil {
		t.Fatalf("fail: %s", err)
	}
	if strings.Join(keys, "") != "123" {
		t.Fatalf("fail: %s, actual:'%v'", err, keys)
	}

	keys, err = db.Keys("samples")
	if err != nil {
		t.Fatalf("fail: %s", err)
	}
	if strings.Join(keys, "") != "123" {
		t.Fatalf("fail: %s, actual:'%v'", err, keys)
	}

	err = db.Remove("samples", "2")
	if err != nil {
		t.Fatalf("fail: %s", err)
	}
	keys, err = db.Keys("samples")
	if err != nil {
		t.Fatalf("fail: %s", err)
	}
	if strings.Join(keys, "") != "13" {
		t.Fatalf("fail: %s, actual:'%v'", err, keys)
	}
}

func TestAutoincrement(t *testing.T) {
	storageRoot := "/Users/a17847869/go/src/github.com/nskforward/godb/tests/tmp"
	db := godb.NewStorage(storageRoot)
	err := db.RemoveAll("samples")
	if err != nil {
		t.Fatalf("fail: %s", err)
	}
	key, err := db.Autoincrement("samples")
	if err != nil {
		t.Fatalf("fail: %s", err)
	}
	err = db.Write("samples", strconv.FormatInt(key, 10), []byte("1"))
	if err != nil {
		t.Fatalf("fail: %s", err)
	}
	data, err := db.Read("samples", "1")
	if err != nil {
		t.Fatalf("fail: %s", err)
	}
	if !bytes.Equal(data, []byte("1")) {
		t.Fatalf("fail: bytes are not the same")
	}
	key, err = db.Autoincrement("samples")
	if err != nil {
		t.Fatalf("fail: %s", err)
	}
	err = db.Write("samples", strconv.FormatInt(key, 10), []byte("2"))
	if err != nil {
		t.Fatalf("fail: %s", err)
	}
	data, err = db.Read("samples", "2")
	if err != nil {
		t.Fatalf("fail: %s", err)
	}
	if !bytes.Equal(data, []byte("2")) {
		t.Fatalf("fail: bytes are not the same")
	}
}

func TestReadAll(t *testing.T) {
	storageRoot := "/Users/a17847869/go/src/github.com/nskforward/godb/tests/tmp"
	db := godb.NewStorage(storageRoot)
	err := db.RemoveAll("samples")
	if err != nil {
		t.Fatalf("fail: %s", err)
	}
	err = db.Write("samples", "1", []byte("1"))
	if err != nil {
		t.Fatalf("fail: %s", err)
	}
	err = db.Write("samples", "2", []byte("2"))
	if err != nil {
		t.Fatalf("fail: %s", err)
	}
	err = db.Write("samples", "3", []byte("3"))
	if err != nil {
		t.Fatalf("fail: %s", err)
	}
	data, err := db.ReadAll("samples")
	if err != nil {
		t.Fatalf("fail: %s", err)
	}
	var buf bytes.Buffer
	for _, item := range data {
		buf.Write(item)
	}
	if !bytes.Equal(buf.Bytes(), []byte("123")) {
		t.Fatalf("fail: bytes are not the same")
	}
}
