package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"encoding/binary"
	"fmt"
	"time"
)

func Totp(secret string, diff uint64, digits int) (code uint32, left uint64, err error) {
	key, err := base32.StdEncoding.DecodeString(secret)
	if err != nil {
		return
	}

	b := make([]byte, 8)
	timeNow := uint64(time.Now().Unix())
	binary.BigEndian.PutUint64(b, timeNow/diff)

	left = diff - (timeNow % diff)

	hmacSha1 := hmac.New(sha1.New, key)
	hmacSha1.Write(b)
	hash := hmacSha1.Sum(nil)

	offset := hash[len(hash)-1] & 0x0F
	hashParts := hash[offset : offset+4]
	hashParts[0] &= 0x7F

	number := binary.BigEndian.Uint32(hashParts)
	mask := uint32(1)
	for digits > 0 {
		digits--
		mask *= 10
	}
	code = number % mask

	return
}

