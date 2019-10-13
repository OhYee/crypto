package bits

import (
	"fmt"
)

const (
	byteLength = 1 * 8
)

type Bits struct {
	value []byte
}

// stringToByte transfer string to byte.
func stringToByte(s string) (b byte) {
	fmt.Sscanf(s, "%x", &b)
	return
}

// NewBits declare a Bits struct
func NewBits(s string) *Bits {
	for len(s)%2 != 0 {
		s = "0" + s
	}
	bits := make([]byte, len(s)/2)
	for i := 0; i < len(s)/2; i++ {
		bits[len(s)/2-i-1] = stringToByte(s[i*2 : i*2+2])
	}
	return &Bits{
		value: bits,
	}
}

// NewBitsFromBytes declare a Bits struct from bytes
func NewBitsFromBytes(b ...byte) *Bits {
	return &Bits{
		value: b,
	}
}

// Bool return true if the Bits is not 0, otherwise false.
func (b Bits) Bool() bool {
	for _, v := range b.value {
		if v != 0x00 {
			return true
		}
	}
	return false
}

// Get the bit value of the Bits
func (b Bits) Get(pos int) *Bits {
	return NewBitsFromBytes(b.value[pos>>3] & (1 << (pos & 7)))
}

// Set the bit value with 1
func (b *Bits) Set(pos int) Bits {
	b.value[pos>>3] |= (1 << (pos & 7))
	return *b
}

// Clr the bit value with 0
func (b *Bits) Clr(pos int) Bits {
	b.value[pos>>3] &= ^(1 << (pos & 7))
	return *b
}

// SetValue the bit value with boolean
func (b *Bits) SetValue(pos int, bb bool) Bits {
	if bb {
		return b.Set(pos)
	}
	return b.Clr(pos)
}

// Mask return the 0~l sub-bits
func (b Bits) Mask(l int) *Bits {
	bits := b.value[:][0 : l/8]
	if l%8 != 0 {
		bits = append(bits, b.value[l/8]&((1<<(l%8))-1))
	}
	return NewBitsFromBytes(bits...)
}

// func (b Bits) LeftLoop(pos int, len int) Bits {
// 	value := b & ((1 << len) - 1)
// 	for i := 0; i < pos; i++ {
// 		overflow := value.Get(len - 1)
// 		value <<= 1
// 		value.SetValue(0, overflow)
// 	}
// 	return value & ((1 << len) - 1)
// }
