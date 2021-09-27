package tests

import (
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/nskforward/godb"
)

func TestDiskCreate(t *testing.T) {
	db := godb.NewStorage("storage", "json")
	_, err := db.DiskCreate(
		"samples",
		[]string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"},
		[]godb.Label{
			{Name: "lang", Value: "eng"},
			{Name: "format", Value: "short"},
		},
	)
	if err != nil {
		t.Fatalf(`failed: %v`, err)
	}
	_, err = ioutil.ReadFile(filepath.Join(godb.ProcessDir(), "storage", "samples", "1.json"))
	if err != nil {
		t.Fatalf(`failed: %v`, err)
	}
}
