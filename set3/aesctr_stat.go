package set3

import (
	"bufio"
	"bytes"
	"github.com/dbalan/cryptopals/common"
	"github.com/dbalan/cryptopals/set1"
	"os"
)

// reads input strings, generates ciphertexts
func encryptChallengeText(key []byte) (resp [][]byte) {
	fp, err := os.Open("./20.txt")
	if err != nil {
		panic("should not have happened!")
	}

	scanner := bufio.NewScanner(fp)

	for scanner.Scan() {
		input := common.DecodeB64(scanner.Bytes())
		enc, err := AES128CTR(input, key, uint64(0))

		if err != nil {
			panic(err)
		}
		resp = append(resp, enc)
	}
	return
}

func stripToShortest(input [][]byte) (merged []byte, keysize int) {
	if len(input) == 0 {
		panic(common.BadDataErr)
	}

	// shortest block, also keysize (for XOR)
	keysize = len(input[0])

	for _, in := range input {
		if len(in) < keysize {
			keysize = len(in)
		}
	}

	for _, in := range input {
		merged = append(merged, in[0:keysize]...)
	}
	return
}

func decrypt(input []byte, keysize int) []byte {
	key := set1.FindKey(input, keysize)
	dec := set1.DecryptRepeatXOR(input, key)
	return dec
}

func CH20() []byte {
	key, err := common.RandBytes(16)
	if err != nil {
		panic(err)
	}

	data := encryptChallengeText(key)
	dec := decrypt(stripToShortest(data))
	return bytes.Replace(dec, []byte("/"), []byte("\n"), -1)
}
