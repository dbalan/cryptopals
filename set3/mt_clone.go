package set3

// takes a current MT19937 output and untempers state bit that pro// duced it
func untemper(y uint32) uint32 {
	y ^= (y >> l)
	y ^= (y << t) & c
	for i := 0; i < s; i++ {
		y ^= (y << s) & b
	}

	for i := 0; i < u; i++ {
		y ^= (y >> u)
	}
	return y
}

func cloneMT(seq []uint32) *mt19937 {
	if len(seq) != n {
		panic("check sequence len")
	}

	state := []uint32{}

	for _, s := range seq {
		state = append(state, untemper(s))
	}

	return &mt19937{state, 624}
}
