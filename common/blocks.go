package common

func EqualBlocks(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}

	for ci, cv := range a {
		if cv != b[ci] {
			return false
		}
	}

	return true
}

// a is of lenth n * bs where n is a real number
// split into a blocks
func Blocks(d []byte, bs int) [][]byte {
	resp := [][]byte{}
	for i := 0; i < len(d); i += bs {
		resp = append(resp, d[i:i+bs])
	}
	return resp
}

func PadTo(cur int, size int) int {
	// pad to next blocks
	left := cur % size
	if left != 0 {
		return cur + (size - left)
	}

	return cur
}
