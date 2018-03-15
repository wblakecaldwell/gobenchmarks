package maps

import (
	"fmt"
	"testing"
)

type IPV4From8s struct {
	a1 uint8
	b1 uint8
	c1 uint8
	d1 uint8
}

type IPV4PairFrom8s struct {
	a1 uint8
	b1 uint8
	c1 uint8
	d1 uint8
	a2 uint8
	b2 uint8
	c2 uint8
	d2 uint8
}

type IPV4PairFrom32s struct {
	a uint32
	b uint32
}

func buildIPV4From8s(a uint8) IPV4From8s {
	return IPV4From8s{a, a, a, a}
}

func buildIPV4PairFrom8s(a uint8) IPV4PairFrom8s {
	return IPV4PairFrom8s{a, a, a, a, a, a, a, a}
}

func buildIPV4PairFrom32s(a uint8) IPV4PairFrom32s {
	return IPV4PairFrom32s{uint32(a), uint32(a)}
}

func buildUint32(a uint8) uint32 {
	return uint32(a)
}

func buildUint64(a uint8) uint64 {
	b := uint64(buildUint32(a))
	return uint64(b)<<32 + uint64(b)
}

func buildIPv4String(a uint8) string {
	return fmt.Sprintf("%d.%d.%d.%d", a, a, a, a)
}

func buildIPv4PairString(a uint8) string {
	return fmt.Sprintf("%d.%d.%d.%d:%d.%d.%d.%d", a, a, a, a, a, a, a, a)
}

func buildIPV4Bytes(a uint8) [4]byte {
	return [4]byte{a, a, a, a}
}

func buildIPV4PairBytes(a uint8) [8]byte {
	return [8]byte{a, a, a, a, a, a, a, a}
}

func build9Bytes(a uint8) [9]byte {
	return [9]byte{
		a, a, a, a, a, a, a, a,
		a}
}

func build16Bytes(a uint8) [16]byte {
	return [16]byte{
		a, a, a, a, a, a, a, a,
		a, a, a, a, a, a, a, a}
}

func build64Bytes(a uint8) [64]byte {
	return [64]byte{
		a, a, a, a, a, a, a, a,
		a, a, a, a, a, a, a, a,
		a, a, a, a, a, a, a, a,
		a, a, a, a, a, a, a, a,
		a, a, a, a, a, a, a, a,
		a, a, a, a, a, a, a, a,
		a, a, a, a, a, a, a, a,
		a, a, a, a, a, a, a, a,
	}
}

func BenchmarkIPV4From8s(b *testing.B) {
	keys := make([]IPV4From8s, 0, 256)
	for i := 0; i < 256; i++ {
		keys = append(keys, buildIPV4From8s(uint8(i)))
	}

	m := make(map[IPV4From8s]int)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for i := 0; i < 256; i++ {
			m[keys[i]]++
		}
	}
}

func BenchmarkIPV4PairFromUint8s(b *testing.B) {
	keys := make([]IPV4PairFrom8s, 0, 256)
	for i := 0; i < 256; i++ {
		keys = append(keys, buildIPV4PairFrom8s(uint8(i)))
	}

	m := make(map[IPV4PairFrom8s]int)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for i := 0; i < 256; i++ {
			m[keys[i]]++
		}
	}
}

func BenchmarkIPV4PairFromUint32s(b *testing.B) {
	keys := make([]IPV4PairFrom32s, 0, 256)
	for i := 0; i < 256; i++ {
		keys = append(keys, buildIPV4PairFrom32s(uint8(i)))
	}

	m := make(map[IPV4PairFrom32s]int)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for i := 0; i < 256; i++ {
			m[keys[i]]++
		}
	}
}

func BenchmarkUint32(b *testing.B) {
	keys := make([]uint32, 0, 256)
	for i := 0; i < 256; i++ {
		keys = append(keys, buildUint32(uint8(i)))
	}

	m := make(map[uint32]int)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for i := 0; i < 256; i++ {
			m[keys[i]]++
		}
	}
}

func BenchmarkUint64(b *testing.B) {
	keys := make([]uint64, 0, 256)
	for i := 0; i < 256; i++ {
		keys = append(keys, buildUint64(uint8(i)))
	}

	m := make(map[uint64]int)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for i := 0; i < 256; i++ {
			m[keys[i]]++
		}
	}
}

func BenchmarkIPV4String(b *testing.B) {
	keys := make([]string, 0, 256)
	for i := 0; i < 256; i++ {
		keys = append(keys, buildIPv4String(uint8(i)))
	}

	m := make(map[string]int)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for i := 0; i < 256; i++ {
			m[keys[i]]++
		}
	}
}

func BenchmarkIPV4PairString(b *testing.B) {
	keys := make([]string, 0, 256)
	for i := 0; i < 256; i++ {
		keys = append(keys, buildIPv4PairString(uint8(i)))
	}

	m := make(map[string]int)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for i := 0; i < 256; i++ {
			m[keys[i]]++
		}
	}
}

func Benchmark4ByteArray(b *testing.B) {
	keys := make([][4]byte, 0, 256)
	for i := 0; i < 256; i++ {
		keys = append(keys, buildIPV4Bytes(uint8(i)))
	}

	m := make(map[[4]byte]int)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for i := 0; i < 256; i++ {
			m[keys[i]]++
		}
	}
}

func Benchmark8ByteArray(b *testing.B) {
	keys := make([][8]byte, 0, 256)
	for i := 0; i < 256; i++ {
		keys = append(keys, buildIPV4PairBytes(uint8(i)))
	}

	m := make(map[[8]byte]int)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for i := 0; i < 256; i++ {
			m[keys[i]]++
		}
	}
}

func BenchmarkFun9ByteArray(b *testing.B) {
	keys := make([][9]byte, 0, 256)
	for i := 0; i < 256; i++ {
		keys = append(keys, build9Bytes(uint8(i)))
	}

	m := make(map[[9]byte]int)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for i := 0; i < 256; i++ {
			m[keys[i]]++
		}
	}
}

func Benchmark16ByteArray(b *testing.B) {
	keys := make([][16]byte, 0, 256)
	for i := 0; i < 256; i++ {
		keys = append(keys, build16Bytes(uint8(i)))
	}

	m := make(map[[16]byte]int)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for i := 0; i < 256; i++ {
			m[keys[i]]++
		}
	}
}

func BenchmarkFun64ByteArray(b *testing.B) {
	keys := make([][64]byte, 0, 256)
	for i := 0; i < 256; i++ {
		keys = append(keys, build64Bytes(uint8(i)))
	}

	m := make(map[[64]byte]int)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for i := 0; i < 256; i++ {
			m[keys[i]]++
		}
	}
}
