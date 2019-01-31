package set3

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUntemper(t *testing.T) {
	cases := []struct {
		State uint32
		Value uint32
	}{
		{State: 2450663046, Value: 2422210778},
		{State: 3153726523, Value: 2353087594},
		{State: 1160128244, Value: 2569974907},
		{State: 1819853765, Value: 4049024520},
		{State: 1496513897, Value: 563593555},
		{State: 1293609056, Value: 1794197249},
		{State: 594681420, Value: 2434290377},
		{State: 1278130738, Value: 4222178191},
		{State: 1990424729, Value: 2381045132},
		{State: 3352857227, Value: 1294739153},
	}

	for _, cs := range cases {
		st := untemper(cs.Value)
		assert.Equal(t, cs.State, st)
	}

}

func TestClone(t *testing.T) {
	var someSeed uint32 = 234

	mt := NewMT19937(someSeed)
	seq := []uint32{}

	for i := 0; i < 624; i++ {
		seq = append(seq, mt.Rand())
	}

	cloned := cloneMT(seq)

	for i := 0; i < 10; i++ {
		assert.Equal(t, cloned.Rand(), mt.Rand())
	}

}
