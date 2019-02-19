package common

import (
	"errors"
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

func InvMod(a, n *big.Int) (*big.Int, error) {
	zero := big.NewInt(0)

	t := big.NewInt(0)
	nt := big.NewInt(1)

	r := new(big.Int).Set(n)
	nr := new(big.Int).Set(a)

	for nr.Cmp(zero) != 0 {
		quot := &big.Int{}
		quot.Div(r, nr)

		// PSWAP:
		// t = nt
		// nt = t - quot*nt
		tmp := new(big.Int).Set(nt)
		nt.Mul(nt, quot)
		nt.Sub(t, nt)
		t = tmp

		tmp = new(big.Int).Set(nr)
		nr.Mul(nr, quot)
		nr.Sub(r, nr)
		r = tmp
	}

	if r.Cmp(big.NewInt(1)) == 1 {
		return nil, errors.New("Not invertible")
	}

	if t.Cmp(zero) == -1 {
		t.Add(t, n)
	}
	return t, nil
}
