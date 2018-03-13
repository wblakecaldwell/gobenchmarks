package maps

import (
	"fmt"
	"testing"
)

// IPV4 is an IPv4 as 4 uint8's
type IPV4 struct {
	a uint8
	b uint8
	c uint8
	d uint8
}

func buildStruct(a uint8) IPV4 {
	return IPV4{a, a, a, a}
}

func buildUint32(a uint8) uint32 {
	return uint32(a)<<24 + uint32(a)<<16 + uint32(a)<<8 + uint32(a)
}

func buildString(a uint8) string {
	return fmt.Sprintf("%d.%d.%d.%d.", a, a, a, a)
}

func buildByteArray(a uint8) [4]byte {
	return [4]byte{a, a, a, a}
}

func BenchmarkStruct(b *testing.B) {
	keys := make([]IPV4, 0, 256)
	for i := 0; i < 256; i++ {
		keys = append(keys, buildStruct(uint8(i)))
	}

	m := make(map[IPV4]int)
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

func BenchmarkString(b *testing.B) {
	keys := make([]string, 0, 256)
	for i := 0; i < 256; i++ {
		keys = append(keys, buildString(uint8(i)))
	}

	m := make(map[string]int)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for i := 0; i < 256; i++ {
			m[keys[i]]++
		}
	}
}

func BenchmarkByteArray(b *testing.B) {
	keys := make([][4]byte, 0, 256)
	for i := 0; i < 256; i++ {
		keys = append(keys, buildByteArray(uint8(i)))
	}

	m := make(map[[4]byte]int)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for i := 0; i < 256; i++ {
			m[keys[i]]++
		}
	}
}
