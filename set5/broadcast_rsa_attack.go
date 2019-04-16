package set5

import (
	"fmt"
	"github.com/dbalan/cryptopals/common"
	"github.com/dbalan/cryptopals/rsa"
	"math/big"
)

const keySize = 128

// mimics grabbing ciphertext out of wire
func getCipherTextE3(msg []byte) (ct, n *big.Int, err error) {
	three := big.NewInt(3)
	pub := new(big.Int)
try_again:
	pub, _, n, err = rsa.GenKeyPair(keySize)
	if err != nil {
		return
	}

	if pub.Cmp(three) != 0 {
		// try to get p, q such that (p-1)(q-1) and 3 are coprime
		goto try_again
	}

	ct = rsa.Encrypt(msg, pub, n)
	return
}

func HastadBCAttack() error {
	msg := []byte("HEL")
	cts := []*big.Int{}
	nlist := []*big.Int{}

	for i := 0; i < 3; {
		ct, n, err := getCipherTextE3(msg)
		if err != nil {
			return err
		}
		for _, pn := range nlist {
			if pn.Cmp(n) == 0 {
				// we need atleast 3 different n
				continue
			}
		}

		cts = append(cts, ct)
		nlist = append(nlist, n)
		i++
	}

	// we have three ciphertexts and three corresponding Ns
	// read: https://crypto.stanford.edu/pbc/notes/numbertheory/crt.html
	// CRT to compute C ≡ (M^3) mod (N1*N2*N3)
	// (c_0 * m_s_0 * invmod(m_s_0, n_0)) +
	// (c_1 * m_s_1 * invmod(m_s_1, n_1)) +
	// (c_2 * m_s_2 * invmod(m_s_2, n_2)) mod N_012 == M^3

	MS0 := new(big.Int).Mul(nlist[1], nlist[2])
	MS1 := new(big.Int).Mul(nlist[0], nlist[2])
	MS2 := new(big.Int).Mul(nlist[0], nlist[1])

	IM0, err := common.InvMod(MS0, nlist[0])
	if err != nil {
		return err
	}

	IM1, err := common.InvMod(MS1, nlist[1])
	if err != nil {
		return err
	}

	IM2, err := common.InvMod(MS2, nlist[2])
	if err != nil {
		return err
	}

	M30 := mul3(cts[0], MS0, IM0)
	M31 := mul3(cts[1], MS1, IM1)
	M32 := mul3(cts[2], MS2, IM2)

	N123 := new(big.Int).Mul(nlist[0], nlist[1])
	N123.Mul(N123, nlist[2])

	M3 := new(big.Int).Add(M30, M31)
	M3.Add(M3, M32)

	M3.Mod(M3, N123)

	// M3 == msg^3
	mint := fmt.Sprintf("%x", msg)
	m, _ := new(big.Int).SetString(mint, 16)
	m.Exp(m, big.NewInt(3), nil)

	if m.Cmp(M3) != 0 {
		return fmt.Errorf("extracted doesn't match")
	}
	return nil
}

func mul3(a, b, c *big.Int) *big.Int {
	n := new(big.Int)
	n.Mul(a, b)
	n.Mul(n, c)
	return n
}
