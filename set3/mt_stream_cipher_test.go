package set3

import (
	//	"fmt"
	"math/rand"
	"testing"

	"github.com/dbalan/cryptopals/common"
	"github.com/stretchr/testify/assert"
	"time"
)

func TestMTStream(t *testing.T) {
	mts := NewMTStream(uint32(123))

	buf := make([]byte, 8)

	n, _ := mts.Read(buf)
	assert.Equal(t, 8, n)

	expected := []byte{}
	for _, r := range []uint32{2991312382, 3062119789} {
		for i := 0; i < 4; i++ {
			expected = append(expected, byte(r&0xff))
			r >>= 8
		}

	}
	assert.Equal(t, expected, buf)
}

func TestMTStreamAttack(t *testing.T) {
	rand.Seed(time.Now().Unix())
	key := rand.Uint32() & 0xffff
	knownPT := []byte("AAAAAAAAAAAAAA")
	ct := MTSOracle(knownPT, key)

	suffix := ct[len(ct)-14 : len(ct)]
	keystream, _ := common.XORBlk(suffix, knownPT)

	checkSeed := func(s uint32, suc chan<- uint32) {
		mts := NewMTStream(uint32(s))
		buf := make([]byte, len(ct))

		n, _ := mts.Read(buf)
		if n < len(buf) {
			t.Errorf("bad library")
		}

		// check buf[len(ct)-14:len(ct)] == keystream
		if common.EqualBlocks(buf[len(ct)-14:len(ct)], keystream) {
			suc <- s
		}
		return
	}

	suc := make(chan uint32)
	for s := uint32(0); s < 2<<16+1; s++ {
		go checkSeed(s, suc)
	}
	res := <-suc
	assert.Equal(t, key, res)
}
