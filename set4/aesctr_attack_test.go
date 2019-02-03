package set4

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChallengePT(t *testing.T) {
	_, err := challengePT()
	assert.Nil(t, err)
}

func TestEncryptCTR(t *testing.T) {
	_, _, _, err := encryptCTR()
	assert.Nil(t, err)
}

func TestAESCTRAttack(t *testing.T) {
	ct, key, nonce, err := encryptCTR()
	assert.Nil(t, err)

	oracle := func(c, k []byte, o uint64, nt []byte) ([]byte, error) {
		return edit(c, k, nonce, o, nt)
	}

	pt, err := AESCTRAttack(oracle, ct, key)
	assert.Nil(t, err)

	assert.NotEqual(t, 0, len(pt))
	fmt.Println(string(pt))
}
