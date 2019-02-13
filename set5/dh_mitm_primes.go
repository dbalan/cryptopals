package set5

import (
	"fmt"
	"github.com/dbalan/cryptopals/common"
	"math/big"
)

// MITM attacks
type GMaskMitm struct {
	B                *PersonB
	maskg            *big.Int
	p, g, pubA, pubB *big.Int
}

func newGMaskMitm(maskg *big.Int) *GMaskMitm {
	B := newB()
	return &GMaskMitm{maskg: maskg, B: B}
}

func (m *GMaskMitm) PrimeEx(p, g *big.Int) {
	m.p = p
	m.g = g
	m.B.PrimeEx(p, m.maskg)
}

func (m *GMaskMitm) Kex(pub *big.Int) *big.Int {
	m.pubA = pub
	pubB := m.B.Kex(m.pubA)
	m.pubB = pubB
	return m.pubB
}

func (m *GMaskMitm) Exchange(ct, iv []byte) ([]byte, []byte) {
	if m.maskg.Cmp(big.NewInt(1)) == 0 {
		skey := big.NewInt(1)
		mct := common.CopyBlock(ct)
		miv := common.CopyBlock(iv)
		// (p**priv) mod p = 0!
		msg := decryptWithSK(skey, mct, miv)
		fmt.Println("MITM: g == 1", string(msg))
	}

	if m.maskg.Cmp(m.p) == 0 {
		// p**b mod p = 0 == B
		// B**a mod p = 0 == skeyA
		// A**b mod p != 0 skeyB
		skey := big.NewInt(0)
		mct := common.CopyBlock(ct)
		miv := common.CopyBlock(iv)
		// (p**priv) mod p = 0!
		msg := decryptWithSK(skey, mct, miv)
		fmt.Println("MITM: g == p", string(msg))
	}

	if m.maskg.Add(m.maskg, big.NewInt(1)).Cmp(m.p) == 0 {
		// (p-1)**b % p = 1 | p-1 either 1 or p-1
		check := func(skey *big.Int) {
			defer func() {
				recover()
			}()
			mct := common.CopyBlock(ct)
			miv := common.CopyBlock(iv)
			// (p**priv) mod p = 0!
			msg := decryptWithSK(skey, mct, miv)
			fmt.Printf("MITM: g == p-1, key = %s: %s", skey.String(),
				string(msg))
		}

		check(big.NewInt(1))
		check(m.p.Sub(m.p, big.NewInt(1)))
	}

	return m.B.Exchange(ct, iv)
}
