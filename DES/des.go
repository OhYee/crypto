package des

import (
	"github.com/OhYee/crypto/DES/bits"
	"github.com/OhYee/rainbow/color"
	"github.com/OhYee/rainbow/log"
)

var Logger = log.New().SetOutputToNil().SetPrefix(func(s string) string {
	return color.New().SetFontBold().Colorful("Log     ")
}).SetNewLine(true)

var (
	// 初始置换
	ip = []int{
		58, 50, 42, 34, 26, 18, 10, 2,
		60, 52, 44, 36, 28, 20, 12, 4,
		62, 54, 46, 38, 30, 22, 14, 6,
		64, 56, 48, 40, 32, 24, 16, 8,
		57, 49, 41, 33, 25, 17, 9, 1,
		59, 51, 43, 35, 27, 19, 11, 3,
		61, 53, 45, 37, 29, 21, 13, 5,
		63, 55, 47, 39, 31, 23, 15, 7,
	}
	// 轮函数拓展置换
	ep = []int{
		32, 1, 2, 3, 4, 5,
		4, 5, 6, 7, 8, 9,
		8, 9, 10, 11, 12, 13,
		12, 13, 14, 15, 16, 17,
		16, 17, 18, 19, 20, 21,
		20, 21, 22, 23, 24, 25,
		24, 25, 26, 27, 28, 29,
		28, 29, 30, 31, 32, 1,
	}
	// 轮函数普通置换
	pp = []int{
		16, 7, 20, 21,
		29, 12, 28, 17,
		1, 15, 23, 26,
		5, 18, 31, 10,
		2, 8, 24, 14,
		32, 27, 3, 9,
		19, 13, 30, 6,
		22, 11, 4, 25,
	}
	// 子密钥左移位数
	keyMove = []int{
		1, 1, 2, 2, 2, 2, 2, 2,
		1, 2, 2, 2, 2, 2, 2, 1,
	}
	// 子密钥置换表1
	pc1 = []int{
		57, 49, 41, 33, 25, 17, 9,
		1, 58, 50, 42, 34, 26, 18,
		10, 2, 59, 51, 43, 35, 27,
		19, 11, 3, 60, 52, 44, 36,
		63, 55, 47, 39, 31, 23, 15,
		7, 62, 54, 46, 38, 30, 22,
		14, 6, 61, 53, 45, 37, 29,
		21, 13, 5, 28, 20, 12, 4,
	}
	// 子密钥置换表2
	pc2 = []int{
		14, 17, 11, 24, 1, 5,
		3, 28, 15, 6, 21, 10,
		23, 19, 12, 4, 26, 8,
		16, 7, 27, 20, 13, 2,
		41, 52, 31, 37, 47, 55,
		30, 40, 51, 45, 33, 48,
		44, 49, 39, 56, 34, 53,
		46, 42, 50, 36, 29, 32,
	}
	// S盒代换表
	sp = [][][]byte{
		{
			{14, 4, 13, 1, 2, 15, 11, 8, 3, 10, 6, 12, 5, 9, 0, 7},
			{0, 15, 7, 4, 14, 2, 13, 1, 10, 6, 12, 11, 9, 5, 3, 8},
			{4, 1, 14, 8, 13, 6, 2, 11, 15, 12, 9, 7, 3, 10, 5, 0},
			{15, 12, 8, 2, 4, 9, 1, 7, 5, 11, 3, 14, 10, 0, 6, 13},
		},
		{
			{15, 1, 8, 14, 6, 11, 3, 4, 9, 7, 2, 13, 12, 0, 5, 10},
			{3, 13, 4, 7, 15, 2, 8, 14, 12, 0, 1, 10, 6, 9, 11, 5},
			{0, 14, 7, 11, 10, 4, 13, 1, 5, 8, 12, 6, 9, 3, 2, 15},
			{13, 8, 10, 1, 3, 15, 4, 2, 11, 6, 7, 12, 0, 5, 14, 9},
		},
		{
			{10, 0, 9, 14, 6, 3, 15, 5, 1, 13, 12, 7, 11, 4, 2, 8},
			{13, 7, 0, 9, 3, 4, 6, 10, 2, 8, 5, 14, 12, 11, 15, 1},
			{13, 6, 4, 9, 8, 15, 3, 0, 11, 1, 2, 12, 5, 10, 14, 7},
			{1, 10, 13, 0, 6, 9, 8, 7, 4, 15, 14, 3, 11, 5, 2, 12},
		},
		{
			{7, 13, 14, 3, 0, 6, 9, 10, 1, 2, 8, 5, 11, 12, 4, 15},
			{13, 8, 11, 5, 6, 15, 0, 3, 4, 7, 2, 12, 1, 10, 14, 9},
			{10, 6, 9, 0, 12, 11, 7, 13, 15, 1, 3, 14, 5, 2, 8, 4},
			{3, 15, 0, 6, 10, 1, 13, 8, 9, 4, 5, 11, 12, 7, 2, 14},
		},
		{
			{2, 12, 4, 1, 7, 10, 11, 6, 8, 5, 3, 15, 13, 0, 14, 9},
			{14, 11, 2, 12, 4, 7, 13, 1, 5, 0, 15, 10, 3, 9, 8, 6},
			{4, 2, 1, 11, 10, 13, 7, 8, 15, 9, 12, 5, 6, 3, 0, 14},
			{11, 8, 12, 7, 1, 14, 2, 13, 6, 15, 0, 9, 10, 4, 5, 3},
		},
		{
			{12, 1, 10, 15, 9, 2, 6, 8, 0, 13, 3, 4, 14, 7, 5, 11},
			{10, 15, 4, 2, 7, 12, 9, 5, 6, 1, 13, 14, 0, 11, 3, 8},
			{9, 14, 15, 5, 2, 8, 12, 3, 7, 0, 4, 10, 1, 13, 11, 6},
			{4, 3, 2, 12, 9, 5, 15, 10, 11, 14, 1, 7, 6, 0, 8, 13},
		},
		{
			{4, 11, 2, 14, 15, 0, 8, 13, 3, 12, 9, 7, 5, 10, 6, 1},
			{13, 0, 11, 7, 4, 9, 1, 10, 14, 3, 5, 12, 2, 15, 8, 6},
			{1, 4, 11, 13, 12, 3, 7, 14, 10, 15, 6, 8, 0, 5, 9, 2},
			{6, 11, 13, 8, 1, 4, 10, 7, 9, 5, 0, 15, 14, 2, 3, 12},
		},
		{
			{13, 2, 8, 4, 6, 15, 11, 1, 10, 9, 3, 14, 5, 0, 12, 7},
			{1, 15, 13, 8, 10, 3, 7, 4, 12, 5, 6, 11, 0, 14, 9, 2},
			{7, 11, 4, 1, 9, 12, 14, 2, 0, 6, 10, 13, 15, 3, 5, 8},
			{2, 1, 14, 7, 4, 10, 8, 13, 15, 12, 9, 0, 3, 5, 6, 11},
		},
	}
)

