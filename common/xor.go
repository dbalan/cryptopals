package common

// xors fist block in place
func XORBlk(fst, snd []byte) ([]byte, error) {
	// assume both are of same length
	l := len(fst)
	if l != len(snd) {
		return nil, BadDataErr
	}
	for i := 0; i < l; i++ {
		fst[i] = byte(fst[i] ^ snd[i])
	}
	return fst, nil
}
