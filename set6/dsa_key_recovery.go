package set6

import (
	"github.com/dbalan/cryptopals/dsa"
	"github.com/dbalan/cryptopals/sha"
	"math/big"
)

const (
	msg = `For those that envy a MC it can be hazardous to your health
So be friendly, a matter of life and death, just like a etch-a-sketch
`
	ystr = "84ad4719d044495496a3201c8ff484feb45b962e7302e56a392aee4" +
		"abab3e4bdebf2955b4736012f21a08084056b19bcd7fee56048e004" +
		"e44984e2f411788efdc837a0d2e5abb7b555039fd243ac01f0fb2ed" +
		"1dec568280ce678e931868d23eb095fde9d3779191b8c0299d6e07b" +
		"bb283e6633451e535c45513b2d33c99ea17"
	rstr = "548099063082341131477253921760299949438196259240"
	sstr = "857042759984254168557880549501802188789837994940"
)

func fromStr(s string) *big.Int {
	t := new(big.Int)
	t.SetString(s, 10)
	return t
}

func checkKey(x *big.Int) bool {
	rc, sc, err := dsa.Sign([]byte(msg), x)
	if err != nil {
		panic(err)
	}
	y, _ := new(big.Int).SetString(ystr, 16)
	return dsa.Verify([]byte(msg), rc, sc, y)
}

func getKey() *big.Int {
	hs := new(big.Int).SetBytes(sha.SHA([]byte(msg)))

	// y := fromStr(ystr)
	r := fromStr(rstr)
	s := fromStr(sstr)

	for k := 0; k <= 65536; k++ {
		x := dsa.ComputeKey(big.NewInt(int64(k)), r, s, hs)
		if checkKey(x) {
			return x
		}
	}
	panic("no suitable keys")
}
