package set3

import (
	"github.com/dbalan/cryptopals/common"
	"math/rand"
)

type MTStream struct {
	mt    *mt19937
	cur   uint32
	index int
}

func NewMTStream(seed uint32) *MTStream {
	mt := NewMT19937(seed)
	return &MTStream{mt, 0, 4}
}

func (mts *MTStream) Read(p []byte) (int, error) {
	nbuf := len(p)

	for i := 0; i < nbuf; i++ {
		if mts.index%4 == 0 {
			mts.cur = mts.mt.Rand()
			mts.index = 0
		}

		p[i] = byte(mts.cur & 0xff)
		mts.cur = mts.cur >> 8
		mts.index++
	}

	return nbuf, nil
}

func MTStreamCipher(in []byte, seed uint32) (out []byte, err error) {
	ks := make([]byte, len(in))
	mts := NewMTStream(seed)

	_, err = mts.Read(ks)
	if err != nil {
		return
	}

	return common.XORBlk(in, ks)
}

func MTSOracle(in []byte, key uint32) []byte {
	num := rand.Uint32() % 0xff
	prefix, _ := common.RandBytes(int(num))

	pt := append(prefix, in...)

	ct, _ := MTStreamCipher(pt, key)
	return ct
}
