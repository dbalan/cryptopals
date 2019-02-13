package set5

import (
	"fmt"
	"math/big"
	"testing"
)

func TestTestNormalComm(t *testing.T) {
	normalRecv := newB()
	communicate(normalRecv)
}

func TestKey0MITM(t *testing.T) {
	mal := newK0Mitm()
	communicate(mal)
}

func TestPrimeGMask1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recoverd from botched second leg g = 1")
		}
	}()
	mal := newGMaskMitm(big.NewInt(1))
	communicate(mal)
}

func TestPrimeGMaskP(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recoverd from botched second leg g = p")
		}
	}()
	p, _ := primes()
	mal := newGMaskMitm(p)
	communicate(mal)
}

func TestPrimeGMaskP_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recoverd from botched second leg g = p-1")
		}
	}()
	p, _ := primes()
	mal := newGMaskMitm(p.Sub(p, big.NewInt(1)))
	communicate(mal)
}
