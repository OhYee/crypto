package aes

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/OhYee/cryptography_and_network_security/AES/bits"
	"github.com/OhYee/goutils"
)

func Test_generateSBox(t *testing.T) {
	tests := []struct {
		name string
		want [][]byte
	}{
		{
			name: "sbox",
			want: [][]byte{
				{0x63, 0x7c, 0x77, 0x7b, 0xf2, 0x6b, 0x6f, 0xc5, 0x30, 0x01, 0x67, 0x2b, 0xfe, 0xd7, 0xab, 0x76},
				{0xca, 0x82, 0xc9, 0x7d, 0xfa, 0x59, 0x47, 0xf0, 0xad, 0xd4, 0xa2, 0xaf, 0x9c, 0xa4, 0x72, 0xc0},
				{0xb7, 0xfd, 0x93, 0x26, 0x36, 0x3f, 0xf7, 0xcc, 0x34, 0xa5, 0xe5, 0xf1, 0x71, 0xd8, 0x31, 0x15},
				{0x04, 0xc7, 0x23, 0xc3, 0x18, 0x96, 0x05, 0x9a, 0x07, 0x12, 0x80, 0xe2, 0xeb, 0x27, 0xb2, 0x75},
				{0x09, 0x83, 0x2c, 0x1a, 0x1b, 0x6e, 0x5a, 0xa0, 0x52, 0x3b, 0xd6, 0xb3, 0x29, 0xe3, 0x2f, 0x84},
				{0x53, 0xd1, 0x00, 0xed, 0x20, 0xfc, 0xb1, 0x5b, 0x6a, 0xcb, 0xbe, 0x39, 0x4a, 0x4c, 0x58, 0xcf},
				{0xd0, 0xef, 0xaa, 0xfb, 0x43, 0x4d, 0x33, 0x85, 0x45, 0xf9, 0x02, 0x7f, 0x50, 0x3c, 0x9f, 0xa8},
				{0x51, 0xa3, 0x40, 0x8f, 0x92, 0x9d, 0x38, 0xf5, 0xbc, 0xb6, 0xda, 0x21, 0x10, 0xff, 0xf3, 0xd2},
				{0xcd, 0x0c, 0x13, 0xec, 0x5f, 0x97, 0x44, 0x17, 0xc4, 0xa7, 0x7e, 0x3d, 0x64, 0x5d, 0x19, 0x73},
				{0x60, 0x81, 0x4f, 0xdc, 0x22, 0x2a, 0x90, 0x88, 0x46, 0xee, 0xb8, 0x14, 0xde, 0x5e, 0x0b, 0xdb},
				{0xe0, 0x32, 0x3a, 0x0a, 0x49, 0x06, 0x24, 0x5c, 0xc2, 0xd3, 0xac, 0x62, 0x91, 0x95, 0xe4, 0x79},
				{0xe7, 0xc8, 0x37, 0x6d, 0x8d, 0xd5, 0x4e, 0xa9, 0x6c, 0x56, 0xf4, 0xea, 0x65, 0x7a, 0xae, 0x08},
				{0xba, 0x78, 0x25, 0x2e, 0x1c, 0xa6, 0xb4, 0xc6, 0xe8, 0xdd, 0x74, 0x1f, 0x4b, 0xbd, 0x8b, 0x8a},
				{0x70, 0x3e, 0xb5, 0x66, 0x48, 0x03, 0xf6, 0x0e, 0x61, 0x35, 0x57, 0xb9, 0x86, 0xc1, 0x1d, 0x9e},
				{0xe1, 0xf8, 0x98, 0x11, 0x69, 0xd9, 0x8e, 0x94, 0x9b, 0x1e, 0x87, 0xe9, 0xce, 0x55, 0x28, 0xdf},
				{0x8c, 0xa1, 0x89, 0x0d, 0xbf, 0xe6, 0x42, 0x68, 0x41, 0x99, 0x2d, 0x0f, 0xb0, 0x54, 0xbb, 0x16},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateSBox(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateSBox() = %v, want %v", got, tt.want)
				for i := 0; i < 16; i++ {
					for j := 0; j < 16; j++ {
						fmt.Printf("%x ", got[i][j])
					}
					fmt.Printf("\n")
				}
			}
		})
	}
}

