benchmarking different types of map keys:

		$ go version
		go version go1.9.1 linux/amd64

		$ go test --bench=. .
		goos: linux
		goarch: amd64
		pkg: github.com/wblakecaldwell/gobenchmarks/maps
		BenchmarkIPV4-4             	  100000	     12270 ns/op		# IPv4 as struct with 4 uint8's
		BenchmarkIPV4Pair-4         	  200000	      7758 ns/op		# IPv4 pair as struct with 8 uint8's
		BenchmarkUint32-4           	  200000	      7955 ns/op		# IPv4 as uint32
		BenchmarkUint64-4           	  200000	      7849 ns/op		# IPv4 pair as uint64
		BenchmarkIPV4String-4       	  200000	     10578 ns/op		# IPv4 as string
		BenchmarkIPV4PairString-4   	  100000	     11299 ns/op		# IPv4 pair as string with colon separator
		Benchmark4ByteArray-4       	  200000	      8013 ns/op		# IPv4 as [4]byte
		Benchmark8ByteArray-4       	  200000	      7892 ns/op		# IPv4 pair as [8]byte
		Benchmark9ByteArray-4       	  100000	     19113 ns/op		# just for fun: [9]byte
		Benchmark16ByteArray-4      	  100000	     16565 ns/op		# just for fun: [16]byte
		Benchmark64ByteArray-4      	  100000	     22081 ns/op		# just for fun: [64]byte
		Benchmark16IPString-4       	  100000	     16629 ns/op		# just for fun: string with 16 IP addresses and colon separator
		PASS
		ok  	github.com/wblakecaldwell/gobenchmarks/maps	21.438s	

Go prefers 4 and 8 byte integers and byte arrays. If you can't use those, consider a short string.
