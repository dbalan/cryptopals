package set1

func FixedXOR(fst string, snd string) (string, error) {
	// assume both are of same length
	f := []byte(fst)
	s := []byte(snd)
	l := len(f)
	resp := []byte{}
	for i := 0; i < l; i++ {
		p := decodeHex(f[i])
		q := decodeHex(s[i])
		resp = append(resp, encodeHex(p^q))
	}
	return string(resp), nil
}
