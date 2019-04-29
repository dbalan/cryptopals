package set6

import (
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

func TestOracle(t *testing.T) {
	ct, _, _, oracle := ParityOracle()
	oracle(ct)
}

func TestTimes(t *testing.T) {
	res := times(big.NewInt(16))
	assert.Equal(t, 4, res)
}

func TestBreakOracle(t *testing.T) {
	ct, e, n, oracle := ParityOracle()
	parityOracleAttack(e, n, ct, oracle)
}
