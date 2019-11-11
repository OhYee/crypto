package sha

import (
	"github.com/OhYee/goutils"
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SHA1(tt.input); !goutils.Equal(got, tt.want) {
				t.Errorf("Want %+v got %+v\n", tt.want, got)
			}
		})
	}
}
