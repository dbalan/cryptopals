package set2

import (
	"github.com/dbalan/cryptopals/common"
)

func DetectAESMode(oracle func([]byte) ([]byte, error)) common.AESMode {
	// look for repeat data?
	plainText := []byte{}
	for i := 0; i < 16*3; i++ {
		plainText = append(plainText, byte('a'))
	}

	enc, err := oracle(plainText)
	if err != nil {
		panic("err")
	}

	if findRepeations(enc) {
		return common.ECB
	}
	return common.CBC
}

// IN ECB MODE - middle blocks would be same
func findRepeations(cipherText []byte) bool {
	start := 16
	bs := 16
	fst := cipherText[start : start+bs]
	start = start + bs
	snd := cipherText[start : start+bs]

	for k, f := range fst {
		if f != snd[k] {
			return false
		}
	}
	return true
}
