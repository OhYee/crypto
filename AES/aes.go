package aes

import (
	"fmt"
	"github.com/OhYee/cryptography_and_network_security/AES/bits"
)

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
	return [][]byte{
		{0x63, 0x7c, 0x77, 0x7b, 0xf2, 0x6b, 0x6f, 0xc5, 0x30, 0x01, 0x67, 0x2b, 0xfe, 0xd7, 0xab, 0x76},
		{0xca, 0x82, 0xc9, 0x7d, 0xfa, 0x59, 0x47, 0xf0, 0xad, 0xd4, 0xa2, 0xaf, 0x9c, 0xa4, 0x72, 0xc0},
		{0xb7, 0xfd, 0x93, 0x26, 0x36, 0x3f, 0xf7, 0xcc, 0x34, 0xa5, 0xe5, 0xf1, 0x71, 0xd8, 0x31, 0x15},
		{0x04, 0xc7, 0x23, 0xc3, 0x18, 0x96, 0x05, 0x9a, 0x07, 0x12, 0x80, 0xe2, 0xeb, 0x27, 0xb2, 0x75},
		{0x09, 0x83, 0x2c, 0x1a, 0x1b, 0x6e, 0x5a, 0xa0, 0x52, 0x3b, 0xd6, 0xb3, 0x29, 0xe3, 0x2f, 0x84},
		{0x53, 0xd1, 0x00, 0xed, 0x20, 0xfc, 0xb1, 0x5b, 0x6a, 0xcb, 0xbe, 0x39, 0x4a, 0x4c, 0x58, 0xcf},
		{0xd0, 0xef, 0xaa, 0xfb, 0x43, 0x4d, 0x33, 0x85, 0x45, 0xf9, 0x02, 0x7f, 0x50, 0x3c, 0x9f, 0xa8},
		{0x51, 0xa3, 0x40, 0x8f, 0x92, 0x9d, 0x38, 0xf5, 0xbc, 0xb6, 0xda, 0x21, 0x10, 0xff, 0xf3, 0xd2},
		{0xcd, 0x0c, 0x13, 0xec, 0x5f, 0x97, 0x44, 0x17, 0xc4, 0xa7, 0x7e, 0x3d, 0x64, 0x5d, 0x19, 0x73},
		{0x60, 0x81, 0x4f, 0xdc, 0x22, 0x2a, 0x90, 0x88, 0x46, 0xee, 0xb8, 0x14, 0xde, 0x5e, 0x0b, 0xdb},
		{0xe0, 0x32, 0x3a, 0x0a, 0x49, 0x06, 0x24, 0x5c, 0xc2, 0xd3, 0xac, 0x62, 0x91, 0x95, 0xe4, 0x79},
		{0xe7, 0xc8, 0x37, 0x6d, 0x8d, 0xd5, 0x4e, 0xa9, 0x6c, 0x56, 0xf4, 0xea, 0x65, 0x7a, 0xae, 0x08},
		{0xba, 0x78, 0x25, 0x2e, 0x1c, 0xa6, 0xb4, 0xc6, 0xe8, 0xdd, 0x74, 0x1f, 0x4b, 0xbd, 0x8b, 0x8a},
		{0x70, 0x3e, 0xb5, 0x66, 0x48, 0x03, 0xf6, 0x0e, 0x61, 0x35, 0x57, 0xb9, 0x86, 0xc1, 0x1d, 0x9e},
		{0xe1, 0xf8, 0x98, 0x11, 0x69, 0xd9, 0x8e, 0x94, 0x9b, 0x1e, 0x87, 0xe9, 0xce, 0x55, 0x28, 0xdf},
		{0x8c, 0xa1, 0x89, 0x0d, 0xbf, 0xe6, 0x42, 0x68, 0x41, 0x99, 0x2d, 0x0f, 0xb0, 0x54, 0xbb, 0x16},
	}
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

func plus(a byte, b byte) byte {
	return a ^ b
}

func multiplus(a byte, b byte) (c byte, carry bool) {
	for i := 0; i < 8; i++ {
		if a&(1<<i) != 0 {
			carry = carry || ((int(b)<<i)>>8) > 0
			c ^= b << i
		}
	}
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
				c, c2 := multiplus(a[i][k], b[k][j])
				if c2 {
					c = plus(c, 0b11011)
				}
				output[i][j] = plus(output[i][j], c)
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
	plaintext := bits.NewBitsFromString(input, 16)
	keys := keyExpansion(bits.NewBitsFromString(key, 16))

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

	sbox := func(b byte) byte { return sBoxTransfer(b, sBoxTable) }
	loop := func(f func(int, int)) {
		for i := 0; i < 4; i++ {
			for j := 0; j < 4; j++ {
				f(i, j)
			}
		}
	}

	loop(func(i int, j int) {
		if i == 0 && j == 0 {
			fmt.Printf("开始\n")
		}
		fmt.Printf("%02x ", p[i][j])
		if j == 3 {
			fmt.Printf("\n")
			if i == 3 {
				fmt.Printf("\n")
			}
		}
	})

	loop(func(i int, j int) {
		p[i][j] ^= keys[0+j][i]
	})

	for k := 1; k <= 10; k++ {
		loop(func(i int, j int) {
			if i == 0 && j == 0 {
				fmt.Printf("%d开始\n", k)
			}
			fmt.Printf("%02x ", p[i][j])
			if j == 3 {
				fmt.Printf("\n")
				if i == 3 {
					fmt.Printf("\n")
				}
			}
		})
		// S置换
		loop(func(i int, j int) {
			p[i][j] = sbox(p[i][j])
		})
		loop(func(i int, j int) {
			if i == 0 && j == 0 {
				fmt.Printf("%d置换\n", k)
			}
			fmt.Printf("%02x ", p[i][j])
			if j == 3 {
				fmt.Printf("\n")
				if i == 3 {
					fmt.Printf("\n")
				}
			}
		})
		// 行置换
		p = colTransfer(p)
		loop(func(i int, j int) {
			if i == 0 && j == 0 {
				fmt.Printf("%d行\n", k)
			}
			fmt.Printf("%02x ", p[i][j])
			if j == 3 {
				fmt.Printf("\n")
				if i == 3 {
					fmt.Printf("\n")
				}
			}
		})
		// 列置换
		if k != 10 {
			p = rowTransfer(p)
		}
		loop(func(i int, j int) {
			if i == 0 && j == 0 {
				fmt.Printf("%d列\n", k)
			}
			fmt.Printf("%02x ", p[i][j])
			if j == 3 {
				fmt.Printf("\n")
				if i == 3 {
					fmt.Printf("\n")
				}
			}
		})
		// 密钥相加
		loop(func(i int, j int) {
			p[i][j] ^= keys[k*4+j][i]
		})
		loop(func(i int, j int) {
			if i == 0 && j == 0 {
				fmt.Printf("%d结束\n", k)
			}
			fmt.Printf("%02x ", p[i][j])
			if j == 3 {
				fmt.Printf("\n")
				if i == 3 {
					fmt.Printf("\n")
				}
			}
		})

		loop(func(i int, j int) {
			if i == 0 && j == 0 {
				fmt.Printf("%d key\n", k)
			}
			fmt.Printf("%02x ", keys[k*4+j][i])
			if j == 3 {
				fmt.Printf("\n")
				if i == 3 {
					fmt.Printf("\n")
				}
			}
		})
	}

	p = inverse(p)

	loop(func(i int, j int) {
		output += fmt.Sprintf("%02x", p[i][j])
	})
	return
}
