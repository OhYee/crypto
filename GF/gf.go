package gf

import (
	"github.com/OhYee/crypto/GF/euclid"
)

func abs(n int) int {
	if n < 0 {
		n = -n
	}
	return n
}

// Plus GF(2^n) plus a + b
func Plus(a int, b int) int {
	a = abs(a)
	b = abs(b)
	return a ^ b
}

// Multiplus GF(2^n) multiplus a * b
func Multiplus(a int, b int) (c int) {
	a = abs(a)
	b = abs(b)

	aa := GetHighestBit(a)
	for i := 0; i <= aa; i++ {
		if (a>>i)&1 != 0 {
			c = Plus(c, b<<i)
		}
	}
	return
}

func Divide(a int, b int) (c int, r int) {
	a = abs(a)
	b = abs(b)
	if a == 0 || b == 0 {
		c = 0
		r = 0
		return
	}
	aa := GetHighestBit(a)
	bb := GetHighestBit(b)

	r = 0
	for i := aa; i >= 0; i-- {
		r = (r << 1) | ((a >> i) & 1)
		if (r>>bb)&1 == 1 {
			c = (c << 1) | 1
			r ^= b
		} else {
			c = c << 1
		}
	}
	return
}

func GetHighestBit(n int) (l int) {
	n = abs(n)
	for n != 0 {
		l++
		n /= 2
	}
	l--
	return
}

func MultiPlusTable(n int, m int) (table [][]int) {
	n = abs(n)
	m = abs(m)

	num := 1 << n
	// mm := GetHighestBit(m)

	table = make([][]int, num)
	for i := 0; i < num; i++ {
		table[i] = make([]int, num)
		for j := 0; j < num; j++ {
			table[i][j] = Multiplus(i, j)
			_, r := Divide(table[i][j], m)
			table[i][j] = r
		}
	}
	return
}

func Inverse(a int, m int) (c int) {
	a = abs(a)
	m = abs(m)

	_, x, _ := euclid.ExGCD(a, m, Plus, Multiplus, Divide)
	return x
}
