package set2

import (
	"crypto/rand"
	"github.com/dbalan/cryptopals/common"
	"math/big"
)

func EncOracle(pt []byte) ([]byte, common.AESMode, error) {
	// key
	blockSize := 16
	key, err := common.RandBytes(blockSize)
	if err != nil {
		return nil, -1, err
	}

	// (0, 6] to 5 - 10 bytes + 5
	padLen, err := rand.Int(rand.Reader, big.NewInt(6))
	if err != nil {
		return nil, -1, err
	}
	pl := padLen.Int64() + 5

	pad, err := common.RandBytes(int(pl))
	if err != nil {
		return nil, -1, err
	}

	pt = append(pt, pad...)
	pt = append(pad, pt...)

	// choose ECB/CBC
	mode, err := rand.Int(rand.Reader, big.NewInt(2))
	if err != nil {
		return nil, -1, err
	}

	if mode.Int64() == 0 {
		enc, err := EncAES128ECB(pt, key)
		return enc, common.ECB, err
	}

	iv, err := common.RandBytes(blockSize)
	if err != nil {
		return nil, -1, err
	}
	enc, err := EncAES128CBC(pt, iv, key)
	return enc, common.CBC, err
}
