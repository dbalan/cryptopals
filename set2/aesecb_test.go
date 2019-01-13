package set2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAESECB(t *testing.T) {
	// Twenty Thousand Leagues under the Sea :-)
	pt := "On my arrival at New York the question was at its height.  The theory of the floating island, and the unapproachable sandbank, supported by minds little competent to form a judgment, was abandoned.  And, indeed, unless this shoal had a machine in its stomach, how could it change its position with such astonishing rapidity?"

	key := []byte("YELLOW SUBMARINE")
	enc, err := EncAES128ECB([]byte(pt), key)
	assert.Nil(t, err)

	dec, err := DecAES128ECB(enc, key)
	assert.Nil(t, err)
	assert.Equal(t, pt, string(dec))
}
