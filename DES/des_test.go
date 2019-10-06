package des

import (
	"cryptography_and_network_security/DES/bits"
	"reflect"
	"testing"
)

func Test_inverse(t *testing.T) {
	type args struct {
		raw []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "IP",
			args: args{
				raw: ip,
			},
			want: []int{
				40, 8, 48, 16, 56, 24, 64, 32,
				39, 7, 47, 15, 55, 23, 63, 31,
				38, 6, 46, 14, 54, 22, 62, 30,
				37, 5, 45, 13, 53, 21, 61, 29,
				36, 4, 44, 12, 52, 20, 60, 28,
				35, 3, 43, 11, 51, 19, 59, 27,
				34, 2, 42, 10, 50, 18, 58, 26,
				33, 1, 41, 9, 49, 17, 57, 25,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := inverse(tt.args.raw); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPBox(t *testing.T) {
	data := bits.Bits(0x02468aceeca86420)
	data2 := bits.Bits(0x5a005a003cf03c0f)

	type args struct {
		input        bits.Bits
		key          []int
		inputLength  int
		outputLength int
	}

	tests := []struct {
		name       string
		args       args
		wantOutput bits.Bits
	}{
		{
			name: "simple",
			args: args{
				input:        0b01110011,
				key:          []int{2, 3, 4, 5, 6, 7, 8, 1},
				inputLength:  8,
				outputLength: 8,
			},
			wantOutput: 0b11100110,
		},
		{
			name: "pbox",
			args: args{
				input:        data,
				key:          ip,
				inputLength:  64,
				outputLength: 64,
			},
			wantOutput: data2,
		},
		{
			name: "pbox",
			args: args{
				input:        data2,
				key:          inverse(ip),
				inputLength:  64,
				outputLength: 64,
			},
			wantOutput: data,
		},
		{
			name: "ebox",
			args: args{
				input:        0b010011000111,
				key:          []int{1, 3, 5, 7},
				inputLength:  12,
				outputLength: 4,
			},
			wantOutput: 0b0010,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutput := pbox(tt.args.input, tt.args.key, tt.args.inputLength, tt.args.outputLength); !reflect.DeepEqual(gotOutput, tt.wantOutput) {
				t.Errorf("PBox() = 0x%016x, want 0x%016x", gotOutput, tt.wantOutput)
			}
		})
	}
}

func TestDES(t *testing.T) {
	type args struct {
		input bits.Bits
		key   bits.Bits
	}
	tests := []struct {
		name       string
		plaintext  bits.Bits
		key        bits.Bits
		cryptotext bits.Bits
	}{
		{
			name:       "standard",
			plaintext:  0x02468ACEECA86420,
			key:        0x0F1571C947D9E859,
			cryptotext: 0xda02ce3a89ecac3b,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutput := DES(tt.plaintext, tt.key); !reflect.DeepEqual(gotOutput, tt.cryptotext) {
				t.Errorf("DES() = %016x, want %016x", gotOutput, tt.cryptotext)
			}
		})
	}
}
