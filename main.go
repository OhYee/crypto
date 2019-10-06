package main

import (
	"cryptography_and_network_security/DES"
	"fmt"
	"os"
)

func main() {
	des.Logger.SetOutput(os.Stdout)
	fmt.Printf("%016x\n", des.DES(0x02468ACEECA86420, 0x0F1571C947D9E859))
}
