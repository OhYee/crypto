package main

import (
	"fmt"
	"io/ioutil"
	"regexp"

	"github.com/OhYee/rainbow/color"
)

//go:generate bash -c "CGO_ENABLED=0 GOOS=darwin  GOARCH=amd64 go build -o try_mac     main.go"
//go:generate bash -c "CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o try_win.exe main.go"
//go:generate bash -c "CGO_ENABLED=0 GOOS=linux   GOARCH=amd64 go build -o try_linux   main.go"
const defaultKeys = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
const answer = ""

var red = color.New().SetFrontRed()
var green = color.New().SetFrontGreen()

var k = map[string]string{}
var k2 = map[string]string{}

func readFile(filename string) string {
	keys, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(color.New().SetFrontRed().SetFontBold().Colorful(fmt.Sprintf("Open %s failed, set empty string", filename)))
	}
	return string(keys)
}

func encrypto(input string, key map[string]string, withColor bool) (output string) {
	for _, c := range input {
		s := string(c)
		ordinary, exist := key[s]
		if exist {
			if withColor {
				output += green.Colorful(ordinary)
			} else {
				output += ordinary
			}
		} else {
			if withColor {
				output += red.Colorful(s)
			} else {
				output += s
			}
		}
	}
	return
}

func main() {
	keys := readFile("./keys.txt")
	secret := readFile("./secret.txt")

	reg, err := regexp.Compile("(.)\\s*=>\\s*(.)")
	if err != nil {
		panic(err)
	}
	res := reg.FindAllStringSubmatch(keys, -1)
	for _, item := range res {
		if len(item) == 3 {
			k[item[1]] = item[2]
			k2[item[2]] = item[1]
		}
	}

	fmt.Printf("         %s\n", defaultKeys)
	fmt.Printf("encrypto %s\n", encrypto(defaultKeys, k, true))
	fmt.Printf("decrypto %s\n", encrypto(defaultKeys, k2, true))
	if answer == encrypto(defaultKeys, k, false) {
		fmt.Println(green.Colorful("Success"))
	} else {
		fmt.Println(red.Colorful("Come on!"))
	}
	fmt.Println()
	fmt.Println(encrypto(secret, k2, true))
}
