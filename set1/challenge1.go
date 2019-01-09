package set1

import (
	"fmt"
	"strconv"
)

var Lookup = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz123456789+/")

// covert a 2 sextets into base64 bytes
func convert(k int64) []byte {
	lower := k & 0x3F
	higher := (k >> 6) & 0x3F
	return []byte{Lookup[higher], Lookup[lower]}
}

func hex(a string) int64 {
	parsed, _ := strconv.ParseInt(a, 16, 64)
	return parsed
}

func hex2base64(a string) string {
	// 3 hex points = 12 bits == 2 base64 values
	l := len(a)
	i := 0
	for i < l {
		hexPoints := a[i : i+3]
		fmt.Println(string(convert(hex(hexPoints))))
		if (i + 3) > l {
			// fragment
			fmt.Println("fragment", a[i+3:l])
		}
		i = i + 3
	}
	return a
}
