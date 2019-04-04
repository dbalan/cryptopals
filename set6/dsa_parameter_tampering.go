package set6

import (
	"fmt"
	"github.com/dbalan/cryptopals/dsa"
	"github.com/dbalan/cryptopals/sha"
	"math/big"
)

/*  won't work!!, s would be 0, and we would go on a loop
    s == 0
func getG0() (p, q, g *big.Int) {
	p, q, _ = dsa.GetDefaultParams()
	g = big.NewInt(0)
	return
}

func tamper0() error {
	x, y, err := dsa.KeyPair(getG0)
	if err != nil {
		return err
	}

	println("y", y.Text(16))
	r, s, err := dsa.Sign([]byte("Hey!"), x, getG0)
	if err != nil {
		return err
	}

	println(r.Text(16), s.Text(16))
	return nil
}
*/

func magicSig(msg []byte) (r, s *big.Int) {
	_, q, _ := getGP1()
	hs := new(big.Int).SetBytes(sha.SHA(msg))
	// r := (y ** s) % p % q
	r = big.NewInt(1)
	s = new(big.Int).Exp(hs, big.NewInt(-1), q)
	return
}

func getGP1() (p, q, g *big.Int) {
	p, q, _ = dsa.GetDefaultParams()
	g = new(big.Int).Add(big.NewInt(1), p)
	return
}

func tamperP1(m []byte) error {
	_, y, err := dsa.KeyPair(getGP1)
	if err != nil {
		return err
	}
	r, s := magicSig(m)
	sig := dsa.Verify(m, r, s, y, getGP1)

	if !sig {
		return fmt.Errorf("signature failed")
	}
	return nil
}
