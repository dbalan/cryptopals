package set1

func countBits(b byte) int {
	result := 0

	for b > 0 {
		result += int(b & 1)
		b = b >> 1
	}

	return result
}

func hammingDistance(s1, s2 []byte) int {
	result := 0

	for i := 0; i < len(s1); i++ {
		c1 := s1[i]
		c2 := s2[i]
		xor := c1 ^ c2
		result += countBits(xor)
	}

	return result

}

func normDist(data []byte, keysize int, numBlocks int) float64 {
	cur := 0
	dist := 0.0
	i := 0
	for ; i < numBlocks && len(data) >= (cur+keysize*2); i++ {
		first := data[cur : cur+keysize]
		second := data[cur+keysize : cur+keysize*2]
		cur += keysize * 2
		dist += float64(hammingDistance(first, second)) / float64(keysize)
	}
	return dist / float64(i)
}

func findKeySize(data []byte) int {
	minSoFar := normDist(data, 2, 4)
	bestKeySize := 2

	for i := 3; i <= 40; i++ {
		dist := normDist(data, i, 3)
		//		fmt.Println("KEYSCORE: i: ", i, " dist: ", dist)
		if dist < minSoFar {
			bestKeySize = i
			minSoFar = dist
		}
	}
	return bestKeySize
}

func sliceAndDice(data []byte, keysize int) [][]byte {
	result := [][]byte{}

	for i := 0; i < keysize; i++ {
		transposedBlock := []byte{}
		for j := i; j < len(data); j += keysize {
			transposedBlock = append(transposedBlock, data[j])
		}
		result = append(result, transposedBlock)
	}

	return result

}

func FindKey(data []byte, keysize int) []byte {
	slices := sliceAndDice(data, keysize)
	key := []byte{}
	for _, s := range slices {
		_, k := BestPT(s)
		_ = possibleKeysWithPT(s, 5)

		//		fmt.Println("SLICE IND: ", si)
		//		for _, r := range resp {
		//			fmt.Printf("KEY: %d SCORE: %f PT: %s\n", r.Key, r.Score, r.PT[0:20])
		//		}
		//		fmt.Println("SLICE END")
		key = append(key, k)
	}
	return key
}
