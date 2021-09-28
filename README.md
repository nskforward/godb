# godb
Embedded persistent storage for Golang

- Document oriented key-value storage
- In-memory cache for reading
- Two level mutexes (thread safe)


## Benchmark

cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkReadRandomDocument-12    	 7522686	       156.9 ns/op	       0 B/op	       0 allocs/op