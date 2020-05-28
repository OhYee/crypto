package main

import (
	"fmt"
	"syscall/js"

	des "github.com/OhYee/crypto/DES"
	"github.com/OhYee/crypto/DES/bits"
	"github.com/OhYee/crypto/base64"
	"github.com/OhYee/crypto/hash/sha"
	"github.com/OhYee/crypto/hmac"
	"github.com/OhYee/crypto/replace"
	"github.com/OhYee/crypto/totp"
	"github.com/OhYee/goutils/bytes"
	"github.com/OhYee/rainbow/errors"
	pkg "github.com/OhYee/wasm/package"
)

func readBytes(args []js.Value) ([]byte, error) {
	if len(args) >= 1 {
		input := make([]byte, args[0].Length())
		js.CopyBytesToGo(input, args[0])
		return input, nil
	}
	return []byte{}, errors.New("Argument must be a Uint8Array")
}

func readString(args []js.Value) (string, error) {
	if len(args) >= 1 && args[0].Type() == js.TypeString {
		return args[0].String(), nil
	}
	return "", errors.New("Argument must be a string")
}

func hash(algorithm func([]byte) []byte) func(this js.Value, args []js.Value) (interface{}, error) {
	return func(this js.Value, args []js.Value) (interface{}, error) {
		if len(args) >= 1 && args[0].Type() == js.TypeObject {
			input := make([]byte, args[0].Length())
			js.CopyBytesToGo(input, args[0])

			b := algorithm(input)

			dst := js.Global().Get("Uint8Array").New(len(b))
			js.CopyBytesToJS(dst, b)

			return dst, nil
		}
		return []byte{}, errors.New("Argument must be a Uint8Array")
	}

}

