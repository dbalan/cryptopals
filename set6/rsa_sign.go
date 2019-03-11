package set6

import (
	"github.com/dbalan/cryptopals/sha"
	"math/big"
)

var (
	shaPfx = []int{0x30, 0x21, 0x30, 0x09, 0x06, 0x05, 0x2b, 0x0e, 0x03, 0x02,
		0x1a, 0x05, 0x00, 0x04, 0x14}
)

func forgeSig(data []byte, N *big.Int) {
	hs := sha.SHA(data)

	assemble := []byte{byte(0x00), byte(0x01), byte(0xff), byte(0x00)}
	for _, b := range shaPfx {
		assemble = append(assemble, byte(b))
	}
	actualsig := make([]byte, len(assemble))
	copy(actualsig, assemble)

	assemble = append(assemble, hs...)

	garbage := []byte{byte(0x00)}

	// pad with garbage to module size.
	for i := 0; i < 1024-(len(assemble)*8); i++ {
		garbage = append(garbage, byte(0xff))
	}
	assemble = append(assemble, garbage...)
	//	println("len: ", len(assemble)*8)

	// (0x00 + 0x01 0xff + 0x00 + assemble + garbage)** 1/3
	destPfx := new(big.Int).SetBytes(assemble)
	croot := findClosestCube(destPfx)

	// signature verification
	// FIXME: move out
	croot.Exp(croot, big.NewInt(3), nil)

	sig := croot.Text(16)
	asig := new(big.Int).SetBytes(actualsig)
	expected := asig.Text(16)
	if sig[0:len(expected)] == expected {
		println("Signature verification passed.")
	}

}

func findClosestCube(a *big.Int) *big.Int {
	low := big.NewInt(0)
	hi := new(big.Int).Set(a)

	for low.Cmp(hi) == -1 {
		mid := new(big.Int).Add(low, hi)
		mid.Div(mid, big.NewInt(2))

		m3 := new(big.Int).Exp(mid, big.NewInt(3), nil)

		if m3.Cmp(a) == -1 {
			low = mid.Add(mid, big.NewInt(1))
		} else {
			hi.Set(mid)
		}
	}
	return low
}

/*
def true_cbrt(n):
    lo = 0
    hi = n
    while lo < hi:
        mid = (lo+hi)//2
        if mid**3 < n:
            lo = mid+1
        else:
            hi = mid
    return lo
*/
