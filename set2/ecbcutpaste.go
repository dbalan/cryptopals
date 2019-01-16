package set2

func encryptEmail(email string, key []byte) []byte {
	pt := profileFor(email).encode()
	enc, err := EncAES128ECB([]byte(pt), key)
	if err != nil {
		panic(err)
	}

	return enc
}

func decryptKV(ct []byte, key []byte) string {
	dec, err := DecAES128ECB(ct, key)
	if err != nil {
		panic(err)
	}
	return string(dec)
}
