package set6

import (
	"testing"
)

func TestOracle(t *testing.T) {
	ct, oracle := ParityOracle()
	oracle(ct)
}
