package set2

import (
	"strings"
)

func cbcOracle(in string, key []byte) []byte {
	in = strings.Replace(in, ";", "%3B", -1)
	in = strings.Replace(in, "=", "%3D", -1)
	pt := []byte("comment1=cooking%20MCs;userdata=" + in + ";comment2=%20like%20a%20pound%20of%20bacon")
	lpt := len(pt)
	if (lpt % 16) != 0 {
		newlen := lpt + 16 - (lpt % 16)
		padded, err := PKCS7Padding(pt, newlen)
		if err != nil {
			panic(err)
		}
		pt = padded
	}
	iv := []byte{}
	for i := 0; i < 16; i++ {
		iv = append(iv, byte(0))
	}

	enc, err := EncAES128CBC(pt, iv, key)
	if err != nil {
		panic(err)
	}

	return enc
}

func decryptOracle(ct []byte, key []byte) bool {
	iv := []byte{}
	for i := 0; i < 16; i++ {
		iv = append(iv, byte(0))
	}

	dec, err := DecAES128CBC(ct, iv, key)
	if err != nil {
		panic(err)
	}

	return isAdmin(string(dec))
}

func isAdmin(in string) bool {
	pairs := strings.Split(in, ";")
	if len(pairs) < 1 {
		return false
	}
	for _, pr := range pairs {
		spr := strings.Split(pr, "=")
		if len(spr) != 2 {
			// ignore
			continue
		}

		k := spr[0]
		v := spr[1]

		if k == "admin" && v == "true" {
			return true
		}
	}

	return false
}
