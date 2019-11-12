package sha

import (
	"crypto/sha1"
	"fmt"
	"github.com/OhYee/goutils"
	"math/rand"
	"testing"
)

func TestSHA1(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name  string
		input []byte
		want  []byte
	}{
		{
			name:  "iscbupt",
			input: []byte("iscbupt"),
			want: []byte{
				0x66, 0x4D, 0xC9, 0xF0,
				0x17, 0xDC, 0x1A, 0xEE,
				0x4A, 0x43, 0x66, 0xBC,
				0xFB, 0x85, 0x11, 0xAF,
				0xC8, 0x9F, 0x94, 0x30,
			},
		},
		{
			name:  "Beijing University of Posts and Telecommunications",
			input: []byte("Beijing University of Posts and Telecommunications"),
			want: []byte{
				0xC7, 0x0A, 0xEC, 0x84,
				0xB4, 0x35, 0xD6, 0x96,
				0x59, 0xD2, 0x4A, 0xBA,
				0x72, 0x22, 0x2B, 0x7E,
				0xE1, 0xA6, 0xEB, 0xE1,
			},
		},
		{
			name:  "test 440bit",
			input: []byte("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
			want: []byte{
				0xc1, 0xc8, 0xbb, 0xdc,
				0x22, 0x79, 0x6e, 0x28,
				0xc0, 0xe1, 0x51, 0x63,
				0xd2, 0x08, 0x99, 0xb6,
				0x56, 0x21, 0xd6, 0x5a,
			},
		},
		{
			name:  "test 448bit",
			input: []byte("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
			want: []byte{
				0xc2, 0xdb, 0x33, 0x0f,
				0x60, 0x83, 0x85, 0x4c,
				0x99, 0xd4, 0xb5, 0xbf,
				0xb6, 0xe8, 0xf2, 0x9f,
				0x20, 0x1b, 0xe6, 0x99,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SHA1(tt.input); !goutils.Equal(got, tt.want) {
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
			hash := sha1.New()
			hash.Write(b)
			want := hash.Sum([]byte{})
			got := SHA1(b)
			if !goutils.Equal(got, want) {
				t.Errorf("want %+v, got %+v", want, got)
			}
		})
	}
}
