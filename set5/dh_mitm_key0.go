package set5

import (
	"fmt"
	"github.com/dbalan/cryptopals/common"
	"math/big"
)

type K0Mitm struct {
	B                *PersonB
	p, g, pubA, pubB *big.Int
}

func newK0Mitm() *K0Mitm {
	B := newB()
	return &K0Mitm{B: B}
}

func (m *K0Mitm) PrimeEx(p, g *big.Int) {
	m.B.PrimeEx(p, g)
	m.p = p
	m.g = g
}

func (m *K0Mitm) Kex(pub *big.Int) *big.Int {
	m.pubA = pub
	// send B bogus
	pubB := m.B.Kex(m.p)
	m.pubB = pubB

	// send A bogus
	return m.p
}

func (m *K0Mitm) Exchange(ct, iv []byte) ([]byte, []byte) {
	mct := common.CopyBlock(ct)
	miv := common.CopyBlock(iv)
	// (p**priv) mod p = 0!
	msg := decryptWithSK(big.NewInt(0), mct, miv)
	fmt.Println("MITM: ", string(msg))
	nct, niv := m.B.Exchange(ct, iv)
	return nct, niv
}
