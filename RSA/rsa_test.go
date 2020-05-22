package rsa

import (
	"fmt"
	"math/big"
	"testing"

	cmp "github.com/OhYee/goutils/compare"
)

func Test_pow(t *testing.T) {
	type args struct {
		a int64
		n int64
		m int64
	}
	tests := []struct {
		name    string
		args    args
		wantRes int64
	}{
		{
			name:    "(2^3) mod 5 = 3",
			args:    args{2, 3, 5},
			wantRes: 3,
		},
		{
			name:    "(2^4) mod 5 = 1",
			args:    args{2, 4, 5},
			wantRes: 1,
		},
		{
			name:    "(1^7) mod 2 = 1",
			args:    args{1, 7, 2},
			wantRes: 1,
		},
		{
			name:    "(100^0) mod 5 = 1",
			args:    args{100, 0, 5},
			wantRes: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := pow(big.NewInt(tt.args.a), big.NewInt(tt.args.n), big.NewInt(tt.args.m)); gotRes.Cmp(big.NewInt(tt.wantRes)) != 0 {
				t.Errorf("pow() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func Test_isPrime(t *testing.T) {
	tests := []struct {
		name string
		n    int64
		want bool
	}{
		{
			name: "2 is prime",
			n:    2,
			want: true,
		},
		{
			name: "5 is prime",
			n:    5,
			want: true,
		},
		{
			name: "16 is not prime",
			n:    16,
			want: false,
		},
		{
			name: "65535 is not prime",
			n:    65535,
			want: false,
		},
		{
			name: "100007 is not prime",
			n:    100007,
			want: false,
		},
		{
			name: "10000007 is not prime",
			n:    10000007,
			want: false,
		},
		{
			name: "100000007 is prime",
			n:    100000007,
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPrime(big.NewInt(tt.n)); got != tt.want {
				t.Errorf("isPrime() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func Test_randomUint16(t *testing.T) {
// 	set := make(map[uint16]bool)
// 	for i := 0; i < 100; i++ {
// 		n := randomUint16()
// 		set[n] = true
// 	}
// 	if len(set) < 100 {
// 		t.Errorf("Not randomly. %d", len(set))
// 	}
// }

func Test_gcd(t *testing.T) {
	type args struct {
		a int64
		b int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"gcd(1,2) = 1", args{1, 2}, 1},
		{"gcd(7,14) = 1", args{7, 14}, 7},
		{"gcd(52,28) = 1", args{52, 28}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gcd(big.NewInt(tt.args.a), big.NewInt(tt.args.b)); got.Cmp(big.NewInt(tt.want)) != 0 {
				t.Errorf("gcd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_exgcd(t *testing.T) {
	type args struct {
		a int64
		b int64
	}
	tests := []struct {
		name  string
		args  args
		wantR int64
		wantX int64
		wantY int64
	}{
		{"exgcd(264,19)", args{264, 19}, 1, 9, -125},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotR, gotX, gotY := exgcd(big.NewInt(tt.args.a), big.NewInt(tt.args.b))
			if gotR.Cmp(big.NewInt(tt.wantR)) != 0 {
				t.Errorf("exgcd() gotR = %v, want %v", gotR, tt.wantR)
			}
			if gotX.Cmp(big.NewInt(tt.wantX)) != 0 {
				t.Errorf("exgcd() gotX = %v, want %v", gotX, tt.wantX)
			}
			if gotY.Cmp(big.NewInt(tt.wantY)) != 0 {
				t.Errorf("exgcd() gotY = %v, want %v", gotY, tt.wantY)
			}
		})
	}
}

func Test_Crypto(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			plainText := "Hello World!"
			publicKey, privateKey, _ := Generate()
			t.Logf("%v %v\n", publicKey, privateKey)
			cryptoText := Encrypto([]byte(plainText), privateKey)
			if got := Decrypto(cryptoText, publicKey); string(got) != plainText {
				t.Errorf("Crypto error: (%v, %v)\n\t%v\n\t%v\n", publicKey, privateKey, cryptoText, got)
			}
		})
	}
}

func Test_Crypto2(t *testing.T) {
	type testcase struct {
		plainText  []byte
		privateKey []byte
		cryptoText []byte
		publicKey  []byte
	}
	testcases := []testcase{
		{
			plainText:  []byte{123},
			privateKey: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x13, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x2b},
			cryptoText: []byte{0x01, 0x18},
			publicKey:  []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x8b, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x2b},
		},
		// This data will cause big.Int = 0, and Bytes() returns []byte{}!
		{
			plainText:  []byte{100, 218, 164, 74, 212, 147, 255, 40, 169, 110, 255, 171, 110, 119, 241, 115, 42, 61, 151, 216, 50, 65, 88, 27, 55, 219, 215, 10, 122, 73, 0, 254},
			privateKey: []byte{0x11, 0x1b, 0x41, 0x0e, 0x3c, 0xcf, 0xa6, 0xb7, 0x13, 0xfe, 0xe1, 0x04, 0xa1, 0xd6, 0x56, 0x29},
			cryptoText: []byte{10, 14, 168, 78, 147, 230, 239, 8, 8, 43, 231, 225, 78, 2, 129, 140, 6, 216, 89, 175, 104, 65, 177, 20, 3, 108, 82, 109, 4, 13, 104, 57, 4, 170, 35, 211, 226, 73, 7, 140, 5, 106, 31, 171, 170, 32, 111, 244, 1, 146, 159, 149, 130, 55, 239, 59, 19, 158, 204, 120, 115, 253, 137, 185, 12, 235, 54, 232, 138, 169, 220, 239, 17, 16, 138, 253, 224, 112, 169, 43, 1, 146, 159, 149, 130, 55, 239, 59, 14, 90, 65, 197, 226, 53, 118, 40, 17, 16, 138, 253, 224, 112, 169, 43, 12, 35, 99, 248, 23, 212, 166, 42, 3, 154, 132, 108, 201, 20, 43, 182, 18, 2, 26, 245, 112, 57, 71, 241, 17, 133, 201, 66, 120, 118, 80, 179, 3, 193, 17, 92, 128, 115, 140, 75, 4, 63, 20, 208, 40, 227, 225, 176, 15, 199, 229, 30, 104, 91, 197, 131, 14, 125, 165, 214, 126, 155, 223, 198, 1, 57, 40, 229, 213, 162, 44, 9, 7, 175, 207, 72, 171, 65, 159, 55, 4, 58, 130, 43, 54, 123, 242, 82, 14, 213, 153, 18, 18, 227, 16, 176, 8, 225, 170, 97, 243, 48, 10, 89, 18, 138, 173, 99, 179, 231, 125, 247, 7, 177, 182, 242, 220, 121, 23, 108, 19, 204, 249, 48, 48, 78, 60, 21, 11, 207, 218, 215, 247, 158, 6, 152, 0, 0, 0, 0, 0, 0, 0, 0, 2, 237, 67, 65, 26, 217, 215, 104},
			publicKey:  []byte{0x00, 0x00, 0x00, 0x00, 0x52, 0x55, 0x3d, 0x07, 0x13, 0xfe, 0xe1, 0x04, 0xa1, 0xd6, 0x56, 0x29},
		},
	}

	for _, tt := range testcases {
		gotCryotoText := Encrypto(tt.plainText, tt.privateKey)
		if !cmp.Equal(gotCryotoText, tt.cryptoText) {
			t.Errorf("Encrypto got %v, want %v\n", gotCryotoText, tt.cryptoText)
		}
		gotPlainText := Decrypto(tt.cryptoText, tt.publicKey)
		if !cmp.Equal(gotPlainText, tt.plainText) {
			t.Errorf("Decrypto got %v, want %v\n", gotPlainText, tt.plainText)
		}
	}
}
