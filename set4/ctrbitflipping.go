package set4

import (
	"github.com/dbalan/cryptopals/common"
	"github.com/dbalan/cryptopals/set3"
	"math/rand"
	"strings"
)

func ctrOracle(in string, key []byte, nonce uint64) ([]byte, error) {
	in = strings.Replace(in, ";", "%3B", -1)
	in = strings.Replace(in, "=", "%3D", -1)
	pt := []byte("comment1=cooking%20MCs;userdata=" + in + ";comment2=%20like%20a%20pound%20of%20bacon")

	return set3.AES128CTR(pt, key, nonce)
}

func decryptOracle(ct []byte, key []byte, nonce uint64) (bool, error) {
	pt, err := set3.AES128CTR(ct, key, nonce)
	if err != nil {
		return false, err
	}

	return isAdmin(string(pt)), nil
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

func bitFlip(oracle func(in string) []byte, check func(ct []byte) bool) (bool, error) {
	expected := []byte(";admin=true;")

	repeating := common.Repeat(4*len(expected), byte('A'))

	somect := oracle(string(repeating))

	// recover keystream, its not correct = but it will be for the middle
	// part, all we care is that

	keystream, err := common.XORBlk(somect, common.Repeat(len(somect), byte('A')))
	if err != nil {
		return false, err
	}

	repExp := []byte{}
	for i := 0; i < len(somect)+len(expected); i = i + len(expected) {
		repExp = append(repExp, expected...)
	}
	// prune to correct size
	repExp = repExp[0:len(somect)]
	// create a new CT

	newct, err := common.XORBlk(keystream, repExp)
	if err != nil {
		return false, err
	}
	return check(newct), nil
}

func Attack() (bool, error) {
	key, err := common.RandBytes(16)
	if err != nil {
		panic(err)
	}

	nonce := rand.Uint64()
	oracle := func(in string) []byte {
		ct, err := ctrOracle(in, key, nonce)
		if err != nil {
			panic(err)
		}
		return ct
	}

	check := func(ct []byte) bool {
		status, err := decryptOracle(ct, key, nonce)
		if err != nil {
			panic(err)
		}
		return status
	}

	return bitFlip(oracle, check)
}
