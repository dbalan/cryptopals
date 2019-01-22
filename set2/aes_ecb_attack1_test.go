package set2

import (
	"fmt"
	"github.com/dbalan/cryptopals/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDetectBlockSize(t *testing.T) {
	//	key, err := common.RandBytes(16)
	//	assert.Nil(t, err)
	key := []byte("YELLOW SUBMARINE")
	oracle := func(pt []byte) []byte {
		return AES128ECBOracle(pt, key)
	}
	assert.Equal(t, 16, detectBlockSize(oracle))
}

func TestAES128ECBOracle(t *testing.T) {
	// for a constant key, check block size changes
	key, err := common.RandBytes(16)
	assert.Nil(t, err)

	pt := "A"

	bs := detectBlockSize(func(pt []byte) []byte {
		return AES128ECBOracle(pt, key)
	})

	enc := AES128ECBOracle([]byte(pt), key)

	// bs - (len(enc)+bs % bs) + len(enc)+bs
	curLen := len(enc)

	for i := 0; i < bs-1; i++ {
		pt += "A"
	}

	enc2 := AES128ECBOracle([]byte(pt), key)

	newLen := len(enc2)
	assert.Equal(t, bs, newLen-curLen)
}

func TestDecrypt(t *testing.T) {
	if testing.Short() {
		t.Skip("Too Long!")
	}
	key, err := common.RandBytes(16)
	assert.Nil(t, err)

	oracle := func(pt []byte) []byte {
		return AES128ECBOracle(pt, key)
	}
	resp := fullDecrypt(oracle)
	fmt.Println(resp)
}
