package set2

import (
	"errors"
)

func PKCS7Padding(data []byte, lenExp int) ([]byte, error) {
	ldata := len(data)
	if ldata >= lenExp {
		return data, errors.New("BAD_DATA")
	}
	// pad lenExp - (ldata mod lenExp), lenExp - (ldata mod lenExp) times
	pad := lenExp - (ldata % lenExp)
	for i := 0; i < pad; i++ {
		data = append(data, byte(pad))
	}
	return data, nil
}

// assume the data is padded
func validateStripPKCS7Padding(data []byte) ([]byte, error) {
	length := len(data)
	pad := data[length-1]

	// find sufixes with same pad
	sfxLen := 0
	for i := length; i > 0; i-- {
		if data[i-1] != pad {
			break
		}
		sfxLen++
	}

	if sfxLen != int(pad) {
		return data, errors.New("WRONG_PAD_LENGTH")
	}

	return data[0 : length-sfxLen], nil
}

func PKCS7StripPadding(data []byte) []byte {
	dlen := len(data)
	last := data[dlen-1]

	if int(last) > len(data) {
		return data
	}

	for i := 0; i < int(last); i++ {
		if data[(dlen-1)-i] != last {
			return data
		}
	}

	// looks like padding!
	return data[0 : dlen-int(last)]
}
