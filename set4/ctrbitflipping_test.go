package set4

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCTRBitflipping(t *testing.T) {
	status, err := Attack()
	assert.Nil(t, err)

	assert.True(t, status)
}
