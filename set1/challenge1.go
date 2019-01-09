package set1

import (
	"github.com/pkg/errors"
	"strconv"
)

const lookup = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

// Hex2B64 converts a hexadecimal string to its base64 representation
func Hex2B64(in string) (string, error) {
	resp := []byte{}
	// in is a UTF-8 encoded hex value
	// work on bytes
	val := []byte(in)
	l := len(val)

	// group them into 3 (3 in hex == 12 bits == 2 sextets)
	i := 0
	for i < l {
		if l < i+3 {
			// fragment
			return "", errors.New("NOT IMPL")
		}
		point := val[i : i+3]
		frag, err := convertPoint(point)
		resp = append(resp, frag...)
		if err != nil {
			return "", err
		}

		i += 3
	}
	return string(resp), nil
}

func convertPoint(pt []byte) ([]byte, error) {
	// TODO: implement parsing ourselves
	val, err := strconv.ParseInt(string(pt), 16, 64)
	if err != nil {
		return nil, err
	}
	low := val & 0x3F
	high := (val >> 6) & 0x3F
	return []byte{lookup[high], lookup[low]}, nil
}
