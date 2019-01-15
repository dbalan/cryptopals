package set2

import (
	//	"fmt"
	"github.com/dbalan/cryptopals/common"
)

func AES128ECBOracle(plainText []byte, key []byte) []byte {
	unknown := common.DecodeB64([]byte("Um9sbGluJyBpbiBteSA1LjAKV2l0aCBteSByYWctdG9wIGRvd24gc28gbXkgaGFpciBjYW4gYmxvdwpUaGUgZ2lybGllcyBvbiBzdGFuZGJ5IHdhdmluZyBqdXN0IHRvIHNheSBoaQpEaWQgeW91IHN0b3A/IE5vLCBJIGp1c3QgZHJvdmUgYnkK"))

	pt := append(plainText, unknown...)
	enc, err := EncAES128ECB(pt, key)
	if err != nil {
		panic(err)
	}
	return enc
}

func detectBlockSize(oracle func([]byte) []byte) int {
	plainText := []byte("A")

	enc := oracle(plainText)
	curLen := len(enc)

	for bs := 2; bs < 32; bs++ {
		plainText := []byte{}
		for i := 0; i < bs*3; i++ {
			plainText = append(plainText, byte('A'))
		}

		enc := oracle(plainText)
		// size of cipherText increases by block
		diff := len(enc) - curLen

		if diff > 0 {
			return diff
		}
	}
	return -1
}

func equal(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}

	for ci, chr := range a {
		if chr != b[ci] {
			return false
		}
	}

	return true
}

func fullDecrypt(oracle func([]byte) []byte) string {
	bs := detectBlockSize(oracle)
	totalbs := len(oracle([]byte(""))) / bs

	prefix := &[]byte{}
	for i := 1; i <= totalbs; i++ {
		p := decryptBlock(oracle, bs, i, *prefix)
		prefix = &p
	}

	return string(*prefix)
}

func decryptBlock(oracle func([]byte) []byte, bs, bno int, prefix []byte) []byte {
dec:
	for round := 0; round < bs; round++ {
		pt := []byte{}

		padlen := (bno * bs) - len(prefix) - 1
		for i := 0; i < padlen; i++ {
			pt = append(pt, byte('A'))
		}

		ctMap := map[byte][]byte{}
		for i := 0; i < 127; i++ {
			newPT := append(pt, prefix...)
			newPT = append(newPT, byte(i))
			//			fmt.Printf("NEWPT: % q (%d)\n", newPT, len(newPT))
			enc := oracle(newPT)
			ctMap[byte(i)] = enc[0 : bs*bno]
		}

		enc := oracle(pt)
		for k, v := range ctMap {
			if equal(v, enc[0:bs*bno]) {
				//				fmt.Printf("Got: k % q\n", k)
				prefix = append(prefix, k)
				continue dec
			}
		}
	}
	return prefix
}
