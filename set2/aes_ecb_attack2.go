package set2

import (
	"github.com/dbalan/cryptopals/common"
	"math"
)

func AESECBOracle2(plainText []byte, key []byte, prefix []byte) []byte {
	targetText := common.DecodeB64([]byte("Um9sbGluJyBpbiBteSA1LjAKV2l0aCBteSByYWctdG9wIGRvd24gc28gbXkgaGFpciBjYW4gYmxvdwpUaGUgZ2lybGllcyBvbiBzdGFuZGJ5IHdhdmluZyBqdXN0IHRvIHNheSBoaQpEaWQgeW91IHN0b3A/IE5vLCBJIGp1c3QgZHJvdmUgYnkK"))

	pt := append(prefix, plainText...)
	pt = append(pt, targetText...)

	enc, err := EncAES128ECB(pt, key)
	if err != nil {
		panic(err)
	}

	return enc
}

func findPfxLen(oracle func([]byte) []byte) int {
	// step 0: find ct for a known block
	// get 4 blocks of known pt
	kpt := ""
	for i := 0; i < 3*16; i++ {
		kpt += "A"
	}

	bs := 16
	enc := oracle([]byte(kpt))

	blks := common.Blocks(enc, bs)

	eqblk := []byte{}
	// find first identical blocks in enc
	// FIXME: we are assuming prefix doesnt repeat itself, but we can solve
	// that by couting repeatiions
outer:
	for i := 0; i < len(blks); i++ {
		for j := i + 1; j < len(blks); j++ {
			if common.EqualBlocks(blks[i], blks[j]) {
				eqblk = blks[i]
				break outer
			}
		}
	}

	// we know how a block of A would look like in ciphertext
	// step 1: take a block of A, and increase the length untill we see the
	// block
	kpt = ""

	for i := 0; i < 16; i++ {
		kpt += "A"
	}

	for i := 0; i < 16; i++ {
		enc = oracle([]byte(kpt))

		for bi, b := range common.Blocks(enc, bs) {
			if common.EqualBlocks(eqblk, b) {
				// we know block is at bi-th place
				return ((bi - 1) * 16) + (16 - i)
			}
		}

		kpt += "A"
	}

	return -1
}

func decryptPrefixOracle(oracle func([]byte) []byte) []byte {
	bs := 16
	pfxLen := findPfxLen(oracle)
	enc := oracle([]byte(""))
	nblocks := len(enc) / 16

	dec := []byte{}

	pBlkLen := 1
	if pfxLen > bs {
		// how many blocks of ciphertext is prefix?
		pBlkLen = int(math.Ceil(float64(pfxLen) / float64(bs)))
	}

	// we pad first block away to prefix
	// from second ownwards
	for i := pBlkLen; i <= nblocks; i++ {
		dec = decryptPfxBlock(oracle, pfxLen, i, dec)
	}
	return dec
}

// decryptfirst
func decryptPfxBlock(oracle func([]byte) []byte, prefixLen int, bno int, decrypted []byte) []byte {
	bs := 16
	padLen := bs - (prefixLen % bs)

	startIndex := bno * bs
	//	if prefixLen > bs {
	//		startIndex = int(prefixLen/bs) * bs
	//	}

outer:
	for round := 0; round < bs; round++ {
		pt := []byte{}

		for i := 0; i < padLen; i++ {
			pt = append(pt, byte('A'))
		}

		for i := 0; i < (bs*bno)-len(decrypted)-1; i++ {
			pt = append(pt, byte('A'))
		}

		ctMap := map[byte][]byte{}
		//		fmt.Printf("CURRENT: bno :%d padLen: %d ptlen :%d pt %s\n", bno, padLen, len(pt), pt)
		for i := 0; i < 127; i++ {
			newPT := append(pt, decrypted...)
			newPT = append(newPT, byte(i))
			//			fmt.Printf("TRYING FOR PT: %s\n", string(newPT))
			enc := oracle(newPT)
			ctMap[byte(i)] = enc[startIndex : startIndex+bs]
		}

		enc := oracle(pt)
		for k, v := range ctMap {
			if common.EqualBlocks(v, enc[startIndex:startIndex+bs]) {
				decrypted = append(decrypted, k)
				continue outer
			}
		}
	}

	return decrypted
}
