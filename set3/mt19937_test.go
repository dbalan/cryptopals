package set3

import (
	//	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMT1993Init(t *testing.T) {
	m := NewMT19937(123)
	last := m.state[n-1]
	assert.Equal(t, uint32(3987333491), last)

	// from https://github.com/james727/MTP
	seq := []uint32{
		2991312382,
		3062119789,
		1228959102,
		1840268610,
	}

	for i := 0; i < len(seq); i++ {
		r := m.Rand()
		assert.Equal(t, seq[i], r)
	}

}

func TestMT1993Twist(t *testing.T) {
	m := NewMT19937(0)
	//	1294739153
	for i := 0; i < 700; i++ {
		_ = m.Rand()
	}

	// from https://github.com/james727/MTP
	assert.Equal(t, uint32(1294739153), m.Rand())

}
