package common

func hexPrettyPrint(b byte) string {
	left := EncodeHex(b >> 4 & 0xF)
	right := EncodeHex(b & 0xF)
	return string(left) + string(right)
}

func EncodeHex(n byte) byte {
	switch {
	case int(n) >= 0 && int(n) <= 9:
		return n + '0'
	default:
		return byte(int(n)-10) + 'a'
	}
}

func EncodeHexString(in []byte) (r string) {
	for _, b := range in {
		r += hexPrettyPrint(b)
	}
	return r
}
