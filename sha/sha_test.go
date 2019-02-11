package sha

import (
	//	"fmt"
	"github.com/dbalan/cryptopals/common"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestPreprocessing(t *testing.T) {
	input := []byte("The quick brown fox jumps over the lazy dog")

	out := Preprocess(input)
	assert.Equal(t, 0, (len(out)*8)%512)
}

func TestBEEncode(t *testing.T) {
	outp := BEEncodeUint64(0xdeadbeef)
	expected := []byte{byte(0), byte(0), byte(0), byte(0), byte(222), byte(173), byte(190), byte(239)}

	assert.Equal(t, expected, outp)
}

func TestSHA(t *testing.T) {
	testCases := []struct {
		In  string
		Out string
	}{
		{"The quick brown fox jumps over the lazy dog", "2fd4e1c67a2d28fced849ee1bb76e7391b93eb12"},
		{"The quick brown fox jumps over the lazy cog", "de9f2c7fd25e1b3afad3e85a0bd17d9b100db4b3"},
		{"", "da39a3ee5e6b4b0d3255bfef95601890afd80709"},
	}
	for _, tc := range testCases {
		result := SHA([]byte(tc.In))
		assert.Equal(t, tc.Out, common.EncodeHexString(result))
	}
}

var res []byte

func BenchmarkSHA1(b *testing.B) {
	var r []byte
	for i := 0; i < b.N; i++ {
		r = SHA([]byte("what"))
	}
	res = r
}

func BenchmarkSHA1RunningTime(b *testing.B) {
	var r []byte
	var ts []time.Duration

	for i := 0; i < b.N; i++ {
		start := time.Now()
		r = SHA([]byte("hello world"))
		extm := time.Since(start)
		ts = append(ts, extm)
	}
	res = r

	common.StandardDeviation(ts)
}

func TestSHA1RunningTime(t *testing.T) {
	var r []byte
	var ts []time.Duration

	for p := 0; p < 10; p++ {
		for i := 0; i < 1000; i++ {
			start := time.Now()
			r = SHA([]byte("hello worldqweqyuuqywequw"))
			//		time.Sleep(1 * time.Millisecond)
			extm := time.Since(start)
			ts = append(ts, extm)
		}
		res = r

		common.StandardDeviation(ts)
	}
}
