package set3

import (
	"fmt"
	"github.com/dbalan/cryptopals/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEncryptOracle(t *testing.T) {
	key, err := common.RandBytes(16)
	assert.Nil(t, err)

	iv, err := common.RandBytes(16)
	assert.Nil(t, err)

	pt := getString()
	_ = encryptOracle([]byte(pt), iv, key)

}

func TestDecrypt(t *testing.T) {
	key := []byte("YELLOW SUBMARINE")
	iv := []byte{}
	for i := 0; i < 16; i++ {
		iv = append(iv, byte(0))
	}

	pt := getString()
	decoded := common.DecodeB64([]byte(pt))
	ct := encryptOracle(decoded, iv, key)

	oracle := func(ct, iv []byte) error {
		return decryptOracle(ct, iv, key)
	}

	gotPT := decryptWithPaddingOracle(ct, iv, oracle)
	fmt.Println(string(gotPT))
	assert.Equal(t, common.DecodeB64([]byte(pt)), gotPT)
}
