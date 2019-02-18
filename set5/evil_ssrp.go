package set5

import (
	"math/big"
)

type evilSSRP struct {
	salt     uint64
	N, g     *big.Int
	A        *big.Int
	password string
}

func newEvilSSRP() *evilSSRP {
	N, g := primes()
	var salt uint64 = 0

	return &evilSSRP{salt: salt, N: N, g: g}
}

func (s *evilSSRP) ExchangePub(A *big.Int) (uint64, *big.Int, *big.Int) {
	s.A = A
	b := big.NewInt(1)
	B := &big.Int{}
	B.Exp(g, b, N) // == g

	return s.salt, B, big.NewInt(1)
}

func (s *evilSSRP) CheckAuth(cauth string) bool {
	for _, p := range passwords() {
		// S = A * (v ** 1)**1 % N
		// S = A * v
		// v = g**x
		x := saltHmac(0, p)
		v := &big.Int{}
		v.Exp(s.g, x, s.N)

		S := &big.Int{}
		S.Mod(s.A, s.N)
		S.Mul(S, v)
		S.Mod(S, s.N)
		K := SHA256Int(S)
		if HMAC_SHA256(K.Text(16), s.salt) == cauth {
			s.password = p
			return false
		}
	}
	return false
}

func (s *evilSSRP) GetPassword() string {
	return s.password
}
