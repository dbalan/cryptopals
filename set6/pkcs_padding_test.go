package set6

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const pt = "Dora: We did it"

func TestPKCSOracle(t *testing.T) {
	msg := []byte(pt)
	ct, _, _, oracle := PKCSOracle(msg)
	assert.True(t, oracle(ct), "unchanged ct should give you pkcs conforming")
}

func TestPKCSPaddingAttack(t *testing.T) {
	msg := []byte(pt)
	ct, pub, n, oracle := PKCSOracle(msg)
	decMsg, err := PKCSOracleAttack(ct, pub, n, oracle)
	assert.Nil(t, err, "no error should've happened")
	assert.Equal(t, decMsg, msg, "message should match")
}
