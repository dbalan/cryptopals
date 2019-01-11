package set1

import (
	"fmt"
	"testing"
)

func TestCh3(t *testing.T) {
	ciphertexts := []string{"1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736", "7b5a4215415d544115415d5015455447414c155c46155f4058455c5b523f"}

	for _, ct := range ciphertexts {
		raw := decodeHexString(ct)
		pt, key := BestPT(raw)
		fmt.Printf("CT: %s\nPT: %s\nKey: %x\n", ct, pt, int(key))
	}
}
