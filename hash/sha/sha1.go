package sha

import (
	"encoding/binary"
	"math/bits"
)

var SHA1 = func() func([]byte) []byte {
	getK := func(t int) (k uint32) {
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
	makeW := func(group []byte) (w []uint32) {
		w = make([]uint32, 80)
		for t := 0; t < 80; t++ {
			if t < 16 {
				w[t] = binary.BigEndian.Uint32(group[t*4 : (t+1)*4])
			} else {
				w[t] = bits.RotateLeft32(w[t-3]^w[t-8]^w[t-14]^w[t-16], 1)
			}
		}
		return
	}

	f := func(x uint32, y uint32, z uint32, t int) (ret uint32) {
		switch {
		case t <= 19:
			ret = uint32Ch(x, y, z)
		case t >= 20 && t <= 39, t >= 60 && t <= 79:
			ret = x ^ y ^ z
		case t >= 40 && t <= 59:
			ret = uint32Maj(x, y, z)
		}
		return
	}

	return func(input []byte) []byte {
		input = uint32InputInitial(input)
		reg := newRegister32(
			0x67452301, // A 0
			0xEFCDAB89, // B 1
			0x98BADCFE, // C 2
			0x10325476, // D 3
			0xC3D2E1F0, // E 4
		)

		l := len(input)
		for i := 0; i*64 < l; i++ {
			group := input[i*64 : (i+1)*64]
			reg2 := reg.copy()

			// init sub group
			w := makeW(group)

			for t := 0; t < 80; t++ {
				T1 := uint32Sum(bits.RotateLeft32(reg2.v[0], 5), f(reg2.v[1], reg2.v[2], reg2.v[3], t), reg2.v[4], w[t], getK(t))
				reg2.v[4] = reg2.v[3]
				reg2.v[3] = reg2.v[2]
				reg2.v[2] = bits.RotateLeft32(reg2.v[1], 30)
				reg2.v[1] = reg2.v[0]
				reg2.v[0] = T1

				Logger.Printf("%02d %s\n", t, reg2.toString())
			}
			reg.op(reg2, func(a uint32, b uint32) uint32 { return uint32Sum(a, b) })
			Logger.Printf("   %s\n", reg.toString())
		}
		return reg.toBytes()
	}
}()