func Test_sBoxTransfer(t *testing.T) {
	tests := []struct {
		name  string
		b     byte
		table [][]byte
		want  byte
	}{
		{name: "sample 12", b: 0x12, table: sBoxTable, want: 0xC9},
		{name: "sample EA", b: 0xEA, table: sBoxTable, want: 0x87},
		{name: "sample 04", b: 0x04, table: sBoxTable, want: 0xF2},
		{name: "sample 65", b: 0x65, table: sBoxTable, want: 0x4D},
		{name: "sample 85", b: 0x85, table: sBoxTable, want: 0x97},
		{name: "sample 83", b: 0x83, table: sBoxTable, want: 0xEC},
		{name: "sample 45", b: 0x45, table: sBoxTable, want: 0x6E},
		{name: "sample 5D", b: 0x5D, table: sBoxTable, want: 0x4C},
		{name: "sample 96", b: 0x96, table: sBoxTable, want: 0x90},
		{name: "sample 5C", b: 0x5C, table: sBoxTable, want: 0x4A},
		{name: "sample 33", b: 0x33, table: sBoxTable, want: 0xC3},
		{name: "sample 98", b: 0x98, table: sBoxTable, want: 0x46},
		{name: "sample B0", b: 0xB0, table: sBoxTable, want: 0xE7},
		{name: "sample F0", b: 0xF0, table: sBoxTable, want: 0x8C},
		{name: "sample 2D", b: 0x2D, table: sBoxTable, want: 0xD8},
		{name: "sample AD", b: 0xAD, table: sBoxTable, want: 0x95},
		{name: "sample C5", b: 0xC5, table: sBoxTable, want: 0xA6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sBoxTransfer(tt.b, tt.table); got != tt.want {
				t.Errorf("sBoxTransfer() = %x, want %x", got, tt.want)
			}
		})
	}
}

