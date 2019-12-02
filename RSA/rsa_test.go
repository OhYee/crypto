package rsa

import (
	"fmt"
	"github.com/OhYee/goutils"
	"math/big"
	"testing"
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
			cryptoText := Encrypto([]byte(plainText), privateKey)
			if got := Decrypto(cryptoText, publicKey); string(got) != plainText {
				t.Errorf("Crypto error: (%v, %v)\n\t%v\n\t%v\n", publicKey, privateKey, cryptoText, got)
			}
		})
	}
}

func Test_Crypto2(t *testing.T) {
	plainText := []byte{123}
	privateKey := []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x13, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x2b}
	cryptoText := []byte{0x01, 0x18}
	publicKey := []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x8b, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x2b}
	gotCryotoText := Encrypto(plainText, privateKey)
	if !goutils.Equal(gotCryotoText, cryptoText) {
		t.Errorf("Encrypto got %v, want %v\n", gotCryotoText, cryptoText)
	}
	gotPlainText := Decrypto(gotCryotoText, publicKey)
	if !goutils.Equal(gotPlainText, plainText) {
		t.Errorf("Decrypto got %v, want %v\n", gotPlainText, plainText)
	}
}
