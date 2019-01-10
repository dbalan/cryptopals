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

func encodeHex(n byte) byte {
	switch {
	case int(n) >= 0 && int(n) <= 10:
		return n + '0'
	default:
		return byte(int(n)-10) + 'a'
	}
}

func decodeHex(n byte) byte {
	switch {
	case n >= '0' && n <= '9':
		return n - '0'
	case n >= 'a' && n <= 'f':
		return n - 'a' + 10
	case n >= 'A' && n <= 'F':
		return n - 'A' + 10
	}
	panic("errored out")
}
