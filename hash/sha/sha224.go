package sha

import (
	"encoding/binary"
)

var SHA224 = func() func([]byte) []byte {
	k := []uint32{
		0x428a2f98, 0x71374491, 0xb5c0fbcf, 0xe9b5dba5,
		0x3956c25b, 0x59f111f1, 0x923f82a4, 0xab1c5ed5,
		0xd807aa98, 0x12835b01, 0x243185be, 0x550c7dc3,
		0x72be5d74, 0x80deb1fe, 0x9bdc06a7, 0xc19bf174,
		0xe49b69c1, 0xefbe4786, 0x0fc19dc6, 0x240ca1cc,
		0x2de92c6f, 0x4a7484aa, 0x5cb0a9dc, 0x76f988da,
		0x983e5152, 0xa831c66d, 0xb00327c8, 0xbf597fc7,
		0xc6e00bf3, 0xd5a79147, 0x06ca6351, 0x14292967,
		0x27b70a85, 0x2e1b2138, 0x4d2c6dfc, 0x53380d13,
		0x650a7354, 0x766a0abb, 0x81c2c92e, 0x92722c85,
		0xa2bfe8a1, 0xa81a664b, 0xc24b8b70, 0xc76c51a3,
		0xd192e819, 0xd6990624, 0xf40e3585, 0x106aa070,
		0x19a4c116, 0x1e376c08, 0x2748774c, 0x34b0bcb5,
		0x391c0cb3, 0x4ed8aa4a, 0x5b9cca4f, 0x682e6ff3,
		0x748f82ee, 0x78a5636f, 0x84c87814, 0x8cc70208,
		0x90befffa, 0xa4506ceb, 0xbef9a3f7, 0xc67178f2,
	}
	makeW := func(group []byte) (w []uint32) {
		w = make([]uint32, 64)
		for t := 0; t < 64; t++ {
			if t < 16 {
				w[t] = binary.BigEndian.Uint32(group[t*4 : (t+1)*4])
			} else {
				w[t] = uint32Sum(uint32sigma1(w[t-2]), w[t-7], uint32sigma0(w[t-15]), w[t-16])
			}
		}
		return
	}

	return func(input []byte) []byte {
		input = uint32InputInitial(input)
		reg := newRegister32(
			0xc1059ed8, // A 0
			0x367cd507, // B 1
			0x3070dd17, // C 2
			0xf70e5939, // D 3
			0xffc00b31, // E 4
			0x68581511, // F 5
			0x64f98fa7, // G 6
			0xbefa4fa4, // H 7
		)

		l := len(input)
		for i := 0; i*64 < l; i++ {
			group := input[i*64 : (i+1)*64]
			reg2 := reg.copy()

			// init sub group
			w := makeW(group)

			for t := 0; t < 64; t++ {
				T1 := uint32Sum(reg2.v[7], uint32Ch(reg2.v[4], reg2.v[5], reg2.v[6]), uint32Sigma1(reg2.v[4]), w[t], k[t])
				T2 := uint32Sum(uint32Maj(reg2.v[0], reg2.v[1], reg2.v[2]), uint32Sigma0(reg2.v[0]))
				reg2.v[7] = reg2.v[6]
				reg2.v[6] = reg2.v[5]
				reg2.v[5] = reg2.v[4]
				reg2.v[4] = uint32Sum(reg2.v[3], T1)
				reg2.v[3] = reg2.v[2]
				reg2.v[2] = reg2.v[1]
				reg2.v[1] = reg2.v[0]
				reg2.v[0] = uint32Sum(T1, T2)

				Logger.Printf("%02d %s\n", t, reg2.toString())
			}
			reg.op(reg2, func(a uint32, b uint32) uint32 { return uint32Sum(a, b) })
			Logger.Printf("   %s\n", reg.toString())
		}
		reg.v = reg.v[:reg.l-1]
		return reg.toBytes()
	}
}()
