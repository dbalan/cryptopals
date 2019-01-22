package set2

import (
	//	"fmt"
	"github.com/dbalan/cryptopals/common"

	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCBCOracle(t *testing.T) {
	key, err := common.RandBytes(16)

	assert.Nil(t, err)

	ct := cbcOracle("hello", key)
	resp := decryptOracle(ct, key)
	assert.Equal(t, false, resp)
}

func TestIsAdmin(t *testing.T) {
	cases := []struct {
		In  string
		Out bool
	}{
		{"something=what;admin=false;ehat=wer", false},
		{"admin=true", true},
		{"", false},
		{"admin=true;what=when", true},
	}

	for _, cs := range cases {
		assert.Equal(t, cs.Out, isAdmin(cs.In))

	}
}

func TestFliping(t *testing.T) {
	key, err := common.RandBytes(16)
	assert.Nil(t, err)
	// pad data, see if flipping works
	in := "?admin?true"
	data := []byte("comment1=cooking%20MCs;userdata=" + in + ";comment2=%20like%20a%20pound%20of%20bacon")

	if len(data)%16 != 0 {
		padto := len(data) + 16 - (len(data) % 16)
		//		fmt.Println(padto, len(data))
		d, err := PKCS7Padding(data, padto)
		assert.Nil(t, err)
		data = d
	}

	// should not work
	enc := cbcOracle(in, key)
	// cbc function is inplace, so work on a copy
	enccopy := make([]byte, len(enc))
	copy(enccopy, enc)
	// return false
	assert.Equal(t, false, decryptOracle(enccopy, key))

	// flip bits
	// 1. convert ? -> ;
	fst := flippedBit(enc[16], '?', ';')
	enc[16] = fst

	// 2 convert ? -> =
	snd := flippedBit(enc[22], '?', '=')
	enc[22] = snd

	// and here we go
	assert.Equal(t, true, decryptOracle(enc, key))
}
