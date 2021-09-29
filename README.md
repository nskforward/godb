# godb
High-performance embedded persistent key-value storage for Golang

- In-memory cache for reading
- Thread safe


## Benchmark

```
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkReadRandomDocument-12   7522686   156.9 ns/op   0 B/op   0 allocs/op
```

## Get started

Installation
```
go get -u github.com/nskforward/godb
```

Simple example
```
db := godb.NewStorage("storage") // will create "storage" folder near executable file

err := godb.Write("samples", "key", []byte("hello world"))
if err != nil {
    panic(err)
}

_, data, err := godb.Read("samples", "key")
if err != nil {
    panic(err)
}

fmt.Println(string(data)) // -> "hello world"

// Your data is persistent now. Try "cat storage/samples/key" and see the result.
```