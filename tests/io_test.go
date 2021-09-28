package tests

import (
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
