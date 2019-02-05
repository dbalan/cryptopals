package common

func PackUint32(w ...byte) uint32 {
	l := len(w)

	var acc uint32
	for i := 0; i < l; i++ {
		acc = acc | uint32(w[l-i-1])<<uint(i*8)
	}
	return acc
}
