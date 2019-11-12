package sha

import (
	"crypto/sha512"
	"fmt"
	"github.com/OhYee/goutils"
	"math/rand"
	"testing"
)

func TestSHA512(t *testing.T) {
	tests := []struct {
		name  string
		input []byte
		want  []byte
	}{
		{
			name:  "iscbupt",
			input: []byte("iscbupt"),
			want: []byte{
				0x2D, 0x13, 0xE3, 0xB5,
				0x52, 0x39, 0xBB, 0xDF,
				0x69, 0x25, 0x83, 0x50,
				0xB0, 0xF9, 0x1A, 0x33,
				0x36, 0x01, 0x58, 0xAA,
				0x5C, 0x6C, 0x63, 0x6C,
				0xFD, 0xA7, 0xA4, 0xC4,
				0x60, 0xEF, 0xBE, 0x7F,
				0x8D, 0xD5, 0xF2, 0x81,
				0xBC, 0x31, 0x43, 0x9B,
				0xAB, 0xC7, 0xEC, 0x79,
				0x82, 0x08, 0x27, 0x3D,
				0xF4, 0x8E, 0xD2, 0x86,
				0x4A, 0x6B, 0x11, 0xFD,
				0xBC, 0x87, 0xA6, 0x36,
				0xB9, 0x0D, 0x7E, 0x03,
			},
		},
		{
			name:  "Beijing University of Posts and Telecommunications",
			input: []byte("Beijing University of Posts and Telecommunications"),
			want: []byte{
				0x61, 0x98, 0xBF, 0x51,
				0x7B, 0x30, 0xAF, 0xA7,
				0x84, 0x1C, 0x1C, 0xD3,
				0x5B, 0xEB, 0x13, 0x0C,
				0x68, 0x44, 0x12, 0x32,
				0x45, 0xCC, 0x9D, 0x30,
				0xA0, 0x32, 0x48, 0xF0,
				0x49, 0x9A, 0x6C, 0x0F,
				0xF9, 0xFB, 0xB0, 0x85,
				0x2C, 0xEA, 0x4B, 0x0A,
				0xC9, 0x80, 0xEA, 0x40,
				0x71, 0x29, 0x88, 0x04,
				0x2A, 0xB2, 0xFF, 0xB4,
				0x96, 0x35, 0xE3, 0xD1,
				0x8D, 0x87, 0xE7, 0x33,
				0x72, 0x71, 0x92, 0xB2,
			},
		},
		{
			name:  "test 440bit",
			input: []byte("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
			want: []byte{
				0xB0, 0x22, 0x0C, 0x77,
				0x2C, 0xBF, 0x6C, 0x18,
				0x22, 0xE2, 0xCB, 0x38,
				0xA4, 0x37, 0xD0, 0xE1,
				0xD5, 0x87, 0x72, 0x41,
				0x7A, 0x4B, 0xBB, 0x21,
				0xC9, 0x61, 0x36, 0x4F,
				0x8B, 0x61, 0x43, 0xE0,
				0x5A, 0xA6, 0x31, 0x6D,
				0xCA, 0x8D, 0x1D, 0x7B,
				0x19, 0xE1, 0x64, 0x48,
				0x41, 0x90, 0x76, 0x39,
				0x5F, 0x60, 0x86, 0xCB,
				0x55, 0x10, 0x1F, 0xBD,
				0x6D, 0x54, 0x97, 0xB1,
				0x48, 0xE1, 0x74, 0x5F,
			},
		},
		{
			name:  "test 448bit",
			input: []byte("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
			want: []byte{
				0x96, 0x2B, 0x64, 0xAA,
				0xE3, 0x57, 0xD2, 0xA4,
				0xFE, 0xE3, 0xDE, 0xD8,
				0xB5, 0x39, 0xBD, 0xC9,
				0xD3, 0x25, 0x08, 0x18,
				0x22, 0xB0, 0xBF, 0xC5,
				0x55, 0x83, 0x13, 0x3A,
				0xAB, 0x44, 0xF1, 0x8B,
				0xAF, 0xE1, 0x1D, 0x72,
				0xA7, 0xAE, 0x16, 0xC7,
				0x9C, 0xE2, 0xBA, 0x62,
				0x0A, 0xE2, 0x24, 0x2D,
				0x51, 0x44, 0x80, 0x91,
				0x61, 0x94, 0x5F, 0x13,
				0x67, 0xF4, 0x1B, 0x39,
				0x72, 0xE2, 0x6E, 0x04,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SHA512(tt.input); !goutils.Equal(got, tt.want) {
				t.Errorf("Want %+v got %+v\n", tt.want, got)
			}
		})
	}

	for i := 0; i < 10; i++ {
		length := rand.Intn(10000)
		b := make([]byte, length)
		for j := 0; j < length; j++ {
			b[j] = uint8(rand.Uint32() % 0xff)
		}

		t.Run(fmt.Sprintf("Random test %d", i), func(t *testing.T) {
			hash := sha512.New()
			hash.Write(b)
			want := hash.Sum([]byte{})
			got := SHA512(b)
			if !goutils.Equal(got, want) {
				t.Errorf("want %+v, got %+v", want, got)
			}
		})
	}
}