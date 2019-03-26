package set6

import (
	"errors"
	"io/ioutil"
	"math/big"
	"sort"
	"strings"

	"github.com/dbalan/cryptopals/dsa"
)

type Pair struct {
	Msg string
	R   *big.Int
	S   *big.Int
	M   *big.Int
}

type ByR []Pair

func (b ByR) Len() int      { return len(b) }
func (b ByR) Swap(i, j int) { b[i], b[j] = b[j], b[i] }

func (b ByR) Less(i, j int) bool {
	r1 := b[i].R
	r2 := b[j].R
	return r1.Cmp(r2) == -1
}

func readData() ([]Pair, error) {
	data, err := ioutil.ReadFile("./44.txt")
	if err != nil {
		return nil, err
	}
	text := string(data)

	chunks := strings.Split(text, "\n")

	pairs := []Pair{}
	for i := 0; i < len(chunks); i += 4 {
		msg := chunkVal(chunks[i])
		sStr := chunkVal(chunks[i+1])
		rStr := chunkVal(chunks[i+2])
		mStr := chunkVal(chunks[i+3])
		p := Pair{Msg: msg, R: fromStr(rStr, 10),
			S: fromStr(sStr, 10), M: fromStr(mStr, 16)}

		pairs = append(pairs, p)
	}

	return pairs, nil
}

// msg: But in a in an' a out de dance em
// But in a in an' a out de dance em
func chunkVal(s string) string {
	part := strings.Split(s, ": ")
	return part[1]
}

func findRepeatingK() (p1, p2 *Pair, err error) {
	pairs, err := readData()
	if err != nil {
		return
	}

	sort.Sort(ByR(pairs))

	prev := pairs[0]
	pairs = pairs[1:len(pairs)]

	for _, p := range pairs {
		if prev.R.Cmp(p.R) == 0 {
			return &prev, &p, nil
		}
	}

	return nil, nil, errors.New("No pairs found")
}

/*
 *     M1 - M2
 * K = ------  mod q
 *     S1 - S2
 */
func computeK(p1, p2 *Pair) *big.Int {
	q := fromStr("f4f47f05794b256174bba6e9b396a7707e563c5b", 16)

	m1 := p1.M
	m2 := p2.M
	s1 := p1.S
	s2 := p2.S

	lower := new(big.Int).Set(s2)

	lower.Sub(lower, s1)

	// (S1-S2)**-1
	lower.Exp(lower, big.NewInt(-1), q)

	upper := new(big.Int).Set(m1)

	if upper.Cmp(m2) == 1 {
		upper.Sub(upper, m2)
	} else {
		upper.Sub(m2, upper)
	}

	upper.Mod(upper, q)
	lower.Mul(lower, upper).Mod(lower, q)
	return lower
}

func findPrivateKey() *big.Int {

	p1, p2, err := findRepeatingK()
	if err != nil {
		panic(err)
	}
	k := computeK(p1, p2)

	return dsa.ComputeKey(k, p1.R, p1.S, p1.M)
}
