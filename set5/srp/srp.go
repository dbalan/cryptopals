// this is the SRP server implimentation, set5, ch 36
package main

import (
	"math/big"
	"math/rand"

	"fmt"
	"github.com/dbalan/cryptopals/set5"
)

const username = "user@email.com"

var (
	g = big.NewInt(2)
	k = big.NewInt(3)
	N = big.NewInt(0)
)

type Server struct {
	u, v *big.Int
	salt uint64
	I    string
	A, B *big.Int
	b    uint64
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

	b := rand.Uint64()
	s.b = b
	s.A = A
	B = &big.Int{}

	B1 := &big.Int{}
	B1.Mul(k, s.v)
	B.Exp(g, big.NewInt(int64(b)), N)
	B.Add(B, B1)
	s.B = B

	// compute u and store it
	u := SHA256Int(A, B)
	s.u = u
	return s.salt, B
}

func (s *Server) CalculateKey() {
	//	S = (A * v**u) ** b % N
	S := &big.Int{}
	S.Exp(s.v, s.u, N)
	S.Mul(S, s.A)
	S.Exp(S, big.NewInt(int64(s.b)), N)
	KEY := SHA256Int(S)

}

type Client struct{}

// FIXME: move to interactive
func communication() {
	// don't really care here.
	I := username

	svr := NewServer()
	// send I, A=g**a % N
	a := rand.Uint64()
	A := &big.Int{}

	A.Exp(g, big.NewInt(int64(a)), N)
	fmt.Println("ExchangePub keys")
	salt, B := svr.ExchangePub(I, A)
	fmt.Println("Computing keys")

	u := SHA256Int(A, B)

	password := "5upers4cr4t"
	x := saltHmac(salt, password)
	S := &big.Int{}
	// S = (B - k * g**x)**(a + u * x) % N
	P2 := big.NewInt(0)
	P2.Mul(u, x)
	P2.Add(P2, big.NewInt(int64(a)))

	S.Exp(g, x, N)
	S.Mul(S, k)
	S.Sub(B, S)
	S.Exp(S, P2, N)

	KEY := SHA256Int(S)
	fmt.Println(KEY.Text(16))
	svr.CalculateKey()
}

func main() {
	p, _ := set5.GetNistPrimes()
	N = p

	communication()
}
