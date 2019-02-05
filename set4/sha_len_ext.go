package set4

import (
	"github.com/dbalan/cryptopals/common"
	"github.com/dbalan/cryptopals/sha"
)

// checkAdmin verifies if msg passes the mac test, and then checks if admin=true
func checkAdmin(key, msg, mac []byte) bool {
	actualmac := SHA1MAC(key, msg)
	if !common.EqualBlocks(actualmac, mac) {
		return false
	}

	return isAdmin(string(msg))
}

func gluePad(msg []byte, l uint64) []byte {
	m := make([]byte, len(msg))
	copy(m, msg)

	m = append(m, byte(0x80))
	minLen := len(msg) * 8

	// say keylen is len
	left := 512 - (minLen+int(l*8)+8)%512
	var padLen int = 0

	if left < 64 {
		padLen = left + 512 - 64
	} else if left > 64 {
		padLen = left - 64
	}

	m = append(m, common.Repeat(padLen/8, byte(0))...)
	m = append(m, sha.BEEncodeUint64(uint64(minLen+int(l*8)))...)
	return m
}

func appendMsg(msg, extend []byte, keylen uint64) ([]byte, []byte) {
	// forged = msg || glue || extend
	// msg || 0x80 || 0....0 || len || extend

	forged := append(gluePad(msg, keylen), extend...)
	// pad this new one
	return forged, gluePad(forged, keylen)
}

func decompose(mac []byte) (h0, h1, h2, h3, h4 uint32) {
	// decompose mac
	var h []uint32

	if len(mac) != 5*4 {
		panic("UNKNOWN LEN")
	}

	for i := 0; i < len(mac); i += 4 {
		hv := common.PackUint32(mac[i : i+4]...)
		h = append(h, hv)
	}

	return h[0], h[1], h[2], h[3], h[4]
}

func shaLenExtAttack(mac, msg []byte, check func([]byte, []byte) bool) (newmac,
	newmsg []byte, keylen int) {

	a, b, c, d, e := decompose(mac)
	extend := []byte(";admin=true;")

	// bruteforce on a keylen where you can find a new hmac
	for keylen = 0; keylen < 100; keylen++ {
		var paddedNewMsg []byte
		newmsg, paddedNewMsg = appendMsg(msg, extend, uint64(keylen))
		// fmt.Printf("newmsg: % x\n", newmsg)
		//fmt.Printf("pnewms: % x\n", paddedNewMsg)

		toencode := paddedNewMsg[64-uint64(keylen):]
		//	fmt.Printf("chunk2: % x\n", toencode)

		newmac = sha.PartialSHA(toencode, a, b, c, d, e)
		//	fmt.Printf("newmsg: ")
		if check(newmac, newmsg) {
			return
		}
	}
	panic("NO DICE")
}
