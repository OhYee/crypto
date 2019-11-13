package hmac

import ()

func HMAC(key []byte, message []byte, hash func([]byte) []byte, b int) (output []byte) {
	if len(key) >= b {
		key = hash(key)
	}
	keyPlus := make([]byte, b)
	copy(keyPlus, key)

	ipad := make([]byte, b)
	opad := make([]byte, b)
	for i := 0; i < b; i++ {
		ipad[i] = keyPlus[i] ^ 0x36
		opad[i] = keyPlus[i] ^ 0x5c
	}

	output = hash(append(opad, hash(append(ipad, message...))...))
	return
}
