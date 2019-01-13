package set2

import (
	"crypto/aes"
	"github.com/dbalan/cryptopals/common"
)

func aes128ecb(data []byte, key []byte, enc bool) (resp []byte, err error) {
	if len(key) != 16 {
		err = common.BadDataErr
		return
	}

	cipher, err := aes.NewCipher(key)
	if err != nil {
		return
	}

	l := len(data)

	if enc {
		// pad data
		if l%16 != 0 {
			data, _ = PKCS7Padding(data, l+16-(l%16))
		}
	}

	blocksize := 16 // 128 / 8

	for i := 0; i < l; i = i + blocksize {
		ctBlk := data[i : i+blocksize]
		if enc {
			cipher.Encrypt(ctBlk, ctBlk)
		} else {
			cipher.Decrypt(ctBlk, ctBlk)
		}
	}
	if !enc {
		// unpad
		data = PKCS7StripPadding(data)
	}
	return data, nil
}

func EncAES128ECB(pt []byte, key []byte) ([]byte, error) {
	return aes128ecb(pt, key, true)
}

func DecAES128ECB(ct []byte, key []byte) ([]byte, error) {
	return aes128ecb(ct, key, false)
}
