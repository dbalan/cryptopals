package common

import (
	"math/big"
)

// Extended eucledean algo
// A = B⋅Q+R
// GCD(A,B) == GCD(B,R)

func EGCD(A, B *big.Int) *big.Int {
	zero := big.NewInt(0)
	if A.Cmp(zero) == 0 {
		return B
	}

	if B.Cmp(zero) == 0 {
		return A
	}

	quote := &big.Int{}
	quote.Rem(A, B)

	return EGCD(B, quote)
}
