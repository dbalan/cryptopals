package common

// FIXME: rewrite in loop
func decodeb64frag(frag []byte) (r []byte) {
	var lookup = map[byte]int{
		'A': 0, 'B': 1, 'C': 2, 'D': 3, 'E': 4, 'F': 5, 'G': 6, 'H': 7, 'I': 8, 'J': 9, 'K': 10, 'L': 11, 'M': 12, 'N': 13, 'O': 14, 'P': 15, 'Q': 16, 'R': 17, 'S': 18, 'T': 19, 'U': 20, 'V': 21, 'W': 22, 'X': 23, 'Y': 24, 'Z': 25, 'a': 26, 'b': 27, 'c': 28, 'd': 29, 'e': 30, 'f': 31, 'g': 32, 'h': 33, 'i': 34, 'j': 35, 'k': 36, 'l': 37, 'm': 38, 'n': 39, 'o': 40, 'p': 41, 'q': 42, 'r': 43, 's': 44, 't': 45, 'u': 46, 'v': 47, 'w': 48, 'x': 49, 'y': 50, 'z': 51, '0': 52, '1': 53, '2': 54, '3': 55, '4': 56, '5': 57, '6': 58, '7': 59, '8': 60, '9': 61, '+': 62, '/': 63, '=': 0}
	a := frag[0]
	b := frag[1]
	c := frag[2]
	d := frag[3]

	da := lookup[a]
	db := lookup[b]
	dc := lookup[c]
	dd := lookup[d]

	numbt := 3
	if c == '=' {
		numbt = 1
	} else if d == '=' {
		numbt = 2
	}

	switch numbt {
	case 3:
		b1 := (da << 2) + (db >> 4)
		b2 := (db << 4) + (dc >> 2)
		b3 := (dc << 6) + dd
		return []byte{byte(b1), byte(b2), byte(b3)}
	case 2:
		b1 := (da << 2) + (db >> 4)
		b2 := (db << 4) + (dc >> 2)
		return []byte{byte(b1), byte(b2)}
	case 1:
		b1 := (da << 2) + (db >> 4)
		return []byte{byte(b1)}
	}
	return r
}

// takes a b64 encoded string as input and returns correspoinding bytes
func DecodeB64(in []byte) (r []byte) {
	// for each 4 chars -> char * 4 = 24 == 24 = 3 * byte
	for i := 0; i < len(in); i = i + 4 {
		frag := decodeb64frag(in[i : i+4])
		r = append(r, frag...)
	}
	return r
}
