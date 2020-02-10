# io bench

`go test -bench . -benchmem`

# result

```
goos: darwin
goarch: amd64
pkg: github.com/sueken5/io-bench
BenchmarkEx1-8   	   88365	     12195 ns/op	    2392 B/op	      10 allocs/op
BenchmarkEx2-8   	1000000000	         0.000065 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/sueken5/io-bench	1.224s
```