//go:generate bash -c "GOARCH=wasm GOOS=js go build -o ../../blotter_page/public/static/crypto.wasm main.go"
func main() {
	des.Logger.SetOutputToStdout()

	crypto := pkg.NewPackage("ohyee_crypto")
	crypto.ExportVar("defaultBase64Key", js.ValueOf(base64.DefaultBase64))
	crypto.ExportFunction("base64", func(this js.Value, args []js.Value) (interface{}, error) {
		if len(args) >= 1 {
			input := make([]byte, args[0].Length())
			js.CopyBytesToGo(input, args[0])
			key := ""
			if len(args) == 2 && args[1].Type() == js.TypeString {
				key = args[1].String()
			}
			return base64.Base64(input, []rune(key)...), nil
		}
		return []byte{}, errors.New("Argument must be a Uint8Array")
	})
	crypto.ExportFunction("debase64", func(this js.Value, args []js.Value) (interface{}, error) {
		if len(args) >= 1 && args[0].Type() == js.TypeString {
			key := ""
			if len(args) == 2 && args[1].Type() == js.TypeString {
				key = args[1].String()
			}
			b := base64.DeBase64(args[0].String(), []rune(key)...)
			dst := js.Global().Get("Uint8Array").New(len(b))
			js.CopyBytesToJS(dst, b)
			return dst, nil
		}
		return nil, errors.New("Argument must be a string")
	})
	crypto.ExportFunction("sha1", hash(sha.SHA1))
	crypto.ExportFunction("sha224", hash(sha.SHA224))
	crypto.ExportFunction("sha256", hash(sha.SHA256))
	crypto.ExportFunction("sha384", hash(sha.SHA384))
	crypto.ExportFunction("sha512", hash(sha.SHA512))
	crypto.ExportFunction("encrypto_replace", func(this js.Value, args []js.Value) (interface{}, error) {
		if len(args) >= 2 &&
			args[0].Type() == js.TypeString &&
			args[1].Type() == js.TypeString {
			plaintext := args[0].String()
			key := args[1].String()
			return replace.ReplacePassword(plaintext, replace.GenerateReplaceTable(key, 0)), nil
		}
		fmt.Println("error")
		return nil, errors.New("Want 2 string arguments")
	})
	crypto.ExportFunction("decrypto_replace", func(this js.Value, args []js.Value) (interface{}, error) {
		if len(args) >= 2 &&
			args[0].Type() == js.TypeString &&
			args[1].Type() == js.TypeString {
			plaintext := args[0].String()
			key := args[1].String()
			return replace.ReplacePassword(plaintext, replace.GenerateDecryptionKey(replace.GenerateReplaceTable(key, 0))), nil
		}
		return nil, errors.New("Want 2 string arguments")
	})
	crypto.ExportVar("caesar", js.ValueOf(replace.Caesar))
	crypto.ExportFunction("hmac", func(this js.Value, args []js.Value) (interface{}, error) {
		if len(args) >= 3 && args[2].Type() == js.TypeString {
			plaintext := make([]byte, args[0].Length())
			key := make([]byte, args[1].Length())
			js.CopyBytesToGo(plaintext, args[0])
			js.CopyBytesToGo(key, args[1])
			hash := args[2].String()

			var algorithm func([]byte) []byte
			length := 0

			switch hash {
			case "sha1":
				algorithm = sha.SHA1
				length = 64
			case "sha224":
				algorithm = sha.SHA224
				length = 64
			case "sha256":
				algorithm = sha.SHA256
				length = 64
			case "sha384":
				algorithm = sha.SHA384
				length = 64
			case "sha512":
				algorithm = sha.SHA512
				length = 128
			}
			if length != 0 {
				b := hmac.HMAC(key, plaintext, algorithm, length)
				dst := js.Global().Get("Uint8Array").New(len(b))
				js.CopyBytesToJS(dst, b)
				return dst, nil
			}
			return nil, errors.New("Can not use hash algorithm %s", hash)
		}
		return nil, errors.New("Want Uint8Array, Uint8Array, string(hash algorithm) arguments")
	})
	crypto.ExportFunction("totp", func(this js.Value, args []js.Value) (interface{}, error) {
		if len(args) >= 3 &&
			args[0].Type() == js.TypeObject &&
			args[1].Type() == js.TypeNumber &&
			args[2].Type() == js.TypeNumber {
			key := make([]byte, args[0].Length())
			js.CopyBytesToGo(key, args[0])
			diff := args[1].Int()
			digits := args[2].Int()

			code, left, err := totp.Totp(key, uint64(diff), digits)
			if err != nil {
				return nil, err
			}

			return js.ValueOf(map[string]interface{}{
				"code": code,
				"left": left,
			}), nil
		}
		return nil, errors.New("Want Uint8Array, number, number arguments")
	})
	crypto.ExportFunction("des_encrypto", func(this js.Value, args []js.Value) (interface{}, error) {
		if len(args) >= 2 &&
			args[0].Type() == js.TypeObject &&
			args[1].Type() == js.TypeObject {
			input := make([]byte, args[0].Length())
			key := make([]byte, args[1].Length())

			js.CopyBytesToGo(input, args[0])
			js.CopyBytesToGo(key, args[1])

			return bytes.FromUint64(
				uint64(des.Encrypto(
					bits.Bits(bytes.ToUint64(input)),
					bits.Bits(bytes.ToUint64(key)),
				)),
			), nil
		}
		return nil, errors.New("Argument must be Uint8Array, Uint8Array")
	})
	crypto.ExportFunction("des_decrypto", func(this js.Value, args []js.Value) (interface{}, error) {
		if len(args) >= 2 &&
			args[0].Type() == js.TypeObject &&
			args[1].Type() == js.TypeObject {
			input := make([]byte, args[0].Length())
			key := make([]byte, args[1].Length())

			js.CopyBytesToGo(input, args[0])
			js.CopyBytesToGo(key, args[1])

			return bytes.FromUint64(
				uint64(des.Decrypto(
					bits.Bits(bytes.ToUint64(input)),
					bits.Bits(bytes.ToUint64(key)),
				)),
			), nil
		}
		return nil, errors.New("Argument must be Uint8Array, Uint8Array")
	})
	crypto.Run()
}
