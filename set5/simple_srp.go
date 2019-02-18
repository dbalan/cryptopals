package set5

import (
	"math/big"
	"math/rand"
)

type SSRPSvr interface {
	ExchangePub(*big.Int) (uint64, *big.Int, *big.Int)
	CheckAuth(string) bool
}

type goodSSRP struct {
	u, v, A, b, B *big.Int

	salt uint64
	N, g *big.Int
}

func newSimpleSRPSvr(pass string) *goodSSRP {
	N, g := primes()

	salt := rand.Uint64()
	x := saltHmac(salt, pass)
	v := &big.Int{}
	v.Exp(g, x, N)
	return &goodSSRP{v: v, salt: salt, N: N, g: g}
}

func (s *goodSSRP) ExchangePub(A *big.Int) (uint64, *big.Int, *big.Int) {
	s.A = A
	b := randUint()

	B := &big.Int{}
	B.Exp(s.g, b, s.N)

	uX := randUint()

	u := SHA256Int(uX)

	s.u = u
	s.B = B
	s.b = b
	return s.salt, B, u
}

func (s *goodSSRP) CheckAuth(cauth string) bool {
	//S = (A * V ** u) ** b % N
	S := &big.Int{}
	S.Exp(s.v, s.u, s.N)
	S.Mul(S, s.A)
	S.Exp(S, s.b, s.N)
	K := SHA256Int(S)

	if HMAC_SHA256(K.Text(16), s.salt) == cauth {
		return true
	}
	return false
}

func smplSRPLogin(svr SSRPSvr, pass string) bool {
	N, g := primes()
	a := randUint()
	A := &big.Int{}
	A.Exp(g, a, N)

	salt, B, u := svr.ExchangePub(A)

	x := saltHmac(salt, pass)
	S := &big.Int{}
	P1 := &big.Int{}
	P1.Mul(u, x)
	P1.Add(P1, a)
	S.Exp(B, P1, N)

	K := SHA256Int(S)

	auth := HMAC_SHA256(K.Text(16), salt)
	return svr.CheckAuth(auth)
}
