package tests

import (
	"math/rand"
	"sync/atomic"
	"testing"

	"github.com/nskforward/godb"
)

/*
	cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
	BenchmarkWriteInsert-12    	    9092	    189235 ns/op	     639 B/op	       8 allocs/op
*/

func BenchmarkWriteInsert(b *testing.B) {
	storageRoot := "/Users/a17847869/go/src/github.com/nskforward/godb/tests/tmp"
	db := godb.NewStorage(storageRoot)
	payload := []byte(`{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}`)
	var counter int64
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			atomic.AddInt64(&counter, 1)
			err := db.Write("samples", counter, payload)
			if err != nil {
				b.Fatalf(`failed: %v`, err)
			}
		}
	})
}

/*
	cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
	BenchmarkDiskWriteUpdate-12    	    6214	    196626 ns/op	     642 B/op	       7 allocs/op
*/

func BenchmarkWriteUpdate(b *testing.B) {
	storageRoot := "/Users/a17847869/go/src/github.com/nskforward/godb/tests/tmp"
	db := godb.NewStorage(storageRoot)
	payload := []byte(`{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}`)
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			key := rand.Int63n(1000) + 1
			err := db.Write("samples", key, payload)
			if err != nil {
				b.Fatalf(`failed: %v`, err)
			}
		}
	})
}

/*
	Memory
	cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
	BenchmarkCache-12    	 7522686	       156.9 ns/op	       0 B/op	       0 allocs/op

	Disk
	cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
	BenchmarkRead-12    	   48409	     24107 ns/op	    1274 B/op	       8 allocs/op
*/
func BenchmarkCache(b *testing.B) {
	storageRoot := "/Users/a17847869/go/src/github.com/nskforward/godb/tests/tmp"
	db := godb.NewStorage(storageRoot)
	_, _, err := db.Read("samples", 10)
	if err != nil {
		b.Fatalf(`failed: %v`, err)
	}
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			key := rand.Int63n(100) + 1
			_, _, err := db.Read("samples", key)
			if err != nil {
				b.Fatalf(`failed: %v`, err)
			}
		}
	})
}

func TestReadCache(t *testing.T) {
	storageRoot := "/Users/a17847869/go/src/github.com/nskforward/godb/tests/tmp"
	db := godb.NewStorage(storageRoot)
	err := db.Write("samples", 10, []byte("1"))
	if err != nil {
		t.Fatalf("fail: %s", err)
	}
	cache, _, err := db.Read("samples", 10)
	if err != nil {
		t.Fatalf("fail: %s", err)
	}
	if cache {
		t.Fatalf("cannot be cache")
	}
	cache, _, err = db.Read("samples", 10)
	if err != nil {
		t.Fatalf("fail: %s", err)
	}
	if !cache {
		t.Fatalf("must be cache")
	}
}

func TestWriteCache(t *testing.T) {
	storageRoot := "/Users/a17847869/go/src/github.com/nskforward/godb/tests/tmp"
	db := godb.NewStorage(storageRoot)
	err := db.Write("samples", 10, []byte("1"))
	if err != nil {
		t.Fatalf("fail: %s", err)
	}
	cache, _, err := db.Read("samples", 10)
	if err != nil {
		t.Fatalf("fail: %s", err)
	}
	if cache {
		t.Fatalf("cannot be cache")
	}
	err = db.Write("samples", 10, []byte("1"))
	if err != nil {
		t.Fatalf("fail: %s", err)
	}
	cache, _, err = db.Read("samples", 10)
	if err != nil {
		t.Fatalf("fail: %s", err)
	}
	if cache {
		t.Fatalf("cannot be cache")
	}
}
