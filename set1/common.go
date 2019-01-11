package set1

func encodeHex(n byte) byte {
	switch {
	case int(n) >= 0 && int(n) <= 9:
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

func decodeHexString(s string) []byte {
	// 1 byte == 8 bits, 2^4 *2
	l := len(s)
	resp := []byte{}
	if l%2 != 0 {
		s = "0" + s
	}

	i := 0
	for i < l {
		resp = append(resp, decodeHex(byte(s[i]))<<4+decodeHex(byte(s[i+1])))
		i += 2
	}
	return resp
}

// b00001011 -> "0b"
// int() -> 11
// "0b"
func hexPrettyPrint(b byte) string {
	left := encodeHex(b >> 4 & 0xF)
	right := encodeHex(b & 0xF)
	return string(left) + string(right)
}

// takes a b64 encoded string as input and returns correspoinding bytes
func base64decode(in []byte) []byte {
	return in
}
