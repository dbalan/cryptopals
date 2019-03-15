package set6

import (
	"github.com/dbalan/cryptopals/sha"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

func TestChallengeSHA(t *testing.T) {
	// comapare our sha with theirs
	text := `For those that envy a MC it can be hazardous to your health
So be friendly, a matter of life and death, just like a etch-a-sketch
`

	val := sha.SHA([]byte(text))
	hs := new(big.Int).SetBytes(val)
	theirs := new(big.Int)
	theirs.SetString("d2d0714f014a9784047eaeccf956520045c45265", 16)
	assert.Equal(t, 0, theirs.Cmp(hs))
}
