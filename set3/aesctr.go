package set3

import (
	"crypto/aes"
	//	"fmt"
	"github.com/dbalan/cryptopals/common"
)

func encodeUint64(x uint64) []byte {
	buf := make([]byte, 8)
	for i := 0; i < 8; i++ {
		buf[i] = byte(x & 0xff)
		x = x >> 8
	}
	return buf
}

func AES128CTR(input, key []byte, nonce uint64) ([]byte, error) {
	bs := 16
	lct := len(input)

	upper := encodeUint64(nonce)

	cipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	for b := 0; b < lct; b += bs {

		lower := encodeUint64(uint64(b / bs))
		ksIn := append(upper, lower...)

		kstream := make([]byte, bs)
		cipher.Encrypt(kstream, ksIn)

		currentInput := input[b : b+bs]
		currentInput, err = common.XORBlk(currentInput, kstream)
		if err != nil {
			return nil, err
		}
	}

	return input, nil
}