func Test_colTransfer(t *testing.T) {
	tests := []struct {
		name       string
		input      [][]byte
		wantOutput [][]byte
	}{
		{
			name: "sample",
			input: [][]byte{
				{0x87, 0xf2, 0x4d, 0x97},
				{0xec, 0x6e, 0x4c, 0x90},
				{0x4a, 0xc3, 0x46, 0xe7},
				{0x8c, 0xd8, 0x95, 0xa6},
			},
			wantOutput: [][]byte{
				{0x87, 0xf2, 0x4d, 0x97},
				{0x6e, 0x4c, 0x90, 0xec},
				{0x46, 0xe7, 0x4a, 0xc3},
				{0xa6, 0x8c, 0xd8, 0x95},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutput := colTransfer(tt.input); !reflect.DeepEqual(gotOutput, tt.wantOutput) {
				t.Errorf("colTransfer() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}

func Test_rowTransfer(t *testing.T) {
	tests := []struct {
		name       string
		input      [][]byte
		wantOutput [][]byte
	}{
		{
			name: "sample",
			input: [][]byte{
				{0x87, 0xf2, 0x4d, 0x97},
				{0x6e, 0x4c, 0x90, 0xec},
				{0x46, 0xe7, 0x4a, 0xc3},
				{0xa6, 0x8c, 0xd8, 0x95},
			},
			wantOutput: [][]byte{
				{0x47, 0x40, 0xa3, 0x4c},
				{0x37, 0xd4, 0x70, 0x9f},
				{0x94, 0xe4, 0x3a, 0x42},
				{0xed, 0xa5, 0xa6, 0xbc},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutput := rowTransfer(tt.input); !reflect.DeepEqual(gotOutput, tt.wantOutput) {
				t.Errorf("rowTransfer() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}


func Test_keyExpansion(t *testing.T) {
	tests := []struct {
		name  string
		key   *bits.Bits
		wantW [][]byte
	}{
		{
			name: "1",
			key: bits.NewBitsFromBytes(
				0x3c, 0xa1, 0x0b, 0x21,
				0x57, 0xf0, 0x19, 0x16,
				0x90, 0x2e, 0x13, 0x80,
				0xac, 0xc1, 0x07, 0xbd,
			),
			wantW: [][]byte{
				{0x3c, 0xa1, 0x0b, 0x21},
				{0x57, 0xf0, 0x19, 0x16},
				{0x90, 0x2e, 0x13, 0x80},
				{0xac, 0xc1, 0x07, 0xbd},
				{0x45, 0x64, 0x71, 0xb0},
				{0x12, 0x94, 0x68, 0xa6},
				{0x82, 0xba, 0x7b, 0x26},
				{0x2e, 0x7b, 0x7c, 0x9b},
				{0x66, 0x74, 0x65, 0x81},
				{0x74, 0xe0, 0x0d, 0x27},
				{0xf6, 0x5a, 0x76, 0x01},
				{0xd8, 0x21, 0x0a, 0x9a},
				{0x9f, 0x13, 0xdd, 0xe0},
				{0xeb, 0xf3, 0xd0, 0xc7},
				{0x1d, 0xa9, 0xa6, 0xc6},
				{0xc5, 0x88, 0xac, 0x5c},
				{0x53, 0x82, 0x97, 0x46},
				{0xb8, 0x71, 0x47, 0x81},
				{0xa5, 0xd8, 0xe1, 0x47},
				{0x60, 0x50, 0x4d, 0x1b},
				{0x10, 0x61, 0x38, 0x96},
				{0xa8, 0x10, 0x7f, 0x17},
				{0x0d, 0xc8, 0x9e, 0x50},
				{0x6d, 0x98, 0xd3, 0x4b},
				{0x76, 0x07, 0x8b, 0xaa},
				{0xde, 0x17, 0xf4, 0xbd},
				{0xd3, 0xdf, 0x6a, 0xed},
				{0xbe, 0x47, 0xb9, 0xa6},
				{0x96, 0x51, 0xaf, 0x04},
				{0x48, 0x46, 0x5b, 0xb9},
				{0x9b, 0x99, 0x31, 0x54},
				{0x25, 0xde, 0x88, 0xf2},
				{0x0b, 0x95, 0x26, 0x3b},
				{0x43, 0xd3, 0x7d, 0x82},
				{0xd8, 0x4a, 0x4c, 0xd6},
				{0xfd, 0x94, 0xc4, 0x24},
				{0x32, 0x89, 0x10, 0x6f},
				{0x71, 0x5a, 0x6d, 0xed},
				{0xa9, 0x10, 0x21, 0x3b},
				{0x54, 0x84, 0xe5, 0x1f},
				{0x5b, 0x50, 0xd0, 0x4f},
				{0x2a, 0x0a, 0xbd, 0xa2},
				{0x83, 0x1a, 0x9c, 0x99},
				{0xd7, 0x9e, 0x79, 0x86},
			},
		},
		{
			name: "2",
			key: bits.NewBitsFromString(
				"0f1571c947d9e8590cb7add6af7f6798",
				16,
			),
			wantW: [][]byte{
				{0x0f, 0x15, 0x71, 0xc9},
				{0x47, 0xd9, 0xe8, 0x59},
				{0x0c, 0xb7, 0xad, 0xd6},
				{0xaf, 0x7f, 0x67, 0x98},
				{0xdc, 0x90, 0x37, 0xb0},
				{0x9b, 0x49, 0xdf, 0xe9},
				{0x97, 0xfe, 0x72, 0x3f},
				{0x38, 0x81, 0x15, 0xa7},
				{0xd2, 0xc9, 0x6b, 0xb7},
				{0x49, 0x80, 0xb4, 0x5e},
				{0xde, 0x7e, 0xc6, 0x61},
				{0xe6, 0xff, 0xd3, 0xc6},
				{0xc0, 0xaf, 0xdf, 0x39},
				{0x89, 0x2f, 0x6b, 0x67},
				{0x57, 0x51, 0xad, 0x06},
				{0xb1, 0xae, 0x7e, 0xc0},
				{0x2c, 0x5c, 0x65, 0xf1},
				{0xa5, 0x73, 0x0e, 0x96},
				{0xf2, 0x22, 0xa3, 0x90},
				{0x43, 0x8c, 0xdd, 0x50},
				{0x58, 0x9d, 0x36, 0xeb},
				{0xfd, 0xee, 0x38, 0x7d},
				{0x0f, 0xcc, 0x9b, 0xed},
				{0x4c, 0x40, 0x46, 0xbd},
				{0x71, 0xc7, 0x4c, 0xc2},
				{0x8c, 0x29, 0x74, 0xbf},
				{0x83, 0xe5, 0xef, 0x52},
				{0xcf, 0xa5, 0xa9, 0xef},
				{0x37, 0x14, 0x93, 0x48},
				{0xbb, 0x3d, 0xe7, 0xf7},
				{0x38, 0xd8, 0x08, 0xa5},
				{0xf7, 0x7d, 0xa1, 0x4a},
				{0x48, 0x26, 0x45, 0x20},
				{0xf3, 0x1b, 0xa2, 0xd7},
				{0xcb, 0xc3, 0xaa, 0x72},
				{0x3c, 0xbe, 0x0b, 0x38},
				{0xfd, 0x0d, 0x42, 0xcb},
				{0x0e, 0x16, 0xe0, 0x1c},
				{0xc5, 0xd5, 0x4a, 0x6e},
				{0xf9, 0x6b, 0x41, 0x56},
				{0xb4, 0x8e, 0xf3, 0x52},
				{0xba, 0x98, 0x13, 0x4e},
				{0x7f, 0x4d, 0x59, 0x20},
				{0x86, 0x26, 0x18, 0x76},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotW := keyExpansion(tt.key); !goutils.Equal(gotW, tt.wantW) {
				t.Errorf("keyExpansion() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}

func TestAES(t *testing.T) {
	tests := []struct {
		name       string
		input      string
		key        string
		wantOutput string
	}{
		{
			name:       "sample",
			input:      "0123456789abcdeffedcba9876543210",
			key:        "0f1571c947d9e8590cb7add6af7f6798",
			wantOutput: "ff0b844a0853bf7c6934ab4364148fb9",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutput := AES(tt.input, tt.key); gotOutput != tt.wantOutput {
				t.Errorf("AES() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}