package set3

const (
	// coefficients for MT19937
	w = 32  // wordsize
	n = 624 // degree of recurrance
	m = 397 // middle word
	r = 31  // seperation point of one word

	a = 0x9908B0DF // coefficients of the rational normal form twist matrix

	// TGFSR(r) - tempering bit masks
	b = 0x9D2C5680
	c = 0xEFC60000

	// TGFSR(r) - tempering bit shifts
	s = 7
	t = 15

	// Mersenne Twister tempering bit shifts/masks
	u = 11
	d = 0xFFFFFFFF
	l = 18

	// constant for MT19937 (not sure why!)
	f = 1812433253
)

var (
	// masks for MT19937
	lower uint32 = (1 << r) - 1
	upper uint32 = (1 << r) // & 0x7fffffff
)

type mt19937 struct {
	state []uint32
	index int
}

func (mt *mt19937) seed(seed uint32) {
	state := make([]uint32, n)

	state[0] = seed
	for i := 1; i < n; i++ {
		state[i] = f*(state[i-1]^(state[i-1]>>(w-2))) + uint32(i)
	}

	mt.state = state

	// force a twist first time we generate a random number.
	// since state0 can't be used to generate random numbers
	mt.index = n
}

func (mt *mt19937) twist() {
	for cur := 0; cur < n; cur++ {
		next := (cur + 1) % n
		temp := (mt.state[cur] & upper) + (mt.state[next] & lower)
		shifted := temp >> 1

		if temp%2 != 0 {
			shifted = shifted ^ a
		}

		mt.state[cur] = mt.state[(cur+m)%n] ^ shifted
	}
	mt.index = 0
}

func (mt *mt19937) Rand() uint32 {
	if mt.index >= n {
		mt.twist()
	}

	y := mt.state[mt.index]

	y ^= (y >> u)
	y ^= (y << s) & b
	y ^= (y << t) & c
	y ^= (y >> l)

	mt.index += 1

	return y
}

func NewMT19937(seed uint32) *mt19937 {
	if seed <= 0 {
		seed = 5489
	}

	m := mt19937{}
	m.seed(seed)
	return &m
}
