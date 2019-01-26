package set3

import (
	"fmt"
	//	"github.com/dbalan/cryptopals/common"
	"github.com/dbalan/cryptopals/set2"
	"math/rand"
	"strings"
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

func encryptOracle(pt, iv, key []byte) []byte {
	//	bs := 16
	// kraft a random IV
	/*	iv, err := common.RandBytes(bs)
			if err != nil {
		panic(err)
		}
	*/
	ct, err := set2.EncAES128CBC(pt, iv, key)
	if err != nil {
		panic(err)
	}
	return ct
}

func decryptOracle(ct, iv, key []byte) error {
	_, err := set2.DecAES128CBC(ct, iv, key)
	return err
}

func decryptWithPaddingOracle(ct, iv []byte, oracle func([]byte, []byte) error) {
	pt := []byte{}
	// we could do this over 1 round, however, what if the oracle rejects
	// padding > bs?
	for numBlocks := int(len(ct) / 16); numBlocks > 1; numBlocks-- {
		p := decryptLastBlock(ct, iv, oracle)
		pt = append(p, pt...)
		ct = ct[0 : (numBlocks-1)*16]
	}
	p := decryptLastBlock(ct, iv, oracle)
	pt = append(p, pt...)
	fmt.Println(strings.TrimSpace(string(pt)))
}

func decryptLastBlock(ct, iv []byte, oracle func([]byte, []byte) error) []byte {
	bs := 16
	lenct := len(ct)

	aesDec := make([]byte, bs)
	ptBlock := make([]byte, bs)
outer:
	for curbit := 15; curbit >= 0; curbit-- {
		pad := 16 - curbit

		// craft a new CT
		nct := make([]byte, lenct)
		copy(nct, ct)

		var xorblk []byte
		// if this is the lastblock
		if lenct == bs {
			xorblk = iv[0:16]
		} else {
			// the block that gets xor'd with decrypted output
			xorblk = nct[lenct-32 : lenct-16]
		}

		for i := 0; i < pad-1; i++ {
			xorblk[15-i] = aesDec[15-i] ^ byte(pad)
		}

		orig := xorblk[16-pad]
		for b := 0; b <= 0xff; b++ {
			xorblk[16-pad] = byte(b)

			// another copy of ciphertext
			// oracle messes with the buffer passed to it.
			nnct := make([]byte, lenct)
			copy(nnct, nct)
			if oracle(nnct, iv) == nil {
				ad := byte(b) ^ byte(pad)
				aesDec[16-pad] = ad
				ptBlock[16-pad] = ad ^ orig
				continue outer
			}
		}

		panic("should not")
	}

	return set2.PKCS7StripPadding(ptBlock)
}
