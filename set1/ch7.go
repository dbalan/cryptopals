package set1

import (
	"crypto/aes"
	//	"fmt"
	"github.com/pkg/errors"
)

func DecryptAES128ECB(ct []byte, key []byte) (pt []byte, err error) {
	if len(key) != 16 {
		err = errors.New("BAD_KEY")
		return
	}

	cipher, err := aes.NewCipher(key)
	if err != nil {
		return
	}

	l := len(ct)
	blocksize := 16 // 128 / 8

	for i := 0; i < l; i = i + blocksize {
		ctBlk := ct[i : i+blocksize]
		cipher.Decrypt(ctBlk, ctBlk)
	}
	return ct, nil
}
