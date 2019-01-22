package set2

import (
	"crypto/rand"
	"fmt"
	"github.com/dbalan/cryptopals/common"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

func TestFindPfxLen(t *testing.T) {
	// random prefix
	count, err := rand.Int(rand.Reader, big.NewInt(10))
	assert.Nil(t, err)
	prefix, err := common.RandBytes(int(count.Int64()))
	assert.Nil(t, err)

	key, err := common.RandBytes(16)
	assert.Nil(t, err)
	oracle := func(pt []byte) []byte {
		return AESECBOracle2(pt, key, prefix)
	}

	// find prefixpad
	pfxLen := findPfxLen(oracle)
	assert.Equal(t, len(prefix), pfxLen)
}

func TestDecrypt2(t *testing.T) {
	if testing.Short() {
		t.Skip("too long")
	}
	// random prefix
	count, err := rand.Int(rand.Reader, big.NewInt(10))
	assert.Nil(t, err)
	prefix, err := common.RandBytes(int(count.Int64()))
	assert.Nil(t, err)

	key, err := common.RandBytes(16)
	assert.Nil(t, err)
	oracle := func(pt []byte) []byte {
		return AESECBOracle2(pt, key, prefix)
	}

	decrypted := decryptPrefixOracle(oracle)
	fmt.Printf("DECRYPTED: %s\n", decrypted)
}
