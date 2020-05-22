package main

import (
	"fmt"
	"os"

	aes "github.com/OhYee/crypto/AES"
	des "github.com/OhYee/crypto/DES"
	gf "github.com/OhYee/crypto/GF"
	"github.com/OhYee/crypto/GF/euclid"
)

func main() {
	des.Logger.SetOutput(os.Stdout)
	fmt.Printf("%016x\n", des.DES(0x02468ACEECA86420, 0x0F1571C947D9E859))
	aes.Logger.SetOutput(os.Stdout)
	fmt.Printf("%s\n", aes.AES("0123456789abcdeffedcba9876543210", "0f1571c947d9e8590cb7add6af7f6798"))
	euclid.Logger.SetOutput(os.Stdout)
	r, x, y := euclid.ExGCD(97, 35, euclid.Plus, euclid.Multiplus, euclid.Divide)
	fmt.Printf("%d %d %d\n", r, x, y)
	r, x, y = euclid.ExGCD(0b100011011, 0b10000011, gf.Plus, gf.Multiplus, gf.Divide)
	fmt.Printf("%d %d %d\n", r, x, y)

}
