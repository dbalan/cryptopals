package set1

import (
	"bytes"
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

var freqEn = map[byte]float64{
	'a': 8.167,
	'b': 1.492,
	'c': 2.782,
	'd': 4.253,
	'e': 12.70,
	'f': 2.228,
	'g': 2.015,
	'h': 6.094,
	'i': 6.966,
	'j': 0.153,
	'k': 0.772,
	'l': 4.025,
	'm': 2.406,
	'n': 6.749,
	'o': 7.507,
	'p': 1.929,
	'q': 0.095,
	'r': 5.987,
	's': 6.327,
	't': 9.056,
	'u': 2.758,
	'v': 0.978,
	'w': 2.360,
	'x': 0.150,
	'y': 1.974,
	'z': 0.074,
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
		resp = append(resp, PTKeyPair{pt, byte(i), score(pt)})
	}
	return resp
}

func getCharFreq(s []byte) map[byte]float64 {
	resp := map[byte]float64{}

	for _, c := range s {
		var cStr byte
		if c >= 'a' && c <= 'z' {
			cStr = c
		} else if c >= 'A' && c <= 'Z' {
			cStr = byte(c + 32)
		} else {
			continue
		}

		if _, ok := resp[cStr]; !ok {
			resp[cStr] = 1
		} else {
			resp[cStr] += 1
		}

	}

	l := len(s)

	for k, v := range resp {
		resp[k] = v / float64(l)
	}

	return resp
}

func dot(v, canon map[byte]float64) float64 {
	acc := 0.0
	for k, v := range v {
		acc += canon[k] * v
	}
	return acc
}

func score(s []byte) float64 {
	freq := getCharFreq(s)
	return dot(freq, freqEn)
}

func scoreByWord(s []byte) float64 {
	words := bytes.Split(s, []byte(" "))
	acc := 0.0
	for _, w := range words {
		freq := getCharFreq(w)
		acc += dot(freq, freqEn)
	}
	return acc
}

func possibleKeysWithPT(ct []byte, num int) (resp []PTKeyPair) {
	ptkp := allPossibleDecryptions(ct)
	sort.Sort(sort.Reverse(byScore(ptkp)))
	return ptkp[0:num]
}

func countCaps(pt []byte) int {
	acc := 0
	for _, l := range pt {
		if l >= byte('A') && l <= byte('Z') {
			acc += 1
		}
	}
	return acc
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
