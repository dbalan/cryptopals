package set1

import (
	"fmt"
	"testing"
)

const ct = "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"

func TestCh3(t *testing.T) {
	fmt.Println("PlainText: ", BestPT(ct))
}

func TestCh4CT(t *testing.T) {
	cipher := "7b5a4215415d544115415d5015455447414c155c46155f4058455c5b523f"
	fmt.Println("PlainText ch4: ", BestPT(cipher))
}
