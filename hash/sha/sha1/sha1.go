package sha

import (
	"bytes"
	"encoding/binary"
	"github.com/OhYee/cryptography_and_network_security/util/blackhole"
	"github.com/OhYee/goutils/functional"
	"log"
	"math/bits"
)

// Logger 日志
var Logger = log.New(blockhole.BlackHole{}, "", 0)

const (
	initA = 0x67452301
	initB = 0xEFCDAB89
	initC = 0x98BADCFE
	initD = 0x10325476
	initE = 0xC3D2E1F0
)

type register struct {
	A uint32
	B uint32
	C uint32
	D uint32
	E uint32
}

func newRegister() *register {
	return &register{
		A: initA,
		B: initB,
		C: initC,
		D: initD,
		E: initE,
	}
}

func (reg *register) copy() *register {
	return &register{
		A: reg.A,
		B: reg.B,
		C: reg.C,
		D: reg.D,
		E: reg.E,
	}
}

func SHA1(input []byte) []byte {
	// Ensure l≡448 mod 512
	first := true
	for (len(input)*8)%512 != 448 {
		if first {
			input = append(input, 0b10000000)
			first = false
		} else {
			input = append(input, 0b00000000)
		}
	}

	// Add the length at the end of input
	length := make([]byte, 8)
	l := len(input)
	binary.BigEndian.PutUint64(length, uint64(l))
	input = append(input, length...)

	// init register
	reg := newRegister()

	// split by group
	l = len(input)
	for i := 0; i*64 < l; i++ {
		group := input[i*64 : (i+1)*64]

		// init sub group
		w := make([]uint32, 80)
		for t := 0; t < 80; t++ {
			if t < 16 {
				w[t] = binary.BigEndian.Uint32(group[t*4 : (t+1)*4])
			} else {
				w[t] = bits.RotateLeft32(w[t-3]^w[t-8]^w[t-14]^w[t-16], 1)
			}
		}

		for t := 0; t < 80; t++ {
			oldReg := reg.copy()

			reg.A = sum(bits.RotateLeft32(oldReg.A, 5), f(oldReg.B, oldReg.C, oldReg.D, t), oldReg.E, w[t], getK(t))
			reg.B = oldReg.A
			reg.C = bits.RotateLeft32(oldReg.B, 30)
			reg.D = oldReg.C
			reg.E = oldReg.D
			Logger.Printf("%d 0x%08x 0x%08x 0x%08x 0x%08x 0x%08x\n", t, reg.A, reg.B, reg.C, reg.D, reg.E)
		}

		reg.A = sum(reg.A, initA)
		reg.B = sum(reg.B, initB)
		reg.C = sum(reg.C, initC)
		reg.D = sum(reg.D, initD)
		reg.E = sum(reg.E, initE)
		Logger.Printf("%d 0x%08x 0x%08x 0x%08x 0x%08x 0x%08x\n", 80, reg.A, reg.B, reg.C, reg.D, reg.E)
	}

	buf := bytes.NewBuffer([]byte{})
	temp := make([]byte, 4)
	fp.MapUint32(func(data uint32) uint32 {
		binary.BigEndian.PutUint32(temp, data)
		buf.Write(temp)
		return 0
	}, []uint32{reg.A, reg.B, reg.C, reg.D, reg.E})
	return buf.Bytes()
}

func getK(t int) (k uint32) {
	switch {
	case t <= 19:
		k = 0x5A827999
	case t >= 20 && t <= 39:
		k = 0x6ED9EBA1
	case t >= 40 && t <= 59:
		k = 0x8F1BBCDC
	case t >= 60 && t <= 79:
		k = 0xCA62C1D6
	}
	return
}

func sum(nums ...uint32) uint32 {
	var res uint64
	for _, value := range nums {
		res += uint64(value) % uint64((1 << 32))
	}
	return uint32(res)
}

func f(x uint32, y uint32, z uint32, t int) (ret uint32) {
	switch {
	case t <= 19:
		ret = (x & y) ^ ((^x) & z)
	case t >= 20 && t <= 39, t >= 60 && t <= 79:
		ret = x ^ y ^ z
	case t >= 40 && t <= 59:
		ret = (x & y) ^ (x & z) ^ (y & z)
	}
	return
}
