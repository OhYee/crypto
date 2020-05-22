package hmac

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"math/rand"
	"reflect"
	"testing"

	"github.com/OhYee/crypto/hash/sha"
	cmp "github.com/OhYee/goutils/compare"
)

func TestHMAC(t *testing.T) {
	type args struct {
		key     []byte
		message []byte
		hash    func([]byte) []byte
		b       int
	}
	tests := []struct {
		name       string
		args       args
		wantOutput []byte
	}{
		{
			name: "abcdefg",
			args: args{
				key:     []byte("123456"),
				message: []byte("abcdefg"),
				hash:    sha.SHA1,
				b:       64,
			},
			wantOutput: []byte{
				0x04, 0xb0, 0x95, 0x61,
				0xdb, 0x1d, 0x5a, 0xa5,
				0xe8, 0xa6, 0xed, 0x62,
				0x28, 0xd9, 0xf8, 0xa4,
				0xf9, 0x33, 0xad, 0x1e,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutput := HMAC(tt.args.key, tt.args.message, tt.args.hash, tt.args.b); !reflect.DeepEqual(gotOutput, tt.wantOutput) {
				t.Errorf("HMAC() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}

	for i := 0; i < 10; i++ {
		messageL := rand.Intn(10000)
		message := make([]byte, messageL)
		for j := 0; j < messageL; j++ {
			message[j] = uint8(rand.Uint32() % 0xff)
		}

		keyL := rand.Intn(1000)
		key := make([]byte, keyL)
		for j := 0; j < keyL; j++ {
			key[j] = uint8(rand.Uint32() % 0xff)
		}

		t.Run(fmt.Sprintf("Random SHA1 test %d", i), func(t *testing.T) {
			hash := hmac.New(sha1.New, key)
			hash.Write(message)
			want := hash.Sum([]byte{})
			got := HMAC(key, message, sha.SHA1, 64)
			if !cmp.Equal(got, want) {
				t.Errorf("want %+v, got %+v", want, got)
			}
		})

		t.Run(fmt.Sprintf("Random SHA256 test %d", i), func(t *testing.T) {
			hash := hmac.New(sha256.New, key)
			hash.Write(message)
			want := hash.Sum([]byte{})
			got := HMAC(key, message, sha.SHA256, 64)
			if !cmp.Equal(got, want) {
				t.Errorf("want %+v, got %+v", want, got)
			}
		})

		t.Run(fmt.Sprintf("Random SHA512 test %d", i), func(t *testing.T) {
			hash := hmac.New(sha512.New, key)
			hash.Write(message)
			want := hash.Sum([]byte{})
			got := HMAC(key, message, sha.SHA512, 128)
			if !cmp.Equal(got, want) {
				t.Errorf("want %+v, got %+v", want, got)
			}
		})
	}
}
