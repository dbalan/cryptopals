package set1

import (
	"strings"
)

var freqEn = map[string]float64{
	"a": 8.167,
	"b": 1.492,
	"c": 2.782,
	"d": 4.253,
	"e": 12.70,
	"f": 2.228,
	"g": 2.015,
	"h": 6.094,
	"i": 6.966,
	"j": 0.153,
	"k": 0.772,
	"l": 4.025,
	"m": 2.406,
	"n": 6.749,
	"o": 7.507,
	"p": 1.929,
	"q": 0.095,
	"r": 5.987,
	"s": 6.327,
	"t": 9.056,
	"u": 2.758,
	"v": 0.978,
	"w": 2.360,
	"x": 0.150,
	"y": 1.974,
	"z": 0.074,
}

func decryptSingleXOR(ct string, key byte) string {
	raw := decodeHexString(ct)

	resp := []byte{}
	for _, c := range raw {
		resp = append(resp, c^key)
	}
	return string(resp)
}

func allPossibleDecryptions(ct string) []string {
	resp := []string{}

	for i := 0; i <= 0xff; i++ {
		resp = append(resp, decryptSingleXOR(ct, byte(i)))
	}
	return resp
}

func getCharFreq(s string) map[string]float64 {
	resp := map[string]float64{}

	for _, c := range s {
		var cStr string

		if c >= 'a' && c <= 'z' {
			cStr = string(c)
		} else if c >= 'A' && c <= 'Z' {
			cStr = string(c + 32)
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

func dot(v, canon map[string]float64) float64 {
	acc := 0.0
	for k, v := range v {
		acc += canon[k] * v
	}
	return acc
}

func score(s string) float64 {
	words := strings.Split(s, " ")
	acc := 0.0
	for _, w := range words {
		freq := getCharFreq(w)
		acc += dot(freq, freqEn)
	}
	return acc
}

func BestPT(ct string) string {
	curHigh := 0.0
	result := ""

	for _, pt := range allPossibleDecryptions(ct) {
		curScore := score(pt)
		if curScore > curHigh {
			curHigh = curScore
			result = pt
		}
	}
	return result
}
