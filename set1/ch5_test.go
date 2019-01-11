package set1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRepeatingXOR(t *testing.T) {
	plainText := `Burning 'em, if you ain't quick and nimble
I go crazy when I hear a cymbal`
	key := "ICE"
	cipherText := "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"

	encrypted := repeatingXOR(plainText, key)
	assert.Equal(t, cipherText, encrypted)

}

func TestRepId(t *testing.T) {
	pt := "random"
	key := "ran"

	ct := repeatingXOR(pt, key)
	pt2 := decryptRepeatXOR(decodeHexString(ct), []byte(key))
	assert.Equal(t, []byte(pt), pt2)
}
