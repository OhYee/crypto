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
	rawLength := len(input) * 8

	// Ensure l≡448 mod 512
	input = append(input, 0x80)
	for (len(input)*8)%512 != 448 {
		input = append(input, 0b00000000)
	}

	// Add the length at the end of input
	length := make([]byte, 8)
	binary.BigEndian.PutUint64(length, uint64(rawLength))
	input = append(input, length...)

	// init register
	reg := newRegister()

	// split by group
	l := len(input)
	for i := 0; i*64 < l; i++ {
		reg2 := reg.copy()

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
			temp := reg2.copy()

			reg2.A = sum(bits.RotateLeft32(temp.A, 5), f(temp.B, temp.C, temp.D, t), temp.E, w[t], getK(t))
			reg2.B = temp.A
			reg2.C = bits.RotateLeft32(temp.B, 30)
			reg2.D = temp.C
			reg2.E = temp.D
			Logger.Printf("%d 0x%08x 0x%08x 0x%08x 0x%08x 0x%08x\n", t, reg2.A, reg2.B, reg2.C, reg2.D, reg2.E)
		}

		reg.A = sum(reg.A, reg2.A)
		reg.B = sum(reg.B, reg2.B)
		reg.C = sum(reg.C, reg2.C)
		reg.D = sum(reg.D, reg2.D)
		reg.E = sum(reg.E, reg2.E)
		Logger.Printf("%d 0x%08x 0x%08x 0x%08x 0x%08x 0x%08x\n", 80, reg.A, reg.B, reg.C, reg.D, reg.E)
	}

	// transfer from []uint32 to []byte
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
