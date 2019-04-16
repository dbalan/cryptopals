package set6

import (
	"fmt"
	"github.com/dbalan/cryptopals/rsa"
	"math/big"
)

type Server struct {
	pub, priv, n *big.Int
	ctcount      map[*big.Int]int
}

func NewSvr() (*Server, error) {
	pub, priv, n, err := rsa.GenKeyPair(128)
	if err != nil {
		return nil, err
	}
	ctcount := make(map[*big.Int]int)
	return &Server{pub: pub, priv: priv, n: n, ctcount: ctcount}, nil
}

func (s *Server) Encrypt(msg []byte) (ct, pub, n *big.Int) {
	ct = rsa.Encrypt(msg, s.pub, s.n)
	return ct, s.pub, s.n
}

func (s *Server) Decrypt(ct *big.Int) ([]byte, error) {
	if val, ok := s.ctcount[ct]; !ok {
		s.ctcount[ct] = 1
	} else if val >= 1 {
		return nil, fmt.Errorf("Seen this before, go away!")
	}
	return rsa.Decrypt(ct, s.priv, s.n)
}

// EVIL functions start here.

func tamperRsa(ct, pub, n *big.Int) (*big.Int, *big.Int) {
	s := big.NewInt(23445)

	cprime := new(big.Int).Set(pub)
	cprime.Exp(s, cprime, n)
	cprime.Mul(cprime, ct)
	cprime.Mod(cprime, n)
	return cprime, s
}

func recoverRsa(msg []byte, multi, n *big.Int) ([]byte, error) {
	mx := fmt.Sprintf("%x", msg)
	m, _ := new(big.Int).SetString(mx, 16)
	m.DivMod(m, multi, n)

	var recov []byte
	_, err := fmt.Sscanf(m.Text(16), "%x", &recov)
	return recov, err
}
