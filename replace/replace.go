package replace

// DefaultKey for replace crypto
const DefaultKey = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

// Caesar cipher key
const Caesar = "DEFGHIJKLMNOPQRSTUVWXYZABCdefghijklmnopqrstuvwxyzabc3456789012"

// CaesarTable cipher key
var CaesarTable = GenerateReplaceTable("DEFGHIJKLMNOPQRSTUVWXYZABCdefghijklmnopqrstuvwxyzabc3456789012", 0)

// GenerateReplaceTable from key
func GenerateReplaceTable(keyString string, offset int) (key map[rune]rune) {
	if len(keyString)+offset > len(DefaultKey) {
		panic("Generate replace table error: Key is too long")
	}
	key = make(map[rune]rune)
	for idx := range DefaultKey {
		if idx >= len(keyString) {
			return
		}
		key[rune(DefaultKey[idx+offset])] = rune(keyString[idx])
	}
	return
}

// GenerateDecryptionKey using encryption key
func GenerateDecryptionKey(key map[rune]rune) (dKey map[rune]rune) {
	dKey = make(map[rune]rune)
	for k, v := range key {
		dKey[v] = k
	}
	return
}

// ReplacePassword using key
func ReplacePassword(plainText string, key map[rune]rune) (secretText string) {
	secretSlice := make([]rune, len(plainText))
	var exist bool

	for idx, char := range plainText {
		if secretSlice[idx], exist = key[char]; !exist {
			secretSlice[idx] = char
		}
	}
	secretText = string(secretSlice)
	return
}

func main() {
	return
}
