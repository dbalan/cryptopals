package set6

import (
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

func TestSignatureForging(t *testing.T) {
	forgeSig([]byte("hi mom"), big.NewInt(0))
	// 	oddCubeRootPrefix(big.NewInt(2333))
}

func TestCmpBit(t *testing.T) {
	cases := []struct {
		Dst string
		Src string
		Pos int
		Out bool
	}{
		{"1001", "1", 0, true},
		{"1001", "1", 1, true},
		{"1001", "11", 1, false},
		{"1111101", "101", 2, true},
		{"1111101", "101", 3, false},
	}

	for index, cs := range cases {
		out := cmpBit(cs.Dst, cs.Src, cs.Pos)
		assert.Equal(t, cs.Out, out, index)
	}
}

func TestFlipBit(t *testing.T) {
	cases := []struct {
		Src string
		Pos int
		Out string
	}{
		{"1", 0, "0"},
		{"1", 1, "11"},
		{"010", 1, "000"},
		{"1", 10, "10000000001"},
	}

	for index, cs := range cases {
		out := flip(cs.Src, cs.Pos)
		assert.Equal(t, cs.Out, out, index)
	}

}

func TestCubed(t *testing.T) {

	s := new(big.Int)
	s.SetString("101", 2)
	c := cubed(s)
	assert.Equal(t, c, "1111101")
}
