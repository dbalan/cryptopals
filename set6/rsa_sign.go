package set6

import (
	"fmt"
	"github.com/dbalan/cryptopals/sha"
	"math/big"
)

func forgeSig(data []byte, N *big.Int) {
	hs := sha.SHA(data)
	shaPfx := []int{0x30, 0x21, 0x30, 0x09, 0x06, 0x05, 0x2b, 0x0e, 0x03, 0x02,
		0x1a, 0x05, 0x00, 0x04, 0x14}

	assemble := []byte{byte(0xff), byte(0x00)}
	for _, b := range shaPfx {
		assemble = append(assemble, byte(b))
	}

	assemble = append(assemble, hs...)

	// (00 + ff + 00 + assemble + garbage)** 1/3
	fmt.Printf("%b\n", assemble)
	destPfx := new(big.Int).SetBytes(assemble)
	src := oddCubeRootPrefix(destPfx)

	println(src.Text(16))
}

func cubed(src *big.Int) string {
	s := new(big.Int).Set(src)
	s.Exp(s, big.NewInt(3), nil)
	return s.Text(2)
}

func cmpBit(dest, src string, i int) bool {
	var s byte
	// src is less than that bit, so its zero
	if (len(src) - i - 1) < 0 {
		s = byte('0')
	} else {
		s = byte(src[len(src)-i-1])
	}
	return byte(dest[len(dest)-1-i]) == s
}

func flip(src string, i int) string {
	if len(src)-1-i < 0 {
		ns := ""
		padLen := i - (len(src) - 1)
		for p := 0; p < padLen; p++ {
			ns += "0"
		}

		ns += src
		src = ns
	}

	var nb byte
	if []byte(src)[len(src)-1-i] == byte('1') {
		nb = byte('0')
	} else {
		nb = byte('1')
	}

	s := []byte(src)
	s[len(src)-1-i] = nb
	return string(s)
}

func flipbigInt(s *big.Int, i int) {
	dst := flip(s.Text(2), i)
	s.SetString(dst, 2)
}

func oddCubeRootPrefix(destPrefx *big.Int) *big.Int {
	pfx := destPrefx.Text(2)
	src := big.NewInt(1)

	for i := 0; i < len(pfx); i++ {
		// cmp ith bit from last, flip in src
		csrc := cubed(src)

		if !cmpBit(pfx, csrc, i) {
			//println(pfx)
			// println(csrc)
			//println("will flip bit ", i)
			//flip i t bit in src
			flipbigInt(src, i)
		}
	}

	return src
}
