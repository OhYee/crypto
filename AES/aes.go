package aes

import (
	"fmt"
	"github.com/OhYee/rainbow/color"
	"github.com/OhYee/rainbow/log"

	"github.com/OhYee/crypto/AES/bits"
	gf "github.com/OhYee/crypto/GF"
)

// Logger 日志
var Logger = log.New().SetOutputToNil().SetPrefix(func(s string) string {
	return color.New().SetFontBold().Colorful("Log     ")
})

var (
	sBoxTable        = generateSBox()
	rowTransferTable = [][]byte{
		{2, 3, 1, 1},
		{1, 2, 3, 1},
		{1, 1, 2, 3},
		{3, 1, 1, 2},
	}
	keyGenerateTable = []uint32{
		0x01000000, 0x02000000, 0x04000000, 0x08000000, 0x10000000,
		0x20000000, 0x40000000, 0x80000000, 0x1B000000, 0x36000000,
	}
)

func generateSBox() [][]byte {
	table := make([][]byte, 16)
	for i := 0; i < 16; i++ {
		table[i] = make([]byte, 16)
		for j := 0; j < 16; j++ {
			b := byte(gf.Inverse(i*16+j, 0b100011011))
			b2 := byte(0)
			for k := 0; k < 8; k++ {
				b2 |= (((b >> ((k - 0 + 8) % 8)) & 1) ^ ((b >> ((k - 1 + 8) % 8)) & 1) ^ ((b >> ((k - 2 + 8) % 8)) & 1) ^ ((b >> ((k - 3 + 8) % 8)) & 1) ^ ((b >> ((k - 4 + 8) % 8)) & 1) ^ ((0x63 >> (k % 8)) & 1)) << k
			}
			table[i][j] = b2
		}
	}
	return table
}

func sBoxTransfer(b byte, table [][]byte) byte {
	return table[b>>4][b&0xf]
}

func colTransfer(input [][]byte) (output [][]byte) {
	l := len(input)
	output = make([][]byte, l)
	for i := 0; i < l; i++ {
		l2 := len(input[i])
		output[i] = make([]byte, l2)
		for j := 0; j < l2; j++ {
			output[i][j] = input[i][(j+i+l2)%l2]
		}
	}
	return
}

func rowTransfer(input [][]byte) (output [][]byte) {
	output = martixMultiplus(rowTransferTable, input)
	return
}

func martixMultiplus(a [][]byte, b [][]byte) (output [][]byte) {
	row := len(b[0])
	col := len(a)
	l := len(a[0])
	output = make([][]byte, col)
	for i := 0; i < col; i++ {
		output[i] = make([]byte, row)
		for j := 0; j < row; j++ {
			for k := 0; k < l; k++ {
				c := gf.Multiplus(int(a[i][k]), int(b[k][j]))
				if (c >> 8) != 0 {
					c = gf.Plus(c, 0b11011)
				}
				output[i][j] = byte(gf.Plus(int(output[i][j]), c))
			}
		}
	}
	return
}

func keyExpansion(key *bits.Bits) (w [][]byte) {
	sbox := func(b byte) byte { return sBoxTransfer(b, sBoxTable) }
	w = make([][]byte, 44)
	for i := 0; i < 44; i++ {
		for j := 0; j < 4; j++ {
			w[i] = make([]byte, 4)
		}
	}
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			w[i][j] = key.Value[i*4+j]
		}
	}

	for i := 4; i < 44; i++ {
		if i%4 == 0 {
			w[i][0] = sbox(w[i-1][1]) ^ byte(keyGenerateTable[i/4-1]>>24)
			w[i][1] = sbox(w[i-1][2])
			w[i][2] = sbox(w[i-1][3])
			w[i][3] = sbox(w[i-1][0])

			w[i][0] = w[i-4][0] ^ w[i][0]
			w[i][1] = w[i-4][1] ^ w[i][1]
			w[i][2] = w[i-4][2] ^ w[i][2]
			w[i][3] = w[i-4][3] ^ w[i][3]
		} else {
			w[i][0] = w[i-4][0] ^ w[i-1][0]
			w[i][1] = w[i-4][1] ^ w[i-1][1]
			w[i][2] = w[i-4][2] ^ w[i-1][2]
			w[i][3] = w[i-4][3] ^ w[i-1][3]
		}
	}
	return
}

func inverse(input [][]byte) (output [][]byte) {
	output = make([][]byte, len(input))
	for i := 0; i < len(input[0]); i++ {
		output[i] = make([]byte, len(input))
		for j := 0; j < len(input); j++ {
			output[i][j] = input[j][i]
		}
	}
	return
}

func AES(input string, key string) (output string) {
	sbox := func(b byte) byte { return sBoxTransfer(b, sBoxTable) }
	loop := func(f func(int, int)) {
		for i := 0; i < 4; i++ {
			for j := 0; j < 4; j++ {
				f(i, j)
			}
		}
	}
	print := func(data [][]byte, format string, args ...interface{}) {
		Logger.Printf(format, args...)
		for i := 0; i < 4; i++ {
			Logger.Printf("%02x %02x %02x %02x", data[i][0], data[i][1], data[i][2], data[i][3])
		}
		Logger.Printf("\n")
	}

	plaintext := bits.NewBitsFromString(input, 16)
	keys := keyExpansion(bits.NewBitsFromString(key, 16))

	// 明文分组
	p := make([][]byte, 4)
	for i := 0; i < 4; i++ {
		p[i] = make([]byte, 4)
	}
	for i := 0; i < 4; i++ {
		p[i][0] = plaintext.Value[i*4+0]
		p[i][1] = plaintext.Value[i*4+1]
		p[i][2] = plaintext.Value[i*4+2]
		p[i][3] = plaintext.Value[i*4+3]
	}
	p = inverse(p)

	print(p, "开始")

	// 轮密钥加
	loop(func(i int, j int) {
		p[i][j] ^= keys[0+j][i]
	})

	for k := 1; k <= 10; k++ {
		print(p, "第%d轮-开始", k)
		print(keys, "第%d轮-密钥", k)

		// S置换
		loop(func(i int, j int) {
			p[i][j] = sbox(p[i][j])
		})
		print(p, "第%d轮-S置换后", k)
		// 行置换
		p = colTransfer(p)
		print(p, "第%d轮-行置换后", k)

		// 列置换
		if k != 10 {
			p = rowTransfer(p)
		}
		print(p, "第%d轮-列置换后", k)

		// 轮密钥加
		loop(func(i int, j int) {
			p[i][j] ^= keys[k*4+j][i]
		})
		print(p, "第%d轮-结束", k)
	}

	p = inverse(p)

	loop(func(i int, j int) {
		output += fmt.Sprintf("%02x", p[i][j])
	})
	return
}
