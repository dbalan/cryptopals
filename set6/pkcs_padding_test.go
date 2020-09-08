package set6

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPKCSOracle(t *testing.T) {
	ct, _, _, oracle := PKCSOracle()
	assert.True(t, oracle(ct), "unchanged ct should give you pkcs conforming")
}

func TestPKCSPaddingAttack(t *testing.T) {
	ct, pub, n, oracle := PKCSOracle()
	PKCSOracleAttack(ct, pub, n, oracle)
}
