package replace

import (
	"reflect"
	"testing"
)

func TestGenerateReplaceTable(t *testing.T) {
	tests := []struct {
		name      string
		keyString string
		offset    int
		wantKey   map[rune]rune
	}{
		{
			name:      "Caesar cipher",
			keyString: "DEFGHIJKLMNOPQRSTUVWXYZABCdefghijklmnopqrstuvwxyzabc3456789012",
			offset:    0,
			wantKey: map[rune]rune{
				'A': 'D', 'B': 'E', 'C': 'F', 'D': 'G', 'E': 'H',
				'F': 'I', 'G': 'J', 'H': 'K', 'I': 'L', 'J': 'M',
				'K': 'N', 'L': 'O', 'M': 'P', 'N': 'Q', 'O': 'R',
				'P': 'S', 'Q': 'T', 'R': 'U', 'S': 'V', 'T': 'W',
				'U': 'X', 'V': 'Y', 'W': 'Z', 'X': 'A', 'Y': 'B',
				'Z': 'C', 'a': 'd', 'b': 'e', 'c': 'f', 'd': 'g',
				'e': 'h', 'f': 'i', 'g': 'j', 'h': 'k', 'i': 'l',
				'j': 'm', 'k': 'n', 'l': 'o', 'm': 'p', 'n': 'q',
				'o': 'r', 'p': 's', 'q': 't', 'r': 'u', 's': 'v',
				't': 'w', 'u': 'x', 'v': 'y', 'w': 'z', 'x': 'a',
				'y': 'b', 'z': 'c', '0': '3', '1': '4', '2': '5',
				'3': '6', '4': '7', '5': '8', '6': '9', '7': '0',
				'8': '1', '9': '2',
			},
		},
		{
			name:      "Number replace",
			keyString: "1357924680",
			offset:    26 * 2,
			wantKey: map[rune]rune{
				'0': '1', '1': '3', '2': '5', '3': '7', '4': '9',
				'5': '2', '6': '4', '7': '6', '8': '8', '9': '0',
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotKey := GenerateReplaceTable(tt.keyString, tt.offset); !reflect.DeepEqual(gotKey, tt.wantKey) {
				t.Errorf("GenerateReplaceTable() = %v, want %v", gotKey, tt.wantKey)
			}
		})
	}
}

func TestReplacePassword(t *testing.T) {
	tests := []struct {
		name           string
		plainText      string
		key            map[rune]rune
		wantSecretText string
	}{
		{
			name:           "Caesar cipher",
			plainText:      "Caesar Cipher",
			key:            CaesarTable,
			wantSecretText: "Fdhvdu Flskhu",
		},
		{
			name:           "Decrypto Caesar cipher",
			plainText:      "Fdhvdu Flskhu",
			key:            GenerateDecryptionKey(CaesarTable),
			wantSecretText: "Caesar Cipher",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSecretText := ReplacePassword(tt.plainText, tt.key); gotSecretText != tt.wantSecretText {
				t.Errorf("ReplacePassword() = %v, want %v", gotSecretText, tt.wantSecretText)
			}
		})
	}
}
