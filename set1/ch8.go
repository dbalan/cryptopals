package set1

func repeatingBlocks(ct []byte) bool {
	blocks := [][]byte{}

	for i := 0; i < len(ct); i += 16 {
		blocks = append(blocks, ct[i:i+16])
	}

	for i := 0; i < len(blocks); i++ {
		for j := i + 1; j < len(blocks); j++ {
			if hammingDistance(blocks[i], blocks[j]) == 0 {
				return true
			}
		}
	}

	return false
}
