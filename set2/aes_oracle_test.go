package set2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAESOracle(t *testing.T) {
	pt := []byte("On my arrival at New York the question was at its height.  The theory of the floating island, and the unapproachable sandbank, supported by minds little competent to form a judgment, was abandoned.  And, indeed, unless this shoal had a machine in its stomach, how could it change its position with such astonishing rapidity?")

	for i := 0; i < 10; i++ {
		_, err := EncOracle(pt)
		assert.Nil(t, err)
	}
}
