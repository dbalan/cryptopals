package sha

import (
	"github.com/dbalan/cryptopals/common"
)

const (
	h0 = 0x67452301
	h1 = 0xEFCDAB89
	h2 = 0x98BADCFE
	h3 = 0x10325476
	h4 = 0xC3D2E1F0
)

func BEEncodeUint64(x uint64) []byte {
	b := make([]byte, 8)
	for i := 7; i >= 0; i-- {
		b[i] = byte(x & 0xff)
		x = x >> 8
	}

	return b
}

func BEEncodeUint32(x uint32) []byte {
	val := make([]byte, 4)
	for i := 3; i >= 0; i-- {
		val[i] = byte(x & 0xff)
		x >>= 8
	}
	return val
}

func rotateL(x uint32, n uint32) uint32 {
	return (x << n) | (x >> (32 - n))
}

func preprocess(msg []byte) []byte {
	msgLen := len(msg) * 8

	// append 0x80
	msg = append(msg, byte(0x80))
	left := 512 - (msgLen+8)%512
	var padLen int = 0

	if left < 64 {
		padLen = (left + 512 - 64)
	} else if left > 64 {
		padLen = left - 64
	}

	msg = append(msg, common.Repeat(padLen/8, byte(0))...)

	encodedLen := BEEncodeUint64(uint64(msgLen))
	msg = append(msg, encodedLen...)
	return msg
}

func packUint32(w ...byte) uint32 {
	l := len(w)

	var acc uint32
	for i := 0; i < l; i++ {
		acc = acc | uint32(w[l-i-1])<<uint(i*8)
	}
	return acc
}

func chunkEncode(ch []byte, h0, h1, h2, h3, h4 uint32) (p, q, r, s, t uint32) {
	words := make([]uint32, 80)
	groups := common.Blocks(ch, 4)
	for wi, w := range groups {
		words[wi] = packUint32(w...)
	}

	// expand
	for i := 16; i < 80; i++ {
		words[i] = rotateL(words[i-3]^words[i-8]^words[i-14]^words[i-16], 1)
	}

	a := h0
	b := h1
	c := h2
	d := h3
	e := h4

	for i := 0; i < 80; i++ {
		var f, k uint32
		switch {
		case 0 <= i && i < 20:
			f = (b & c) | ((^b) & d)
			k = 0x5A827999
		case 20 <= i && i < 40:
			f = b ^ c ^ d
			k = 0x6ED9EBA1
		case 40 <= i && i < 60:
			f = (b & c) | (b & d) | (c & d)
			k = 0x8F1BBCDC
		case 60 <= i && i < 80:
			f = b ^ c ^ d
			k = 0xCA62C1D6
		default:
			panic("NEVER SHOULD REACH")
		}

		var temp uint32 = uint32(rotateL(a, 5) + f + e + k + words[i])
		e = d
		d = c
		c = rotateL(b, 30)
		b = a
		a = temp

	}

	p = h0 + a
	q = h1 + b
	r = h2 + c
	s = h3 + d
	t = h4 + e
	return
}

func SHA(msg []byte) []byte {
	// msg is 8 bits
	prep := preprocess(msg)
	chs := [][]byte{}
	for i := 0; i < len(prep); i += 64 {
		chs = append(chs, prep[i:i+64])
	}

	var th0 uint32 = h0
	var th1 uint32 = h1
	var th2 uint32 = h2
	var th3 uint32 = h3
	var th4 uint32 = h4

	for _, ch := range chs {
		th0, th1, th2, th3, th4 = chunkEncode(ch, th0, th1, th2, th3, th4)
	}

	resp := []byte{}
	resp = append(resp, BEEncodeUint32(th0)...)
	resp = append(resp, BEEncodeUint32(th1)...)
	resp = append(resp, BEEncodeUint32(th2)...)
	resp = append(resp, BEEncodeUint32(th3)...)
	resp = append(resp, BEEncodeUint32(th4)...)

	return resp
}