// inverse 逆置换
func inverse(raw []int) []int {
	m := make(map[int]int)
	for idx, v := range raw {
		m[v] = idx
	}
	res := make([]int, len(raw))
	for idx := range raw {
		res[idx] = m[idx+1] + 1
	}
	return res
}

// sbox S盒
func sbox(input bits.Bits) (output bits.Bits) {
	// 48 bits
	// pos := 0
	for i := 0; i < 8; i++ {
		offset := i * 6
		row := input.Get(offset)*1 + input.Get(offset+5)*2
		col := input.Get(offset+1)*1 + input.Get(offset+2)*2 + input.Get(offset+3)*4 + input.Get(offset+4)*8
		v := bits.Bits(sp[7-i][row][col])
		output = output | (v << (i * 4))
	}
	// 32 bits
	return
}

// pbox P盒 置换盒
func pbox(input bits.Bits, key []int, inputLength int, outputLength int) (output bits.Bits) {
	for i := 0; i < len(key); i++ {
		v := input.Get(inputLength - key[i])
		output.SetValue(outputLength-i-1, v)
	}
	return
}

func f(input bits.Bits, key bits.Bits) (output bits.Bits) {
	output = input // 32 bits
	Logger.Printf("%8s\t0x%08x", "input", output)
	output = pbox(output, ep, 32, 48) // 48 bits
	Logger.Printf("%8s\t0x%08x", "ep", output)
	output = output ^ key // 48 bits
	Logger.Printf("%8s\t0x%08x", "xor", output)
	output = sbox(output) // 32 bits
	Logger.Printf("%8s\t0x%08x", "sbox", output)
	output = pbox(output, pp, 32, 32) // 32 bits
	Logger.Printf("%8s\t0x%08x", "pp", output)
	return
}

