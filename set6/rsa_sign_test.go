package set6

import (
	//	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

func TestSignatureForging(t *testing.T) {
	forgeSig([]byte("hi mom"), big.NewInt(0))
	// 	oddCubeRootPrefix(big.NewInt(2333))
}
