package set3

import (
	"github.com/dbalan/cryptopals/common"
	"github.com/dbalan/cryptopals/set2"
	"math/rand"
	"time"
)

func getString() string {
	pts := []string{
		"MDAwMDAwTm93IHRoYXQgdGhlIHBhcnR5IGlzIGp1bXBpbmc=",
		"MDAwMDAxV2l0aCB0aGUgYmFzcyBraWNrZWQgaW4gYW5kIHRoZSBWZWdhJ3MgYXJlIHB1bXBpbic=",
		"MDAwMDAyUXVpY2sgdG8gdGhlIHBvaW50LCB0byB0aGUgcG9pbnQsIG5vIGZha2luZw==",
		"MDAwMDAzQ29va2luZyBNQydzIGxpa2UgYSBwb3VuZCBvZiBiYWNvbg==",
		"MDAwMDA0QnVybmluZyAnZW0sIGlmIHlvdSBhaW4ndCBxdWljayBhbmQgbmltYmxl",
		"MDAwMDA1SSBnbyBjcmF6eSB3aGVuIEkgaGVhciBhIGN5bWJhbA==",
		"MDAwMDA2QW5kIGEgaGlnaCBoYXQgd2l0aCBhIHNvdXBlZCB1cCB0ZW1wbw==",
		"MDAwMDA3SSdtIG9uIGEgcm9sbCwgaXQncyB0aW1lIHRvIGdvIHNvbG8=",
		"MDAwMDA4b2xsaW4nIGluIG15IGZpdmUgcG9pbnQgb2g=",
		"MDAwMDA5aXRoIG15IHJhZy10b3AgZG93biBzbyBteSBoYWlyIGNhbiBibG93",
	}

	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(pts))
	return pts[index]
}

func encryptOracle(pt, key []byte) (ct, iv []byte) {
	bs := 16
	// kraft a random IV
	iv, err := common.RandBytes(bs)
	if err != nil {
		panic(err)
	}

	ct, iv, err = set2.EncAES128CBC(pt, iv, key)
	if err != nil {
		panic(err)
	}
	return
}

func decryptOracle(ct, iv, key []byte) error {
	_, err := set2.DecAES128CBC(ct, iv, key)
	return err
}
