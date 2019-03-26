package set6

import (
	"fmt"
	"github.com/dbalan/cryptopals/sha"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

func TestChallengeSHA(t *testing.T) {
	// comapare our sha with theirs
	// text := `For those that envy a MC it can be hazardous to your health
	// So be friendly, a matter of life and death, just like a etch-a-sketch
	//`

	val := sha.SHA([]byte(msg))
	hs := new(big.Int).SetBytes(val)
	theirs := new(big.Int)
	theirs.SetString("d2d0714f014a9784047eaeccf956520045c45265", 16)
	assert.Equal(t, 0, theirs.Cmp(hs))
}

func TestBForceDSAKey(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	x := getKey()
	assert.True(t, matchCksum(x))
}

func matchCksum(x *big.Int) bool {
	exp := "0954edd5e0afe5542a4adf012611a91912a3ec16"
	sum := sha.SHA([]byte(x.Text(16)))
	mine := fmt.Sprintf("%x", sum)
	return mine == exp
}
