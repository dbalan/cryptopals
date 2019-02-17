// this is the SRP server implimentation, set5, ch 36
// see http://srp.stanford.edu/design.html
package set5

import (
	"math/big"
	"math/rand"
)

const username = "user@email.com"

var (
	g = big.NewInt(2)
	k = big.NewInt(3)
	N = big.NewInt(0)
)

type Server struct {
	v, b *big.Int
	salt uint64
	I    string
	A, B *big.Int
}

func NewServer() *Server {
	svr := &Server{I: username}
	password := "5upers4cr4t"
	salt := rand.Uint64()
	x := saltHmac(salt, password)

	svr.salt = salt
	// generate v and store it
	v := &big.Int{}
	v.Exp(g, x, N)
	svr.v = v
	return svr
}

func (s *Server) ExchangePub(I string, A *big.Int) (salt uint64, B *big.Int) {
	if I != s.I {
		panic("username doesn't match")
	}

	// generate session key and send pubkey
	s.b = &big.Int{}
	s.b.SetUint64(rand.Uint64())
	s.A = A
	B = &big.Int{}

	B1 := &big.Int{}
	B1.Mul(k, s.v)
	B.Exp(g, s.b, N)
	B.Add(B, B1)
	s.B = B

	// compute u and store it
	return s.salt, B
}

func (s *Server) CheckAuth(cauth string) bool {
	//	S = (A * v**u) ** b % N
	u := SHA256Int(s.A, s.B)
	S := &big.Int{}
	S.Exp(s.v, u, N)
	S.Mul(S, s.A)
	S.Exp(S, s.b, N)

	key := SHA256Int(S)
	actual := HMAC_SHA256(key.Text(16), s.salt)

	if cauth == actual {
		return true
	}
	return false
}

func login(password string) bool {
	// agree on N
	N, _ = primes()
	// connect to server
	server := NewServer()

	// get a random key
	a := &big.Int{}
	a.SetUint64(rand.Uint64())

	// generate pubKey
	A := &big.Int{}
	A.Exp(g, a, N)

	// exchange keys with server
	salt, B := server.ExchangePub(username, A)

	// shared secret (ish)
	u := SHA256Int(A, B)

	// encode password
	pwhs := saltHmac(salt, password)

	// Session
	// S = (B - k * g**x)**(a + u * x) % N
	P2 := big.NewInt(0)
	P2.Mul(u, pwhs)
	P2.Add(P2, a)
	// S = (B - k * g**x)**P2 % N
	S := &big.Int{}
	S.Exp(g, pwhs, N)
	S.Mul(S, k)
	S.Sub(B, S)
	S.Exp(S, P2, N)

	key := SHA256Int(S)
	cauth := HMAC_SHA256(key.Text(16), salt)

	return server.CheckAuth(cauth)
}

func loginWithZero() bool {
	// agree on N
	N, _ = primes()
	// connect to server
	server := NewServer()

	A := big.NewInt(0)

	// exchange keys with server
	salt, _ := server.ExchangePub(username, A)
	S := big.NewInt(0)

	key := SHA256Int(S)
	cauth := HMAC_SHA256(key.Text(16), salt)

	return server.CheckAuth(cauth)
}

func loginWithN() bool {
	N, _ = primes()

	server := NewServer()
	// try multiples of N
	mul := big.NewInt(rand.Int63n(10))
	mul.Abs(mul)
	mul.Exp(N, mul, nil)
	salt, _ := server.ExchangePub(username, mul)
	S := big.NewInt(0)
	key := SHA256Int(S)
	cauth := HMAC_SHA256(key.Text(16), salt)
	return server.CheckAuth(cauth)
}
