package set5

import (
	"crypto/sha256"
	"fmt"
	"math/big"
	"math/rand"
	"strings"
)

func passwords() []string {
	// q n' d way to get password dict
	txt := "a abruptly absence account acute after again all amounts an and anderson animal annually anxiety any are arousing assess at backward be because been being believe below boilers breach bureau but by by called cape captain casualty cause centimeters charged choice cleaner clear company compartment consequently considerable considerable contain continue could couldn’t damage days delay derelict didn’t discovered dived dock docks done down dry engineers entered established extinguished eyes fifth figure filled for formed fortunately four from furnaces gaped gash had half halt hands have he his hold hole immediate immediately in indeed inexplicable inspect insurance into invaded invasion iron iron isosceles it it its itself job last launched lay leak least liverpool located losses lost made marine maritime meters miles moment moments monster’s motion must needed news no not numbers of on on one or out outrageous over paddle passions patched perfectly perforating piercing power proceeded prodigious produced proved public punch put recorded responsibility resulted sailing sailors scotia scotia sea shape sheet ships shoulder since so speed steam steamer’s straw such supposedly swamped symmetrical that the the their then there they they this this those three to tool toughness—plus triangle truly two two uncommon underside unfortunately vessels voyage was waterline way wheels which whose width with withdraw within without would"
	return strings.Split(txt, " ")
}

func randPasswd() string {
	n := rand.Uint32()
	pass := passwords()
	return pass[int(n)%len(pass)]
}

func randUint() *big.Int {
	x := &big.Int{}
	x.SetUint64(rand.Uint64())
	return x
}

func encodeUint64(x uint64) []byte {
	buf := make([]byte, 8)
	for i := 0; i < 8; i++ {
		buf[i] = byte(x & 0xff)
		x = x >> 8
	}
	return buf
}

func saltHmac(salt uint64, password string) *big.Int {
	encoded := encodeUint64(salt)
	h := sha256.New()
	h.Write(encoded)
	h.Write([]byte(password))
	val := fmt.Sprintf("%x", h.Sum(nil))
	ret := &big.Int{}
	ret.SetString(val, 16)
	return ret
}

func SHA256Int(A ...*big.Int) *big.Int {
	h := sha256.New()
	for _, val := range A {
		h.Write([]byte(val.Text(16)))
	}

	sum := fmt.Sprintf("%x", h.Sum(nil))
	ret := &big.Int{}
	ret.SetString(sum, 16)
	return ret
}

func HMAC_SHA256(key string, salt uint64) string {
	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("%x", salt)))
	h.Write([]byte(key))
	return fmt.Sprintf("%x", h.Sum(nil))
}
