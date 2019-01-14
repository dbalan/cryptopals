package set2

import (
	"fmt"
	"github.com/dbalan/cryptopals/common"
	"testing"
)

func TestAESModeDetect(t *testing.T) {
	oracle := EncOracle

	for i := 0; i < 10; i++ {
		switch DetectAESMode(oracle) {
		case common.CBC:
			fmt.Println("CBC")
		case common.ECB:
			fmt.Println("ECB")
		}
	}
}