func getSubKey(key bits.Bits) (output []bits.Bits) {
	output = make([]bits.Bits, 16)

	// key 64 bit
	temp := pbox(key, pc1, 64, 56) // 56 bits

	C := temp.Mask(28)         // 28 bits
	D := (temp >> 28).Mask(28) // 28 bits

	for i := 0; i < 16; i++ {
		C = C.LeftLoop(keyMove[i], 28)
		D = D.LeftLoop(keyMove[i], 28)

		output[i] = ((D << 28) | C)              // 56 bits
		output[i] = pbox(output[i], pc2, 56, 48) // 48 bits
	}

	return
}

const runCount = 16

// Encrypto DES 加密算法实现
func Encrypto(input bits.Bits, key bits.Bits) (output bits.Bits) {
	var temp bits.Bits
	L := make([]bits.Bits, 2)
	R := make([]bits.Bits, 2)

	output = pbox(input, ip, 64, 64)

	L[0] = (output >> 32).Mask(32) // 32 bits
	R[0] = output.Mask(32)         // 32 bits

	Logger.Printf("%8s\t\t\tL: 0x%08x\tR: 0x%08x", "IP", L[0], R[0])

	keys := getSubKey(key) // 48 bits

	for i := 0; i < runCount; i++ {
		this := i & 1
		next := (i ^ 1) & 1

		temp = f(R[this], keys[i]) // 32 bits

		L[next] = R[this]
		R[next] = temp ^ L[this]

		Logger.Printf("%d\tKey: 0x%012x\tL: 0x%08x\tR: 0x%08x", i+1, keys[i], L[next], R[next])
	}
	output = (R[runCount&1] << 32) | L[runCount&1] // 64 bits
	Logger.Printf("%s\t0x%08x", "output", output)
	output = pbox(output, inverse(ip), 64, 64)
	Logger.Printf("%s\t0x%08x", "invise-ip", output)

	return
}

// Decrypto DES 解密算法实现
func Decrypto(input bits.Bits, key bits.Bits) (output bits.Bits) {
	var temp bits.Bits
	L := make([]bits.Bits, 2)
	R := make([]bits.Bits, 2)

	output = pbox(input, ip, 64, 64)

	L[0] = (output >> 32).Mask(32) // 32 bits
	R[0] = output.Mask(32)         // 32 bits

	Logger.Printf("%8s\t\t\tL: 0x%08x\tR: 0x%08x", "IP", L[0], R[0])

	keys := getSubKey(key) // 48 bits

	for i := 0; i < runCount; i++ {
		this := i & 1
		next := (i ^ 1) & 1

		temp = f(R[this], keys[runCount-1-i]) // 32 bits

		L[next] = R[this]
		R[next] = temp ^ L[this]

		Logger.Printf("%d\tKey: 0x%012x\tL: 0x%08x\tR: 0x%08x", i+1, keys[i], L[next], R[next])
	}
	output = (R[runCount&1] << 32) | L[runCount&1] // 64 bits
	Logger.Printf("%s\t0x%08x", "output", output)
	output = pbox(output, inverse(ip), 64, 64)
	Logger.Printf("%s\t0x%08x", "invise-ip", output)

	return
}
