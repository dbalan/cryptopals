package main

import (
	"crypto/sha256"
	"fmt"
	"math/big"
)

func encodeUint64(x uint64) []byte {
	buf := make([]byte, 8)
	for i := 0; i < 8; i++ {
		buf[i] = byte(x & 0xff)
		x = x >> 8
	}
	return buf
}

func saltHmac(salt uint64, password string) *big.Int {
	encoded := encodeUint64(salt)
	h := sha256.New()
	h.Write(encoded)
	h.Write([]byte(password))
	val := fmt.Sprintf("%x", h.Sum(nil))
	ret := &big.Int{}
	ret.SetString(val, 16)
	return ret
}

func SHA256Int(A ...*big.Int) *big.Int {
	h := sha256.New()
	for _, val := range A {
		h.Write([]byte(val.Text(16)))
	}

	sum := fmt.Sprintf("%x", h.Sum(nil))
	ret := &big.Int{}
	ret.SetString(sum, 16)
	return ret
}
