package set1

import (
	"bytes"
	"github.com/dbalan/cryptopals/common"
	"sort"
)

type PTKeyPair struct {
	PT    []byte
	Key   byte
	Score float64
}

type byScore []PTKeyPair

func (s byScore) Len() int {
	return len(s)
}

func (s byScore) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byScore) Less(i, j int) bool {
	return s[i].Score < s[j].Score
}

func decryptSingleXOR(ct []byte, key byte) []byte {
	resp := []byte{}
	for _, c := range ct {
		resp = append(resp, c^key)
	}
	return resp
}

func allPossibleDecryptions(ct []byte) []PTKeyPair {
	resp := []PTKeyPair{}

	for i := 0; i <= 0xff; i++ {
		pt := decryptSingleXOR(ct, byte(i))
		resp = append(resp, PTKeyPair{pt, byte(i), scoreByWord(pt)})
	}
	return resp
}

func scoreByWord(s []byte) float64 {
	words := bytes.Split(s, []byte(" "))
	acc := 0.0
	for _, w := range words {
		acc += common.TextScore(w)
	}
	return acc
}

func possibleKeysWithPT(ct []byte, num int) (resp []PTKeyPair) {
	ptkp := allPossibleDecryptions(ct)
	sort.Sort(sort.Reverse(byScore(ptkp)))
	return ptkp[0:num]
}

func BestPT(ct []byte) ([]byte, byte) {
	curHigh := 0.0
	result := PTKeyPair{}

	for _, ptkp := range allPossibleDecryptions(ct) {
		if ptkp.Score > curHigh {
			curHigh = ptkp.Score
			result = ptkp
		}
	}
	return result.PT, result.Key
}
