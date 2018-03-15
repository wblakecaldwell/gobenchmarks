Benchmarking different types of map keys:

	$ go version
	go version go1.9.1 linux/amd64

	$ go test --bench=. .
	goos: linux
	goarch: amd64
	pkg: github.com/wblakecaldwell/gobenchmarks/maps
	BenchmarkIPV4From8s-4            	  100000	     13149 ns/op
	BenchmarkIPV4PairFromUint8s-4    	  200000	      8081 ns/op
	BenchmarkIPV4PairFromUint32s-4   	  100000	     11809 ns/op
	BenchmarkUint32-4                	  200000	      7689 ns/op
	BenchmarkUint64-4                	  200000	      7846 ns/op
	BenchmarkIPV4String-4            	  100000	     11026 ns/op
	BenchmarkIPV4PairString-4        	  200000	     10855 ns/op
	Benchmark4ByteArray-4            	  200000	      7537 ns/op
	Benchmark8ByteArray-4            	  200000	      7524 ns/op
	BenchmarkFun9ByteArray-4         	  100000	     18160 ns/op
	Benchmark16ByteArray-4           	  100000	     16653 ns/op
	BenchmarkFun64ByteArray-4        	  100000	     21882 ns/op
	PASS
	ok  	github.com/wblakecaldwell/gobenchmarks/maps	20.739s


##Interpretation:##

**Single IPv4:**

	Benchmark4ByteArray-4            	  200000	      7537 ns/op   # 4 byte array
	BenchmarkUint32-4                	  200000	      7689 ns/op   # uint32
	BenchmarkIPV4String-4            	  100000	     11026 ns/op   # "1.2.3.4"
	BenchmarkIPV4From8s-4            	  100000	     13149 ns/op   # struct with 4 uint8 members

**Pair of IPv4:**

	Benchmark8ByteArray-4            	  200000	      7524 ns/op   # [8]byte
	BenchmarkUint64-4                	  200000	      7846 ns/op   # uint64
	BenchmarkIPV4PairFromUint8s-4    	  200000	      8081 ns/op   # struct with 8 uint8 members
	BenchmarkIPV4PairString-4        	  200000	     10855 ns/op   # "1.2.3.4:5.6.7.8"
	BenchmarkIPV4PairFromUint32s-4   	  100000	     11809 ns/op   # struct with 2 uint32 members
	
**Single IPv6:**

	Benchmark8ByteArray-4            	  200000	      7524 ns/op   # [8]byte
	BenchmarkUint64-4                	  200000	      7846 ns/op   # uint64
	BenchmarkIPV4PairFromUint8s-4    	  200000	      8081 ns/op   # struct with 8 uint8 members
	BenchmarkIPV4PairFromUint32s-4   	  100000	     11809 ns/op   # struct with 2 uint32 members

**Pair of IPv6:**

	Benchmark16ByteArray-4           	  100000	     16653 ns/op   # [16]byte

**Just for fun:**

	Benchmark8ByteArray-4            	  200000	      7524 ns/op   # [8]byte
	Benchmark16ByteArray-4           	  100000	     16653 ns/op   # [16]byte
	BenchmarkFun9ByteArray-4         	  100000	     18160 ns/op   # [9]byte - look how much slower than [8]byte - and even [16]byte
	BenchmarkFun64ByteArray-4        	  100000	     21882 ns/op   # [64]byte


Go prefers 4 and 8 byte integers and byte arrays. If you can't use those, consider a short string.
