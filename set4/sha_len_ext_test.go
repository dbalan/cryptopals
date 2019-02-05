package set4

import (
	"fmt"
	"github.com/dbalan/cryptopals/common"
	"github.com/dbalan/cryptopals/sha"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAppendMessage(t *testing.T) {
	msg := []byte("hello world")
	key := []byte("HELLO") // len 5
	extend := []byte("cruel world")

	keyed := append(key, msg...)
	padded := sha.Preprocess(keyed)
	newmsg := append(padded, extend...)
	expected := sha.Preprocess(newmsg)[len(key):]
	//	fmt.Printf("expected     : % x\n", expected)

	_, forgedPadded := appendMsg(msg, extend, uint64(len(key)))
	//	fmt.Printf("actual append: % x\n", forged)

	assert.Equal(t, expected, forgedPadded)
}

func TestDecomposeSHA1(t *testing.T) {
	mac := common.DecodeB64([]byte("L9ThxnotKPzthJ7hu3bnORuT6xI="))

	var h0 uint32 = 802480582
	var h1 uint32 = 2049779964
	var h2 uint32 = 3984891617
	var h3 uint32 = 3145131833
	var h4 uint32 = 462678802

	a, b, c, d, e := decompose(mac)

	assert.Equal(t, a, h0)
	assert.Equal(t, b, h1)
	assert.Equal(t, c, h2)
	assert.Equal(t, d, h3)
	assert.Equal(t, e, h4)
}

func TestSHALenAttack(t *testing.T) {
	msg := []byte(";what=when;admin=false;")
	key := []byte("hello")
	mac := SHA1MAC(key, msg)

	check := func(mac, msg []byte) bool {
		return checkAdmin(key, msg, mac)
	}

	frmac, frmsg, keylen := shaLenExtAttack(mac, msg, check)
	fmt.Printf("found keylen: %d forged msg: % x \nforged msg: % x\n", keylen, frmsg, frmac)

	assert.Equal(t, len(key), keylen)
}
