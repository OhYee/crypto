package sha

import (
	"math/bits"
)

func uint64Sum(nums ...uint64) (res uint64) {
	for _, value := range nums {
		res = res + value
	}
	return
}

func uint64Ch(x uint64, y uint64, z uint64) uint64 {
	return (x & y) ^ (^x & z)
}

func uint64Maj(x uint64, y uint64, z uint64) uint64 {
	return (x & y) ^ (x & z) ^ (y & z)
}

func uint64Sigma0(x uint64) uint64 {
	return bits.RotateLeft64(x, 64-28) ^ bits.RotateLeft64(x, 64-34) ^ bits.RotateLeft64(x, 64-39)
}

func uint64Sigma1(x uint64) uint64 {
	return bits.RotateLeft64(x, 64-14) ^ bits.RotateLeft64(x, 64-18) ^ bits.RotateLeft64(x, 64-41)
}

func uint64sigma0(x uint64) uint64 {
	return bits.RotateLeft64(x, 64-1) ^ bits.RotateLeft64(x, 64-8) ^ (x >> 7)
}

func uint64sigma1(x uint64) uint64 {
	return bits.RotateLeft64(x, 64-19) ^ bits.RotateLeft64(x, 64-61) ^ (x >> 6)
}
