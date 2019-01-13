package set2

import (
	"crypto/aes"
	"github.com/dbalan/cryptopals/common"
	"math"
)

func EncAES128CBC(pt []byte, iv []byte, key []byte) ([]byte, error) {
	blockSize := 16 // 128/8

	cipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// pad data!
	if len(pt)%16 != 0 {
		lento := math.Ceil(float64(len(pt)) / 16.0)
		pt, err = PKCS7Padding(pt, int(lento))
		if err != nil {
			return nil, err

		}
	}

	prevCT := iv

	// FIXME: assumes pt doesn't need to be padded
	for i := 0; i < len(pt); i += blockSize {
		curBlk := pt[i : i+blockSize]
		curBlk, _ = common.XORBlk(curBlk, prevCT)
		// encrypt this
		cipher.Encrypt(curBlk, curBlk)
		prevCT = curBlk
	}

	return pt, nil
}

func DecAES128CBC(ct []byte, iv []byte, key []byte) ([]byte, error) {
	blockSize := 16

	cipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	prevCT := iv
	for i := 0; i < len(ct); i += blockSize {
		curBlk := ct[i : i+blockSize]

		// store cipher text for next iteration
		tmpBlk := make([]byte, blockSize)
		copy(tmpBlk, curBlk)

		// decrypt in place
		cipher.Decrypt(curBlk, curBlk)
		curBlk, _ = common.XORBlk(curBlk, prevCT)
		prevCT = tmpBlk
	}

	return PKCS7StripPadding(ct), nil
}
