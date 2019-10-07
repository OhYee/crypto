package main

import (
	"github.com/OhYee/cryptography_and_network_security/DES"
	"fmt"
	"os"
)

func exgcd(a int, b int) (r int, x int, y int) {
	if b == 0 {
		return a, 1, 0
	}
	r, x, y = exgcd(b, a%b)
	t := x
	x = y
	y = t - a/b*y
	fmt.Printf("%d * %d + %d * %d = %d\n", a, x, b, y, r)
	return r, x, y
}

func main() {
	des.Logger.SetOutput(os.Stdout)
	fmt.Printf("%016x\n", des.DES(0x02468ACEECA86420, 0x0F1571C947D9E859))
	r, x, y := exgcd(97, 35)
	fmt.Printf("%d %d %d\n", r, x, y)
}
