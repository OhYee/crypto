package euclid

import (
	"github.com/OhYee/cryptography_and_network_security/util/blackhole"
	"log"
)

// GCD get the maximum common factor
func GCD(a int, b int) int {
	if b == 0 {
		return a
	}
	return GCD(b, a%b)
}

// ExGCD extend GCD algorithm
func ExGCD(a int, b int) (r int, x int, y int) {
	if b == 0 {
		return a, 1, 0
	}
	r, x, y = ExGCD(b, a%b)
	t := x
	x = y
	y = t - a/b*y
	Logger.Printf("%d * %d + %d * %d = %d\n", a, x, b, y, r)
	return r, x, y
}

var (
	// Logger 日志
	Logger = log.New(blockhole.BlackHole{}, " |Log| ", 0)
)
