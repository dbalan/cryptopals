package set3

import (
	"fmt"
	"github.com/dbalan/cryptopals/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEncryptInput(t *testing.T) {
	key, err := common.RandBytes(16)
	assert.Nil(t, err)
	_ = encryptChallengeText(key)
}

func TestCH20(t *testing.T) {
	resp := CH20()

	fmt.Printf("CH20 : %s\n", string(resp))
}
