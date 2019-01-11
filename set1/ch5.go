package set1

func repeatingXOR(plainText string, key string) string {
	pt := []byte(plainText)

	lkey := len(key)
	k := []byte(key)
	ct := ""

	for i := 0; i < len(pt); i++ {
		ct += hexPrettyPrint(pt[i] ^ k[i%lkey])
	}
	return ct
}

// FIXME: both functions can be merged
func decryptRepeatXOR(ct []byte, key []byte) []byte {
	lkey := len(key)
	resp := []byte{}
	for i := 0; i < len(ct); i++ {
		resp = append(resp, ct[i]^key[i%lkey])
	}

	return resp
}
