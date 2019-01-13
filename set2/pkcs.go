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
