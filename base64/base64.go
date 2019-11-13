package base64

import (
	"fmt"
	"bytes"
	"strings"
)

const char = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

func BASE64(input []byte) string {

	buf := bytes.NewBufferString("")
	extra := 0
	switch (len(input) * 8) % 6 {
	case 0:
		extra = 0
	case 2:
		extra = 2
		input = append(input, 0x00, 0x00)
	case 4:
		extra = 1
		input = append(input, 0x00)
	}
	getBit := func(pos int) byte {
		return (input[pos>>3] >> (7 - (pos & 0x07))) & 1
	}
	l := len(input)*8/6 - extra
	for i := 0; i < l; i++ {
		num := (getBit(i*6) << 5) | (getBit(i*6+1) << 4) | (getBit(i*6+2) << 3) | (getBit(i*6+3) << 2) | (getBit(i*6+4) << 1) | (getBit(i*6+5) << 0)
		// fmt.Printf("%d %06b\n", num, num)
		buf.WriteByte(char[num])
	}
	for i := 0; i < extra; i++ {
		buf.WriteRune('=')
	}
	return buf.String()
}

func UNBASE64(input string) []byte {
	buf := bytes.NewBuffer([]byte{})
	var temp byte
	var num int
	var extra int
	for _, c := range input {
		if c == '=' {
			extra++
		} else {
			b := byte(strings.IndexByte(char, byte(c)))
			fmt.Printf("%06b\n",b)
			switch num {
			case 0:
				temp = b << 2
				num = 2
			case 2:
				buf.WriteByte(temp | (b>>4))
				temp = b << 4
				num = 4
			case 4:
				buf.WriteByte(temp | (b>>2))
				temp = b << 6
				num = 6
			case 6:
				buf.WriteByte(temp | b)
				temp = 0
				num = 0
			}
		}

	}
	bb := buf.Bytes()
	return bb[:]
}
