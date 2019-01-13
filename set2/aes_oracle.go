package set2

import (
	"crypto/rand"
	"math/big"
)

func randBytes(size int) ([]byte, error) {
	buf := make([]byte, size)
	_, err := rand.Read(buf)
	return buf, err
}

func EncOracle(pt []byte) ([]byte, error) {
	// key
	blockSize := 16
	key, err := randBytes(blockSize)
	if err != nil {
		return nil, err
	}

	// (0, 6] to 5 - 10 bytes + 5
	padLen, err := rand.Int(rand.Reader, big.NewInt(6))
	if err != nil {
		return nil, err
	}
	pl := padLen.Int64() + 5

	pad, err := randBytes(int(pl))
	if err != nil {
		return nil, err
	}

	pt = append(pt, pad...)
	pt = append(pad, pt...)

	// choose ECB/CBC
	mode, err := rand.Int(rand.Reader, big.NewInt(2))
	if err != nil {
		return nil, err
	}

	if mode.Int64() == 0 {
		return EncAES128ECB(pt, key)
	}

	iv, err := randBytes(blockSize)
	if err != nil {
		return nil, err
	}
	return EncAES128CBC(pt, iv, key)
}
