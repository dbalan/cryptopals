package set2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAESModeDetect(t *testing.T) {
	for i := 0; i < 10; i++ {
		detected, actual := DetectAESMode(EncOracle)
		assert.Equal(t, detected, actual)
	}
}
