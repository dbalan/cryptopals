package rsa

import (
	"crypto/rand"
	"fmt"
	"github.com/dbalan/cryptopals/common"
	"github.com/dbalan/cryptopals/sha"
	"math/big"
)

func naivePrimalityCheck(p *big.Int) bool {
	zero := big.NewInt(0)

	if p.Cmp(big.NewInt(1)) == -1 {
		return false
	}

	if p.Cmp(big.NewInt(3)) == -1 {
		return true
	}

	if new(big.Int).Mod(p, big.NewInt(2)).Int64() == 0 {
		return false
	}

	if new(big.Int).Mod(p, big.NewInt(3)).Int64() == 0 {
		return false
	}

	i := big.NewInt(5)

	for {
		sq := new(big.Int).Mul(i, i)
		if sq.Cmp(p) == 1 {
			break
		}

		if new(big.Int).Mod(p, i).Cmp(zero) == 0 {
			return false
		}

		i2 := new(big.Int).Add(i, big.NewInt(2))
		if i2.Mod(p, i2).Cmp(zero) == 0 {
			return false
		}

		i.Add(i, big.NewInt(6))
	}
	return true
}

func prime(size int) (*big.Int, error) {
	//again:
	p, err := rand.Prime(rand.Reader, size)
	if err != nil {
		return nil, err
	}

	//  FIXME: read the guarantees of rand.Prime
	//  Replace with openssl prime functions?
	//	if !naivePrimalityCheck(p) {
	//		goto again
	//	}

	return p, nil
}

// compute fermat's prime Fn
func fermatPrime(n int64) *big.Int {
	f := new(big.Int)
	f.Exp(big.NewInt(2), big.NewInt(n), nil)
	f.Exp(big.NewInt(2), f, nil)
	f.Add(f, big.NewInt(1))
	return f
}

func genprimes(size int) (p, q *big.Int, err error) {
	p, err = prime(size)
	if err != nil {
		return
	}

get_another:
	q, err = prime(size)
	if err != nil {
		return
	}

	if p.Cmp(q) == 0 {
		goto get_another
	}
	return
}

func GenKeyPair(size int) (e, d, n *big.Int, err error) {
	p, q, err := genprimes(size)
	if err != nil {
		return
	}

	n = new(big.Int).Mul(p, q)

	p_1 := new(big.Int).Sub(p, big.NewInt(1))
	q_1 := new(big.Int).Sub(q, big.NewInt(1))
	totient := new(big.Int).Mul(p_1, q_1)

	// e should be coprime to totient, or not choose another
	// Fermat prime F0
	e = fermatPrime(0)
	// fixme: pick a random e in (1, totient)
	var i int64
	for i = 1; common.EGCD(totient, e).Cmp(big.NewInt(1)) != 0; i++ {
		e = fermatPrime(i)
	}

	d, err = common.InvMod(e, totient)
	return
}

func op(data, k, n *big.Int) *big.Int {
	return new(big.Int).Exp(data, k, n)
}

func Encrypt(msg []byte, pub, N *big.Int) *big.Int {
	// naive serialization
	hx := fmt.Sprintf("%x", msg)
	m := new(big.Int)
	m.SetString(hx, 16)

	return op(m, pub, N)
}

func Decrypt(ct, priv, N *big.Int) ([]byte, error) {
	pt := op(ct, priv, N)
	var msg []byte
	_, err := fmt.Sscanf(pt.Text(16), "%x", &msg)
	if err != nil {
		return nil, err
	}
	return msg, nil
}

func Sign(msg []byte, priv, N *big.Int) string {
	// msg to int

	hmsg := common.EncodeHexString(sha.SHA(msg))
	hnum, _ := new(big.Int).SetString(hmsg, 16)

	s := op(hnum, priv, N)
	return s.Text(16)
}

func VerifySign(msg []byte, sign string, pub, N *big.Int) bool {

	hsign, _ := new(big.Int).SetString(sign, 16)

	s := op(hsign, pub, N)
	// s == hash(msg)
	hmsg := common.EncodeHexString(sha.SHA(msg))
	actual, _ := new(big.Int).SetString(hmsg, 16)

	return (actual.Cmp(s) == 0)
}
