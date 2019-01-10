package set1

import (
	"fmt"
	"testing"
)

const ct = "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"

func TestCh3(t *testing.T) {
	fmt.Println("PlainText: ", BestPT(ct))
}
