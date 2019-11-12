package sha

import (
	"math/bits"
)

func uint32Sum(nums ...uint32) uint32 {
	var res uint64
	for _, value := range nums {
		res = (res + uint64(value)) % uint64(1 << 32)
	}
	return uint32(res)
}

func uint32Ch(x uint32, y uint32, z uint32) uint32 {
	return (x & y) ^ (^x & z)
}

func uint32Maj(x uint32, y uint32, z uint32) uint32 {
	return (x & y) ^ (x & z) ^ (y & z)
}

func uint32Sigma0(x uint32) uint32 {
	return bits.RotateLeft32(x, 32-2) ^ bits.RotateLeft32(x, 32-13) ^ bits.RotateLeft32(x, 32-22)
}

func uint32Sigma1(x uint32) uint32 {
	return bits.RotateLeft32(x, 32-6) ^ bits.RotateLeft32(x, 32-11) ^ bits.RotateLeft32(x, 32-25)
}

func uint32sigma0(x uint32) uint32 {
	return bits.RotateLeft32(x, 32-7) ^ bits.RotateLeft32(x, 32-18) ^ (x >> 3)
}

func uint32sigma1(x uint32) uint32 {
	return bits.RotateLeft32(x, 32-17) ^ bits.RotateLeft32(x, 32-19) ^ (x >> 10)
}
