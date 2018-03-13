benchmarking different types of map keys: 

- struct with 4 uint8's
- uint32
- string representation of IPv4
- [4]byte

uint32 and byte array win

		$ go version
		go version go1.9.1 linux/amd64
		
		$ go test --bench=. .
		goos: linux
		goarch: amd64
		pkg: github.com/wblakecaldwell/gobenchmarks/maps
		BenchmarkStruct-4      	  100000	     12515 ns/op
		BenchmarkUint32-4      	  200000	      7914 ns/op
		BenchmarkString-4      	  100000	     11193 ns/op
		BenchmarkByteArray-4   	  200000	      7985 ns/op
		PASS
		ok  	github.com/wblakecaldwell/gobenchmarks/maps	5.998s
