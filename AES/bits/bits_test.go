package bits

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"
)

func Test_stringToByte(t *testing.T) {
	type testCase struct {
		name  string
		s     string
		wantB byte
	}
	tests := make([]testCase, 256)
	for i := 0; i < 256; i++ {
		tests[i] = testCase{
			name:  fmt.Sprintf("0x%x", i),
			s:     fmt.Sprintf("%x", i),
			wantB: byte(i),
		}
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotB := stringToByte(tt.s); gotB != tt.wantB {
				t.Errorf("stringToByte() = %v, want %v", gotB, tt.wantB)
			}
		})
	}
}

func TestNewBits(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want *Bits
	}{
		{name: "0xABCDE", s: "ABCDE", want: &Bits{[]byte{0xDE, 0xBC, 0x0A}}},
		{name: "0x0", s: "0", want: &Bits{[]byte{0x00}}},
		{name: "0x00", s: "00", want: &Bits{[]byte{0x00}}},
		{name: "0x000", s: "000", want: &Bits{[]byte{0x00, 0x00}}},
		{name: "0xF0a", s: "F0a", want: &Bits{[]byte{0x0A, 0x0F}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBits(tt.s); !reflect.DeepEqual(got.value, tt.want.value) {
				t.Errorf("NewBits() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBits_Bool(t *testing.T) {
	type fields struct {
		value []byte
	}
	tests := []struct {
		name string
		b    *Bits
		want bool
	}{
		{name: "0xABCDE", b: &Bits{[]byte{0xDE, 0xBC, 0x0A}}, want: true},
		{name: "0x0", b: &Bits{[]byte{0x00}}, want: false},
		{name: "0x00", b: &Bits{[]byte{0x00}}, want: false},
		{name: "0x000", b: &Bits{[]byte{0x00, 0x00}}, want: false},
		{name: "0xF0a", b: &Bits{[]byte{0x0A, 0x0F}}, want: true},
		{name: "0xA0000000", b: &Bits{[]byte{0x00, 0x00, 0x00, 0x0A}}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.Bool(); got != tt.want {
				t.Errorf("Bits.Bool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewBitsFromBytes(t *testing.T) {
	tests := []struct {
		name string
		b    []byte
		want *Bits
	}{
		{name: "0xABCD", b: []byte{0xCD, 0xab}, want: &Bits{[]byte{0xcd, 0xab}}},
		{name: "0x0000", b: []byte{0x00, 0x00}, want: &Bits{[]byte{0x00, 0x00}}},
		{name: "0xA0", b: []byte{0xA0}, want: &Bits{[]byte{0xA0}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBitsFromBytes(tt.b...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBitsFromBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBits_Get(t *testing.T) {
	tests := []struct {
		name string
		b    *Bits
		pos  int
		want bool
	}{
		{name: "1st byte value", b: NewBits("ABCDEF"), pos: 0, want: true},
		{name: "1st byte value", b: NewBits("ABCDEF"), pos: 2, want: true},
		{name: "1st byte value", b: NewBits("ABCDEF"), pos: 4, want: false},
		{name: "2nd byte value", b: NewBits("ABCDEF"), pos: 8, want: true},
		{name: "2nd byte value", b: NewBits("ABCDEF"), pos: 9, want: false},
		{name: "3rd byte value", b: NewBits("ABCDEF"), pos: 22, want: false},
		{name: "3rd byte value", b: NewBits("ABCDEF"), pos: 23, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.Get(tt.pos).Bool(); got != tt.want {
				t.Errorf("Bits.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBits_Set(t *testing.T) {
	tests := []struct {
		name  string
		b     *Bits
		pos   int
		wantB *Bits
	}{
		{name: "1st byte value", b: NewBits("00"), pos: 0, wantB: NewBits("01")},
		{name: "2st byte value", b: NewBits("0000"), pos: 8, wantB: NewBits("0100")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.Set(tt.pos); !reflect.DeepEqual(got.value, tt.wantB.value) {
				t.Errorf("Bits.Set() = %v, want %v", got, tt.wantB)
			}
		})
	}
}

func TestBits_Clr(t *testing.T) {
	tests := []struct {
		name  string
		b     *Bits
		pos   int
		wantB *Bits
	}{
		{name: "1st byte value", b: NewBits("FF"), pos: 0, wantB: NewBits("FE")},
		{name: "2st byte value", b: NewBits("FFFF"), pos: 8, wantB: NewBits("FEFF")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.Clr(tt.pos); !reflect.DeepEqual(got.value, tt.wantB.value) {
				t.Errorf("Bits.Clr() = %v, want %v", got, tt.wantB)
			}
		})
	}
}

func TestBits_SetValue(t *testing.T) {
	tests := []struct {
		name  string
		b     *Bits
		pos   int
		bb    bool
		wantB *Bits
	}{
		{name: "1st byte value", b: NewBits("00"), pos: 0, bb: true, wantB: NewBits("01")},
		{name: "2st byte value", b: NewBits("0000"), pos: 8, bb: true, wantB: NewBits("0100")},
		{name: "1st byte value", b: NewBits("FF"), pos: 0, bb: false, wantB: NewBits("FE")},
		{name: "2st byte value", b: NewBits("FFFF"), pos: 8, bb: false, wantB: NewBits("FEFF")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.SetValue(tt.pos, tt.bb); !reflect.DeepEqual(got.value, tt.wantB.value) {
				t.Errorf("Bits.SetValue() = %v, want %v", got, tt.wantB)
			}
		})
	}
}

func TestBits_Mask(t *testing.T) {
	tests := []struct {
		name string
		b    *Bits
		l    int
		want *Bits
	}{
		{name: "mask 0", b: NewBits("FFF"), l: 0, want: NewBitsFromBytes()},
		{name: "mask 1", b: NewBits("FFF"), l: 1, want: NewBitsFromBytes(0x01)},
		{name: "mask 4", b: NewBits("FFF"), l: 4, want: NewBitsFromBytes(0x0F)},
		{name: "mask 8", b: NewBits("FFF"), l: 8, want: NewBitsFromBytes(0xFF)},
		{name: "mask 9", b: NewBits("FFF"), l: 9, want: NewBitsFromBytes(0xFF, 0x01)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.Mask(tt.l); !bytes.Equal(got.value, tt.want.value) {
				t.Errorf("Bits.Mask() = %v, want %v", got, tt.want)
			}
		})
	}
}
