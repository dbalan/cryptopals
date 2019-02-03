package set4

import (
	"bytes"
	"fmt"
	"github.com/dbalan/cryptopals/common"
	"github.com/dbalan/cryptopals/set2"
	"github.com/dbalan/cryptopals/set3"
	"io/ioutil"
	"math/rand"
)

func challengePT() ([]byte, error) {
	data, err := ioutil.ReadFile("./25.txt")
	data = bytes.Replace(data, []byte("\n"), []byte(""), -1)
	if err != nil {
		return nil, err
	}

	decoded := common.DecodeB64(data)
	key := []byte("YELLOW SUBMARINE")

	return set2.DecAES128ECB(decoded, key)
}

func encryptCTR() (ct []byte, key []byte, nonce uint64, err error) {
	var pt []byte
	pt, err = challengePT()
	if err != nil {
		return
	}

	key, err = common.RandBytes(16)
	if err != nil {
		return
	}
	nonce = rand.Uint64()

	ct, err = set3.AES128CTR(pt, key, nonce)
	return
}

func edit(ct []byte, key []byte, nonce uint64, offset uint64, newtext []byte) ([]byte, error) {
	pt, err := set3.AES128CTR(ct, key, nonce)
	if err != nil {
		return nil, err
	}

	if int(offset)+len(newtext) > len(ct) {
		return nil, fmt.Errorf("Bad offset")
	}

	for i := 0; i < len(newtext); i++ {
		pt[int(offset)+i] = newtext[i]
	}

	return set3.AES128CTR(pt, key, nonce)
}

func AESCTRAttack(oracle func([]byte, []byte, uint64, []byte) ([]byte, error),
	ct, key []byte) ([]byte, error) {
	lct := len(ct)
	origCT := make([]byte, lct)
	copy(origCT, ct)

	forged := []byte{}
	for i := 0; i < lct; i++ {
		forged = append(forged, byte('A'))
	}

	ct, err := oracle(ct, key, 0, forged)
	if err != nil {
		return nil, err
	}

	keyStream, err := common.XORBlk(ct, forged)
	if err != nil {
		return nil, err
	}

	return common.XORBlk(keyStream, origCT)
}
