package sha

import (
	"encoding/binary"

	"github.com/OhYee/rainbow/color"
	"github.com/OhYee/rainbow/log"
)

var Logger = log.New().SetOutputToNil().SetPrefix(func(s string) string {
	return color.New().SetFontBold().Colorful("Log     ")
})

type sha struct{}

func uint32InputInitial(input []byte) []byte {
	rawLength := len(input) * 8

	// Ensure l≡448 mod 512
	input = append(input, 0x80)
	for (len(input)*8)%512 != 448 {
		input = append(input, 0x00)
	}

	// Add the length at the end of input
	length := make([]byte, 8)
	binary.BigEndian.PutUint64(length, uint64(rawLength))
	input = append(input, length...)

	return input
}

func uint64InputInitial(input []byte) []byte {
	rawLength := len(input) * 8

	// Ensure l≡448 mod 512
	input = append(input, 0x80)
	for (len(input)*8)%1024 != 896 {
		input = append(input, 0x00)
	}

	// Add the length at the end of input
	length := make([]byte, 8)
	binary.BigEndian.PutUint64(length, uint64(0))
	input = append(input, length...)
	binary.BigEndian.PutUint64(length, uint64(rawLength))
	input = append(input, length...)

	return input
}
