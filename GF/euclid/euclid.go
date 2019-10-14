package euclid

import (
	"github.com/OhYee/cryptography_and_network_security/util/blackhole"
	"log"
	"reflect"
)

type Operator func(int, int) int
type Operator2 func(int, int) (int, int)

func Plus(a int, b int) int          { return a + b }
func Multiplus(a int, b int) int     { return a * b }
func Divide(a int, b int) (int, int) { return a / b, a % b }

// GCD get the maximum common factor
func GCD(a int, b int, plus Operator, multiplus Operator, divide Operator2) int {
	if b == 0 {
		return a
	}
	_, rr := divide(a, b)
	return GCD(b, rr, plus, multiplus, divide)
}

// ExGCD extend GCD algorithm
func ExGCD(a int, b int, plus Operator, multiplus Operator, divide Operator2) (r int, x int, y int) {
	if b == 0 {
		return a, 1, 0
	}

	c, d := divide(a, b)
	r, x, y = ExGCD(b, d, plus, multiplus, divide)

	t := x
	x = y
	y = plus(t, -multiplus(c, y))

	if  reflect.ValueOf(plus).Pointer() ==  reflect.ValueOf(Plus).Pointer() {
		Logger.Printf("%d * %d + %d * %d = %d\n", a, x, b, y, r)
	} else {
		Logger.Printf("%b * %b + %b * %b = %b\n", a, x, b, y, r)
	}
	return r, x, y
}

var (
	// Logger 日志
	Logger = log.New(blockhole.BlackHole{}, " |Log| ", 0)
)
