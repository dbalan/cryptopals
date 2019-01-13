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
